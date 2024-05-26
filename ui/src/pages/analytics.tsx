import { Navbar } from "@/components/navbar";
import { MostSearched } from "@/components/analytics/most-searched";
import { BarChart } from "@/components/analytics/search-distribution";

export default function Analytics() {
  return (
    <>
      <Navbar />
      <div className="h-screen max-w-xl mx-auto pt-20">
        <h1 className="text-2xl font-bold pb-4 text-sky-950">Analytics</h1>
        <MostSearched />
        <BarChart
          data={[
            { label: "North", value: 2000 },
            { label: "East", value: 100 },
            { label: "West", value: 1500 },
            { label: "South", value: 500 },
          ]}
        />
        <p className="text-sm text-center mt-4 text-muted-foreground ">
          Number of requests by region
        </p>
      </div>
    </>
  );
}
