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

import { useQuery, useQueryClient } from "react-query";
import { Card, Divider, Skeleton, Text } from "@mantine/core";
import React from "react";
import { getInvoiceControl } from "@/services/OrganizationRequestService";
import { usePageStyles } from "@/assets/styles/PageStyles";
import { InvoiceControlForm } from "@/components/control-files/_partials/InvoiceControlForm";

function InvoiceControlPage() {
  const { classes } = usePageStyles();
  const queryClient = useQueryClient();

  const { data: invoiceControlData, isLoading: isInvoiceControlDataLoading } =
    useQuery({
      queryKey: ["invoiceControl"],
      queryFn: () => getInvoiceControl(),
      initialData: () => queryClient.getQueryData(["invoiceControl"]),
      staleTime: Infinity,
    });

  // Store first element of invoiceControlData in variable
  const invoiceControlDataArray = invoiceControlData?.[0];

  return isInvoiceControlDataLoading ? (
    <Skeleton height={400} />
  ) : (
    <Card className={classes.card}>
      <Text fz="xl" fw={700} className={classes.text}>
        Invoice Controls
      </Text>

      <Divider my={10} />
      {invoiceControlDataArray && (
        <InvoiceControlForm invoiceControl={invoiceControlDataArray} />
      )}
    </Card>
  );
}

export default InvoiceControlPage;
