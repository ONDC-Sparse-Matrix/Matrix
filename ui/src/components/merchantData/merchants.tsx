import { MerchantCard } from "./merchant-card";
import { TimeData } from "./time-data";
import { Spinner } from "@/components/ui/spinner";
import { useEffect } from "react";
import axios from "axios";
import { FETCH_PINCODE_DATA } from "@/lib/endpoints";
import { PincodeData } from "@/lib/types";
import { useState } from "react";

import { testPincodeData } from "@/lib/test-data";

interface MerchantsProps {
  pincode: number;
}

export function Merchants(props: MerchantsProps) {
  const [data, setData] = useState<PincodeData>();
  const [timeTakenForRequest, setTimeTakenForRequest] = useState<number>(0);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const getMerchantData = async (pincode: string) => {
    setIsLoading(true);
    let start = performance.now();
    //* Uncomment this to fetch data from backend.
    // setData(
    //   await axios
    //     .get<PincodeData>(`${FETCH_PINCODE_DATA}/${pincode}`)
    //     .then((res) => res.data)
    // );

    //* Test data here
    await new Promise((resolve) =>
      setTimeout(() => {
        setData(
          pincode.toString() == testPincodeData.current.pincode
            ? testPincodeData
            : undefined
        );
        resolve(void 0);
      }, 3000)
    );

    let end = performance.now();
    setTimeTakenForRequest(end - start); // in ms
    setIsLoading(false);
  };

  useEffect(() => {
    getMerchantData(props.pincode.toString());
  }, [props.pincode]);

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
          {data?.current.merchantList.map((merchant) => (
            <MerchantCard {...merchant} />
          ))}
        </>
      )}
    </>
  );
}
