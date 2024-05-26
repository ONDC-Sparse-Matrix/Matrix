import { MerchantCard } from "./merchant-card";
import { TimeData } from "./time-data";
import { Spinner } from "@/components/ui/spinner";
import { useEffect } from "react";
import { useState } from "react";

import axios from "axios";
import { FETCH_PINCODE_DATA } from "@/lib/endpoints";
import { PincodeData, PincodeDataCache } from "@/lib/types";
import { testPincodeData_1, testPincodeData_2 } from "@/lib/test-search-data";
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

    //* Uncomment this to fetch data from backend.
    // axios
    //   .post<PincodeData>(`${FETCH_PINCODE_DATA}/${pincode}`, {
    //     session: props.sessionId,
    //   })
    //   .then((res) => setData(res.data));

    //* Test data here
    await new Promise((resolve) =>
      setTimeout(() => {
        if (pincode.toString() == testPincodeData_1.current.pincode) {
          setData(testPincodeData_1.current);
          setCache(testPincodeData_1.cache);
        } else if (pincode.toString() == testPincodeData_2.current.pincode) {
          setData(testPincodeData_2.current);
          setCache(testPincodeData_2.cache);
        } else {
          searchCache(pincode).then((res) => setData(res));
        }
        resolve(void 0);
      }, 1000)
    );

    let end = performance.now();
    setTimeTakenForRequest(end - start); // in ms
    setIsLoading(false);
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
