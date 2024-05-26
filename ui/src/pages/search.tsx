import { SearchBar } from "@/components/searchbar";
import { useSearchParams } from "react-router-dom";
import { Navbar } from "@/components/navbar";
import { Merchants } from "@/components/merchantData/merchants";
import styles from "./pages.module.css";

export default function Search() {
  const [searchParams] = useSearchParams();
  const pincode = searchParams.get("pincode") || "";
  const sessionId = searchParams.get("session") || "";
  return (
    <>
      <Navbar />
      <div className={`h-screen overflow-auto mx-auto max-w-xl px-4 pt-20 ${styles.sleek_scrollbar}`}>
        <div className="sticky">
        <SearchBar pincode={pincode} />
        </div>
        <Merchants pincode={parseInt(pincode)} sessionId={sessionId} />
      </div>
    </>
  );
}
