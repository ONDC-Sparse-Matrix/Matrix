import { Clock, AlertCircleIcon } from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";

interface TimeDataProps {
  success: boolean;
  pincode: number;
  time: number;
}

export function TimeData(props: TimeDataProps) {
  return (
    <>
      {props.success ? (
        <>
          <Alert className="mb-8 shadow-inner">
            <Clock className="h-4 w-4" />
            <AlertTitle className="font-bold">Heads up!</AlertTitle>
            <AlertDescription>
              Time taken for request:{" "}
              <span className="text-blue-500 font-bold">
                {props.time / 1000}s
              </span>
            </AlertDescription>
          </Alert>
          <h1 className="text-lg uppercase tracking-widest font-bold text-center text-muted-foreground">
            Search Results for{" "}
            <span className="text-foreground font-extrabold">
              {props.pincode}
            </span>
          </h1>
        </>
      ) : (
        <Alert className="mb-8" variant={"destructive"}>
          <AlertCircleIcon className="h-4 w-4" />
          <AlertTitle className="font-bold">
            No Merchants Available :(
          </AlertTitle>
          <AlertDescription>
            There seems to be no merchants serving at {props.pincode}
          </AlertDescription>
        </Alert>
      )}
    </>
  );
}
