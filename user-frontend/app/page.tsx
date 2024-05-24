import Image from "next/image";
import { SearchBar } from "@/components/searchbar";
import bg from "@/public/background.svg"

export default function Home() {
  return (
    <div className="h-screen mx-auto flex flex-col justify-center items-center" style={{
      backgroundImage: 'radial-gradient(106.89999999999999% 91.8% at 100% 100%, #3da1ff 0%, #ffffff 100%)',
    }}>
      <h1 className="text-8xl font-extrabold bg-gradient-to-r from-blue-500 to-[#5D26C1] bg-clip-text text-transparent">matriX</h1>
      <p className="mb-8 my-2 text-lg text-center uppercase tracking-widest text-blue-800">Pincode Search System</p>
      <SearchBar />
    </div>
  );
}
