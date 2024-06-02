import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { ComponentLoader } from "@/components/ui/component-loader";
import { useTheme } from "@/components/ui/theme-provider";
import { useDailyShipmentCounts } from "@/hooks/useQueries";
import {
  getDateNDaysAgo,
  getDaysBetweenDates,
  getMonthDayString,
  getTodayDate,
} from "@/lib/date";
import { faChartSimple } from "@fortawesome/pro-duotone-svg-icons";
import {
  faArrowDownArrowUp,
  faClock,
} from "@fortawesome/pro-regular-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Point, ResponsiveLine } from "@nivo/line";
import { useMemo } from "react";

function LineChartTooltip({ point }: { point: Point }) {
  return (
    <div className="bg-background border-border rounded-lg border p-2 shadow-md">
      <p className="border-border border-b text-sm font-semibold">
        {point.data.xFormatted}
      </p>
      <div className="flex items-center text-sm">
        <div className="mr-2 size-3 rounded-full bg-green-600" />
        Total:
        <span className="ml-1 text-sm font-semibold">
          {point.data.yFormatted}
        </span>
      </div>
    </div>
  );
}

function LineChart({ data }: { data: any }) {
  const { theme } = useTheme();

  return (
    <div className="h-[15vh]">
      <ResponsiveLine
        data={data}
        margin={{ top: 10 }}
        xScale={{ type: "point" }}
        yFormat=" >-.2f"
        axisTop={null}
        axisRight={null}
        axisBottom={null}
        axisLeft={null}
        pointSize={10}
        pointColor={{ theme: "background" }}
        pointBorderWidth={2}
        pointBorderColor={{ from: "serieColor" }}
        pointLabel="data.yFormatted"
        pointLabelYOffset={-12}
        enableTouchCrosshair={false}
        animate
        colors={["#028ee6", "#774dd7"]}
        curve="natural"
        useMesh={true}
        theme={{
          crosshair: {
            line: {
              stroke: theme === "dark" ? "#fff" : "#000",
            },
          },
        }}
        enableArea={true}
        enableGridX={false}
        enableGridY={false}
        tooltip={({ point }) => {
          return <LineChartTooltip point={point} />;
        }}
      />
    </div>
  );
}

export default function DailyShipmentCounts() {
  const startDate = getDateNDaysAgo(7);
  const endDate = getTodayDate();

  const { formattedData, data, isLoading, isError } = useDailyShipmentCounts(
    startDate,
    endDate,
  );

  // Check if there is actual data to display in the chart
  const hasChartData = useMemo(
    () => formattedData[0]?.data.length > 0,
    [formattedData],
  );

  // Compute the display message for dates
  const dateDisplay = useMemo(() => {
    return `${getDaysBetweenDates(
      startDate,
      endDate,
    )} days (${getMonthDayString(startDate)} - ${getMonthDayString(endDate)})`;
  }, [startDate, endDate]);

  if (isError) {
    return (
      <Card className="relative col-span-2">
        <CardContent className="p-0">
          <div className="flex h-[40vh] items-center justify-center">
            <p className="text-muted-foreground">
              Unable to fetch data. Please try again later.
            </p>
          </div>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card className="col-span-4 lg:col-span-2">
      {isLoading ? (
        <ComponentLoader className="h-[40vh]" />
      ) : (
        <CardContent className="relative p-0">
          <div className="border-border flex items-start justify-between border-b border-dashed p-4">
            <div>
              <div className="flex items-center">
                <p className="text-2xl font-bold">{data?.count || "--"}</p>
                <span className="text-muted-foreground ml-2 text-xs font-normal">
                  {dateDisplay}
                </span>
              </div>
              <h2 className="text-muted-foreground font-semibold">
                Daily Shipment Count
              </h2>
            </div>
            <Button
              size="sm"
              variant="outline"
              className="absolute right-52 top-4"
            >
              <FontAwesomeIcon icon={faArrowDownArrowUp} className="mr-2" />
              Sorted by
              <Badge withDot={false} variant="purple" className="ml-2">
                Created At
              </Badge>
            </Button>
            <Button
              size="sm"
              variant="outline"
              className="absolute right-4 top-4"
            >
              <FontAwesomeIcon icon={faClock} className="mr-2" />
              Last
              <Badge withDot={false} variant="info" className="ml-2">
                Last 30 days
              </Badge>
            </Button>
          </div>
          {hasChartData ? (
            <LineChart data={formattedData} />
          ) : (
            <div className="bg-muted/50 border-border m-5 flex h-[30vh] flex-col items-center justify-center rounded-md border">
              <FontAwesomeIcon icon={faChartSimple} className="mb-2 text-2xl" />
              <h3 className="text-foreground font-semibold">No data to show</h3>
              <p className="text-muted-foreground">
                May be due to lack of shipments.
              </p>
            </div>
          )}
        </CardContent>
      )}
    </Card>
  );
}
