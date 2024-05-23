import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Upload } from "lucide-react";
import { useState } from "react";

export function AddMerchantCard() {
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [pincodes, setPincodes] = useState("");

  const [file, setFile] = useState<File | null>(null);

  const uploadFile = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!file) return;
    //TODO: Arpan - idhar add krna hai streaming ka logic
  };

  const addMerchant = async (e: React.FormEvent) => {
    e.preventDefault();
    const pincodesArray = pincodes
      .split(",")
      .map((pincode) => parseInt(pincode.trim()));
    //TODO: Arpan - idhar daalna hai uplaod ka logic
  };

  return (
    <Card className="flex flex-col md:flex-row">
      <div className="flex flex-col gap-4 p-6 md:w-1/2">
        <h3 className="text-xl font-medium">Add a Merchant Entry</h3>
        <form className="flex flex-col gap-4" onSubmit={addMerchant}>
          <div className="grid gap-1">
            <Label htmlFor="email">Email</Label>
            <Input
              id="email"
              placeholder="merchant@example.com"
              type="email"
              required
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
          <div className="grid gap-1">
            <Label htmlFor="name">Name</Label>
            <Input
              id="name"
              placeholder="Suresh"
              required
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          <div className="grid gap-1">
            <Label htmlFor="pincode">Pincodes</Label>
            <Input
              id="pincode"
              required
              placeholder="123456, 789210 ..."
              onChange={(e) => setPincodes(e.target.value)}
            />
            <p className="text-xs text-muted-foreground">
              Put pincodes separated by a comma (, )
            </p>
          </div>
          <Button className="w-full" type="submit">
            Add Entry
          </Button>
        </form>
      </div>

      <div className="flex flex-col items-center justify-center gap-4 border-t md:border-l border-gray-200 p-6 md:w-1/2 dark:border-gray-800">
        <div className="flex flex-col items-center justify-center gap-2">
          <Upload className="h-8 w-8 text-gray-500 dark:text-gray-400" />
          <h3 className="text-lg font-medium">Upload XLSX File</h3>
          <p className="text-sm text-gray-500 dark:text-gray-400">
            Click to upload an excel file.
          </p>
        </div>
        <form onSubmit={uploadFile}>
          <div className="flex gap-2 items-center">
            <Input
              accept=".xlsx"
              type="file"
              onChange={(e) => {
                const files = e.target.files;
                if (files) {
                  setFile(files[0]);
                }
              }}
              className="flex-1"
            />
            <Button type="submit" size={"icon"}>
              <Upload className="w-4 h-4" />
            </Button>
          </div>
        </form>
      </div>
    </Card>
  );
}
