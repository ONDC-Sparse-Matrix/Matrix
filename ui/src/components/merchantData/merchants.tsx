import { MerchantCard } from "./merchant-card";
import { TimeData } from "./time-data";
import { Spinner } from "@/components/ui/spinner";
import { useEffect } from "react";
import { useState } from "react";

import axios from "axios";
import { FETCH_PINCODE_DATA } from "@/lib/endpoints";
import { PincodeData, PincodeDataCache } from "@/lib/types";
import { testPincodeData_1, testPincodeData_2 } from "@/lib/test-data";
import { updateCache, searchCache } from "@/lib/db";

interface MerchantsProps {
  pincode: number;
  sessionId: string;
}

export function Merchants(props: MerchantsProps) {
  const [data, setData] = useState<PincodeData["current"] | undefined>(
    undefined
  );
  const [cache, setCache] = useState<PincodeDataCache | undefined>(undefined);
  const [timeTakenForRequest, setTimeTakenForRequest] = useState<number>(0);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  // The main function
  const getMerchantData = async (pincode: string) => {
    setIsLoading(true);
    let start = performance.now();

    const cache = await searchCache(pincode);
    if (cache) {
      console.log("Cache found")
      setData(cache);
      let end = performance.now();
      setTimeTakenForRequest(end - start); // in ms
      setIsLoading(false);
    } else {
      axios
        .post<PincodeData>(`http://localhost:3001/pincode/${pincode}`, {
          session: props.sessionId,
        })
        .then((res: any) => { setData(res.data); console.log(res.data); });

      let end = performance.now();
      setTimeTakenForRequest(end - start); // in ms
      setIsLoading(false);
    }
  };

  useEffect(() => {
    const func = async () => {
      await getMerchantData(props.pincode.toString());
    };
    func();
    console.log("pincode changed to: ", props.pincode.toString());
  }, [props.pincode]);

  useEffect(() => {
    const cacheData = async (cache: PincodeDataCache | undefined) => {
      if (cache) {
        console.log("Updating cache");
        await updateCache(cache);
        console.log("Cache updated");
      }
    };
    cacheData(cache);
  }, [cache]);

  return (
    <>
      {isLoading ? (
        <div className="w-full flex justify-center">
          <Spinner className="mt-6 w-8 h-8" />
        </div>
      ) : (
        <>
          <TimeData
            success={data ? true : false}
            pincode={props.pincode}
            time={timeTakenForRequest}
          />
          {data?.merchantList.map((merchant, index) => (
            <MerchantCard key={index} {...merchant} />
          ))}
        </>
      )}
    </>
  );
}
