"use client"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { Button } from "@/components/ui/button";
import { Upload } from "lucide-react";
import { Input } from "@/components/ui/input";
import { useState, ChangeEvent, FormEvent } from "react";
import * as XLSX from 'xlsx';

export function UploadCSV() {
  const [file, setFile] = useState<File | null>(null);

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      setFile(event.target.files[0]);
    }
  };

  const handleUpload = async (event: FormEvent) => {
    event.preventDefault();

    if (!file) return;

    const socket = new WebSocket('ws://localhost:3001/upload');

    socket.onopen = function () {
      console.log('WebSocket connection established.');
    };

    socket.onerror = function (error) {
      console.error('WebSocket error:', error);
    };

    socket.onclose = function () {
      console.log('WebSocket connection closed.');
    };

    const reader = new FileReader();
    reader.onload = async (e) => {
      if (!e.target?.result) return;

      const data = new Uint8Array(e.target.result as ArrayBuffer);
      const workbook = XLSX.read(data, { type: 'array' });
      const sheetName = workbook.SheetNames[0];
      const worksheet = workbook.Sheets[sheetName];
      const rows: any[] = XLSX.utils.sheet_to_json(worksheet, { header: 1 });
      let index = 0;
      for (const row of rows) {
        index++;
        if (row.length < 3) continue;
        const [name, email, pincodes] = row;
        const pincodesArray = pincodes.split(',').map((pincode: string) => pincode.trim());
        const correctPincodeStrings = pincodesArray.filter((pincode: string) => pincode.length === 6);
        const pincodes_int = correctPincodeStrings.map((pincode: string) => parseInt(pincode));
        const data = {
          name,
          email,
          pincodes: pincodes_int,
        }
        console.log(data);
        if(socket.readyState === WebSocket.OPEN){
          socket.send(JSON.stringify(data));
        } else {
          console.error('WebSocket connection is not open.');
        }
      }
      socket.send('done');
    };
    reader.readAsArrayBuffer(file);
  };

  return (
    <Dialog>
      <Tooltip>
        <TooltipTrigger asChild>
          <DialogTrigger asChild>
            <Button variant={"link"}>
              <Upload />
            </Button>
          </DialogTrigger>
        </TooltipTrigger>
        <TooltipContent>
          <p>Upload CSV</p>
        </TooltipContent>
      </Tooltip>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Upload CSV</DialogTitle>
        </DialogHeader>
        <DialogDescription>
          <form className="w-full flex space-x-2" onSubmit={handleUpload}>
            <Input type="file" onChange={handleFileChange} />
            <Button type="submit"><Upload className="w-4 h-4" /></Button>
          </form>
        </DialogDescription>
      </DialogContent>
    </Dialog>
  );
}
