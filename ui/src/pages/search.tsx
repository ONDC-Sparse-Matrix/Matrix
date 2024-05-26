import { SearchBar } from "@/components/searchbar";
import { useSearchParams } from "react-router-dom";
import { Navbar } from "@/components/navbar";
import { Merchants } from "@/components/merchantData/merchants";
import { useEffect } from "react";
import { updateCache } from "@/lib/db";

export default function Search() {
  const [searchParams] = useSearchParams();
  const pincode = searchParams.get("pincode") || "";
  const sessionId = searchParams.get("session") || "";

  useEffect(()=>{
    console.log("useEffect called");
    let socket:WebSocket|undefined;
    if(sessionId){
      socket = new WebSocket(`ws://localhost:3001/sse/${sessionId}`);
      socket.onopen = ()=>{
        console.log("connection opened");
      }
      socket.onmessage = (e:MessageEvent<string>)=>{
        const cache = JSON.parse(e.data);
        updateCache(cache);
      }
      socket.onclose = ()=>{
        console.log("connection closed");
      }

      return ()=>{
          // socket?.close();
     }
    }

  },[sessionId,pincode])
  return (
    <>
      <Navbar />
      <div className="h-screen mx-auto max-w-xl px-4 pt-20">
        <SearchBar pincode={pincode} />
        <Merchants pincode={parseInt(pincode)} sessionId={sessionId} />
      </div>
    </>
  );
}
