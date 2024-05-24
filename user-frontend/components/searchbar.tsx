"use client";

import { Input } from "./ui/input";
import { Search } from "lucide-react";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { MerchantData, PincodeData } from "@/lib/types";
import { useStore } from "@/lib/store";
import axios from "axios";

interface SearchBarProps {
  pincode?: string;
}

export function SearchBar(props: SearchBarProps) {
  const [pincode, setPincode] = useState(props.pincode ? props.pincode : "");
  const router = useRouter();

  const { setMerchantData, setTimeTakenForRequest } = useStore();

  const createEventSource = async () => {
    const eventSource = new EventSource("/api/pincode");
    eventSource.onmessage = (event) => {
      console.log(event.data);
    };
    eventSource.dispatchEvent
  }

  const fetchAndCachePincodeData = async () => {

    const data: PincodeData = await fetch("/api/pincode", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ pincode: pincode }),
    }).then(async (res) => await res.json());
    // console.log(data)
    createEventSource();
    const { current, cache } = data
    // window.localStorage.setItem('merchantData', JSON.stringify(cache))
    return current["merchantList"];
  };

  // const retrieveCurrentPincodeDataFromLocalStorage = () => {
  //   let data: {
  //     "pincode": string,
  //     "merchantList": MerchantData[]
  //   }[] = JSON.parse(
  //     window.localStorage.getItem("merchantData") || "[]"
  //   );
  //   let currentPincodeData = [];
  //   for (let i = 0; i < data.length; i++) {
  //     if (data[i].pincode == pincode) {
  //       currentPincodeData.push(...data[i].merchantList);
  //     }
  //   }
  //   return currentPincodeData;
  // };

  // const renewData = async () => {
  //   let data = await fetchAndCachePincodeData();
  //   window.localStorage.setItem("merchantData", JSON.stringify(data));
  //   let newPincodeData = retrieveCurrentPincodeDataFromLocalStorage();
  //   setMerchantData(newPincodeData);
  // };


  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    router.push(`/search?pincode=${pincode}`);
    let start = performance.now();
    // if (!window.localStorage.getItem("merchantData")) {
    //   let data = await fetchAndCachePincodeData();
    //   setMerchantData(data);
    // } else {
    //   let currentPincodeData = retrieveCurrentPincodeDataFromLocalStorage();
    //   if (currentPincodeData.length === 0) {
    //     await renewData();
    //   } else {
    //     setMerchantData(currentPincodeData);
    //   }
    // }
    let currentMerchantList = await fetchAndCachePincodeData()
    setMerchantData(currentMerchantList);
    let end = performance.now();
    setTimeTakenForRequest(end - start); // in ms

  };

  return (
    <>
      <div className="relative w-full mb-10">
        <div className="max-w-xl mx-auto">
          <div className="absolute inset-y-0 flex items-center ps-3 pointer-events-none">
            <Search className="text-sm text-gray-400 dark:text-gray-600" />
          </div>
          <form onSubmit={handleSubmit}>
            <Input
              placeholder="Enter your pincode"
              type="number"
              defaultValue={props.pincode ? props.pincode : ""}
              onChange={(e) => {
                setPincode(e.target.value);
              }}
              className="bg-sky-50 dark:bg-gray-800 transition-all duration-300 p-6 ps-12 text-md rounded-full focus:shadow-md"
            />
          </form>
        </div>
      </div>
    </>
  );
}
