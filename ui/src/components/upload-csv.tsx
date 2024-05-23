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
  
  export function UploadCSV() {
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
            <form className="w-full flex space-x-2">
              <Input type="file" />
              <Button type="submit"><Upload className="w-4 h-4" /></Button>
            </form>
          </DialogDescription>
        </DialogContent>
      </Dialog>
    );
  }
  