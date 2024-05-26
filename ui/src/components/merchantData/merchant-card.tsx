import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { MerchantData } from "@/lib/types";

export function MerchantCard(props: MerchantData) {
  return (
    <>
      <Card className="drop-shadow-sm mb-4 px-2 transition-all duration-200 hover:shadow-md">
        <CardHeader className="flex flex-row justify-between">
          <CardTitle>{props.name}</CardTitle>
          <CardDescription>{props.email}</CardDescription>
        </CardHeader>
      </Card>
    </>
  );
}
