import { SearchBar } from "@/components/searchbar";
import { useSearchParams } from "react-router-dom";
import { Navbar } from "@/components/navbar";
import { Merchants } from "@/components/merchantData/merchants";

export default function Search() {
  const [searchParams] = useSearchParams();
  const pincode = searchParams.get("pincode") || "";
  return (
    <>
      <Navbar />
      <div className="h-screen overflow-auto mx-auto max-w-xl pt-20">
        <SearchBar pincode={pincode} />
        <Merchants pincode={parseInt(pincode)} />
      </div>
    </>
  );
}
