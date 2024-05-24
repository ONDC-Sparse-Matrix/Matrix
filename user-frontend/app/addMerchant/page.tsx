"use client";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Plus } from "lucide-react";

export default function AddMerchant() {


  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("Form submitted");
  }
  return (
    <>
      <Card className="w-[380px]">
        <CardHeader>
          <CardTitle>Add Merchant</CardTitle>
        </CardHeader>
        <form onSubmit={handleSubmit}>
          <CardContent>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="name">Name</Label>
                <Input id="name" placeholder="Enter name" />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="email">Email</Label>
                <Input id="email" placeholder="Enter email" />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="pincode">Pincode</Label>
                <Input id="pincode" placeholder="Enter pincode array" />
              </div>
            </div>
          </CardContent>
          <CardFooter className="flex justify-between">
            <Button type="submit" className="w-full">
              <Plus className="h-4 w-4 mr-2" />
              Add Merchant
            </Button>
          </CardFooter>
        </form>
      </Card>
    </>
  );
}
