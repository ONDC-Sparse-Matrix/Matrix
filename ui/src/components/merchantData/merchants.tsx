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
import styles from "../components.module.css";
import { MerchantPagination } from "./merchant-pagination";

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

  const itemsPerPage: number = 4;
  const [currentPage, setCurrentPage] = useState<number>(1);
  //TODO: Ye abhi slice kr rakha hai data - gotta transfer it to backend
  const totalItems = data?.merchantList.length || 0;
  const totalPages = Math.ceil(totalItems / itemsPerPage);
  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentItems = data?.merchantList.slice(
    indexOfFirstItem,
    indexOfLastItem
  );

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
      let getData;
      getData = await new Promise<PincodeData | null>((resolve, reject) => {
        axios.post<PincodeData>(`${FETCH_PINCODE_DATA}/${pincode}`, {
          session: props.sessionId
        })
        .then((res) => {
          resolve(res.data)
        }).catch((err) => {
          reject(null);
        })
      })
      setData(getData?.current)

      console.log('getting this data from request and set it up: ', getData)

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

          <div
            className={`h-[50vh] my-4 border-b px-4 overflow-auto ${styles.sleek_scrollbar}`}
          >
            {currentItems?.map((merchant, index) => (
              <MerchantCard key={index} {...merchant} />
            ))}
          </div>
          <MerchantPagination
            currentPage={currentPage}
            setCurrentPage={setCurrentPage}
            totalPages={totalPages}
          />
        </>
      )}
    </>
  );
}
