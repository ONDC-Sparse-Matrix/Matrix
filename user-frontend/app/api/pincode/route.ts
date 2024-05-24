import { MerchantData, PincodeData } from "@/lib/types";

export async function POST(req: Request) {
  const { pincode } = await req.json();
  const res = await fetch(`http://localhost:3001/pincode/${pincode}`)
  const data: PincodeData = await res.json();
  console.log(data);
  return new Response(JSON.stringify(data));
}
