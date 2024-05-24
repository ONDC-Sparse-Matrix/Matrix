"use client";

import { useSearchParams } from "next/navigation";
import { MerchantCard } from "@/components/merchant-card";
import { useStore } from "@/lib/store";
import { SearchBar } from "@/components/searchbar";
import { Clock, AlertCircleIcon } from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";

export default function Search() {
  const searchParams = useSearchParams();
  const pincode = searchParams.get("pincode");
  const { timeTakenForRequest, merchantData } = useStore();

  return (
    <div className="h-screen max-w-xl mx-auto pt-20">
      <SearchBar pincode={pincode ? pincode : ""} />
      {merchantData.length > 0 ? (
        <>
          <Alert className="mb-8 shadow-inner">
            <Clock className="h-4 w-4" />
            <AlertTitle className="font-bold">Heads up!</AlertTitle>
            <AlertDescription>
              Time taken for request:{" "}
              <span className="text-blue-500 font-bold">
                {timeTakenForRequest/1000}s
              </span>
            </AlertDescription>
          </Alert>
          <h1 className="text-lg uppercase tracking-widest font-bold text-center text-muted-foreground">
            Search Results for{" "}
            <span className="text-foreground font-extrabold">{pincode}</span>
          </h1>
          {merchantData.map((merchant, index) => (
            <MerchantCard {...merchant} />
          ))}
        </>
      ) : (
        <Alert className="mb-8" variant={'destructive'}>
          <AlertCircleIcon className="h-4 w-4" />
          <AlertTitle className="font-bold">No Merchants Available :(</AlertTitle>
          <AlertDescription>
            There seems to be no merchants serving at {pincode}
          </AlertDescription>
        </Alert>
      )}
    </div>
  );
}
