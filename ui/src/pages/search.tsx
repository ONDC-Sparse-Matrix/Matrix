import { SearchBar } from "@/components/searchbar";
import { useSearchParams } from "react-router-dom";
import { Navbar } from "@/components/navbar";
import { Merchants } from "@/components/merchantData/merchants";
import { useEffect } from "react";

export default function Search() {
  const [searchParams] = useSearchParams();
  const pincode = searchParams.get("pincode") || "";
  const sessionId = searchParams.get("session") || "";

  useEffect(()=>{
    let eventSource:EventSource|undefined;
    if(sessionId){
      eventSource = new EventSource(`http://localhost:3001/sse/${sessionId}`);
      eventSource.onmessage = (e)=>{
        console.log(e.data);
        //TODO: @Garv save this cache
      }

      return ()=>{
        if(eventSource){
          eventSource.close();
        }
      }
    }

  },[sessionId,pincode])
  return (
    <>
      <Navbar />
      <div className="h-screen overflow-auto mx-auto max-w-xl pt-20">
        <SearchBar pincode={pincode} />
        <Merchants pincode={parseInt(pincode)} sessionId={sessionId} />
      </div>
    </>
  );
}
