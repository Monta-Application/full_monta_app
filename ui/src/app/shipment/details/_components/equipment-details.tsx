import { AsyncSelectField } from "@/components/fields/async-select";
import { InputField } from "@/components/fields/input-field";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { FormControl, FormGroup } from "@/components/ui/form";
import { Icon } from "@/components/ui/icons";
import { ShipmentSchema } from "@/lib/schemas/shipment-schema";
import { faTruck } from "@fortawesome/pro-solid-svg-icons";
import { useFormContext } from "react-hook-form";

export default function EquipmentDetails() {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-x-2">
          <div className="border-border flex size-10 items-center justify-center rounded-lg border">
            <Icon icon={faTruck} className="size-5" />
          </div>
          <div className="flex flex-col gap-0.5">
            <h3 className="text-lg font-semibold">Equipment Details</h3>
            <p className="text-muted-foreground text-xs font-normal">
              Provides a breakdown of equipment details for the shipment.
            </p>
          </div>
        </CardTitle>
      </CardHeader>
      <CardContent>
        <EquipmentForm />
      </CardContent>
    </Card>
  );
}

function EquipmentForm() {
  const { control } = useFormContext<ShipmentSchema>();

  return (
    <FormGroup cols={2} className="gap-4">
      <FormControl>
        <AsyncSelectField
          control={control}
          name="tractorTypeId"
          label="Tractor Type"
          link="/equipment-types/"
          valueKey={["id"]}
          placeholder="Select Tractor Type"
          description="Select the type of tractor used, considering any special requirements (e.g., refrigeration)."
          hasPopoutWindow
          hasPermission
          popoutLink="/equipment/equipment-types/"
          popoutLinkLabel="Equipment Type"
        />
      </FormControl>
      <FormControl>
        <AsyncSelectField
          control={control}
          name="trailerTypeId"
          label="Trailer Type"
          link="/equipment-types/"
          valueKey={["id"]}
          placeholder="Select Trailer Type"
          description="Select the type of trailer used, considering any special requirements (e.g., refrigeration)."
          hasPopoutWindow
          hasPermission
          popoutLink="/equipment/equipment-types/"
          popoutLinkLabel="Equipment Type"
        />
      </FormControl>
      <FormControl>
        <InputField
          name="temperatureMin"
          control={control}
          type="number"
          label="Temperature Min."
          placeholder="Enter Minimum Temperature"
          description="Enter the minimum temperature required during transport."
        />
      </FormControl>
      <FormControl>
        <InputField
          name="temperatureMax"
          control={control}
          type="number"
          label="Temperature Max."
          placeholder="Enter Maximum Temperature"
          description="Enter the maximum temperature allowed during transport."
        />
      </FormControl>
    </FormGroup>
  );
}
