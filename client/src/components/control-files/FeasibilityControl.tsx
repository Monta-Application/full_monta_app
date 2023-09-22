/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * The Monta software is licensed under the Business Source License 1.1. You are granted the right
 * to copy, modify, and redistribute the software, but only for non-production use or with a total
 * of less than three server instances. Starting from the Change Date (November 16, 2026), the
 * software will be made available under version 2 or later of the GNU General Public License.
 * If you use the software in violation of this license, your rights under the license will be
 * terminated automatically. The software is provided "as is," and the Licensor disclaims all
 * warranties and conditions. If you use this license's text or the "Business Source License" name
 * and trademark, you must comply with the Licensor's covenants, which include specifying the
 * Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
 * Grant, and not modifying the license in any other way.
 */

import { useMutation, useQueryClient } from "react-query";
import {
  Box,
  Button,
  Card,
  Divider,
  Group,
  SimpleGrid,
  Skeleton,
  Text,
} from "@mantine/core";
import React from "react";
import { notifications } from "@mantine/notifications";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCheck, faXmark } from "@fortawesome/pro-solid-svg-icons";
import { useForm, yupResolver } from "@mantine/form";
import { usePageStyles } from "@/assets/styles/PageStyles";
import { useFormStyles } from "@/assets/styles/FormStyles";
import axios from "@/helpers/AxiosConfig";
import { APIError } from "@/types/server";
import { emailControlSchema } from "@/helpers/schemas/OrganizationSchema";
import {
  FeasibilityToolControl,
  FeasibilityToolControlFormValues as FormValues,
} from "@/types/dispatch";
import { SelectInput } from "@/components/common/fields/SelectInput";
import { feasibilityOperatorChoices } from "@/helpers/choices";
import { useFeasibilityControl } from "@/hooks/useFeasibilityControl";
import { ValidatedNumberInput } from "@/components/common/fields/NumberInput";

interface Props {
  feasibilityControl: FeasibilityToolControl;
}

function FeasibilityControlForm({
  feasibilityControl,
}: Props): React.ReactElement {
  const { classes } = useFormStyles();
  const [loading, setLoading] = React.useState<boolean>(false);
  const queryClient = useQueryClient();

  const mutation = useMutation(
    (values: FormValues) =>
      axios.put(`/feasibility_tool_control/${feasibilityControl.id}/`, values),
    {
      onSuccess: () => {
        queryClient
          .invalidateQueries({
            queryKey: ["feasibilityControl"],
          })
          .then(() => {
            notifications.show({
              title: "Success",
              message: "Feasibility Control updated successfully",
              color: "green",
              withCloseButton: true,
              icon: <FontAwesomeIcon icon={faCheck} />,
            });
          });
      },
      onError: (error: any) => {
        const { data } = error.response;
        if (data.type === "validation_error") {
          data.errors.forEach((e: APIError) => {
            form.setFieldError(e.attr, e.detail);
            if (e.attr === "nonFieldErrors") {
              notifications.show({
                title: "Error",
                message: e.detail,
                color: "red",
                withCloseButton: true,
                icon: <FontAwesomeIcon icon={faXmark} />,
                autoClose: 10_000, // 10 seconds
              });
            }
          });
        }
      },
      onSettled: () => {
        setLoading(false);
      },
    },
  );

  const form = useForm<FormValues>({
    validate: yupResolver(emailControlSchema),
    initialValues: {
      mpwCriteria: feasibilityControl.mpwCriteria,
      mpwOperator: feasibilityControl.mpwOperator,
      mpdCriteria: feasibilityControl.mpdCriteria,
      mpdOperator: feasibilityControl.mpdOperator,
      mpgCriteria: feasibilityControl.mpgCriteria,
      mpgOperator: feasibilityControl.mpgOperator,
      otpCriteria: feasibilityControl.otpCriteria,
      otpOperator: feasibilityControl.otpOperator,
    },
  });

  const handleSubmit = (values: FormValues) => {
    setLoading(true);
    mutation.mutate(values);
  };

  return (
    <form onSubmit={form.onSubmit((values) => handleSubmit(values))}>
      <Box className={classes.div}>
        <Box>
          <SimpleGrid cols={4} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
            <SelectInput<FormValues>
              form={form}
              name="mpwOperator"
              label="MPW Operator"
              description="Select the operator for MPW"
              data={feasibilityOperatorChoices}
              withAsterisk
            />
            <ValidatedNumberInput<FormValues>
              form={form}
              name="mpwCriteria"
              label="MPW Criteria"
              description="Enter Miles Per Week criteria"
              withAsterisk
            />
            <SelectInput<FormValues>
              form={form}
              name="mpdOperator"
              label="MPD Operator"
              description="Select the operator for MPD"
              data={feasibilityOperatorChoices}
              withAsterisk
            />
            <ValidatedNumberInput<FormValues>
              form={form}
              name="mpdCriteria"
              label="MPD Criteria"
              description="Enter Miles Per Day criteria"
              withAsterisk
            />
            <SelectInput<FormValues>
              form={form}
              name="mpgOperator"
              label="MPG Operator"
              description="Select the operator for MPG"
              data={feasibilityOperatorChoices}
              withAsterisk
            />
            <ValidatedNumberInput<FormValues>
              form={form}
              name="mpgCriteria"
              label="MPG Criteria"
              description="Enter Miles Per Gallon criteria"
              withAsterisk
            />
            <SelectInput<FormValues>
              form={form}
              name="otpOperator"
              label="OTP Operator"
              description="Select the operator for OTP"
              data={feasibilityOperatorChoices}
              withAsterisk
            />
            <ValidatedNumberInput<FormValues>
              form={form}
              name="otpCriteria"
              label="OTP Criteria"
              description="Enter On-Time Percentage Criteria"
              withAsterisk
            />
          </SimpleGrid>
          <Group position="right" mt="md">
            <Button
              color="white"
              type="submit"
              className={classes.control}
              loading={loading}
            >
              Submit
            </Button>
          </Group>
        </Box>
      </Box>
    </form>
  );
}

export default function FeasibilityControlPage() {
  const { classes } = usePageStyles();

  const {
    data: feasibilityControlData,
    isLoading: isFeasibilityControlDataLoading,
  } = useFeasibilityControl();

  const feasibilityControlDataArray = feasibilityControlData?.[0];

  return isFeasibilityControlDataLoading ? (
    <Skeleton height={400} />
  ) : (
    <Card
      className={classes.card}
      sx={{
        overflow: "visible",
      }}
    >
      <Text fz="xl" fw={700} className={classes.text}>
        Feasibility Tool Controls
      </Text>

      <Divider my={10} />
      {feasibilityControlDataArray && (
        <FeasibilityControlForm
          feasibilityControl={feasibilityControlDataArray}
        />
      )}
    </Card>
  );
}
