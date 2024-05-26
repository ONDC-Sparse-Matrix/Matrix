import { SearchBar } from "@/components/searchbar";
import { useSearchParams } from "react-router-dom";
import { Navbar } from "@/components/navbar";
import { Merchants } from "@/components/merchantData/merchants";

export default function Search() {
  const [searchParams] = useSearchParams();
  const pincode = searchParams.get("pincode") || "";
  const sessionId = searchParams.get("session") || "";
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
