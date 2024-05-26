import { AddMerchantCard } from "@/components/add-merchant-card";
import { Navbar } from "@/components/navbar";

export default function AddMerchant() {
  return (
    <>
      <Navbar />
      <div className="h-screen w-full flex items-center justify-center">
        <AddMerchantCard />
      </div>
    </>
  );
}
