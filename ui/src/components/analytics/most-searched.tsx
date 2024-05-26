import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export function MostSearched() {
  return (
    <>
      <div className="w-full flex gap-2 mb-8">
        <Card className="flex-1">
          <CardHeader>
            <CardTitle className="text-muted-foreground text-sm">Most Searched Region</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-center text-2xl font-semibold text-blue-500">North</div>
          </CardContent>
        </Card>

        <Card className="flex-1">
          <CardHeader>
            <CardTitle className="text-muted-foreground text-sm">Most Searched Pincode</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="text-center text-2xl font-semibold text-blue-500">100000</div>
          </CardContent>
        </Card>
      </div>
    </>
  );
}
