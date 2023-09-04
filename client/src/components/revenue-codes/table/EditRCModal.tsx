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

import {
  Box,
  Button,
  Group,
  Modal,
  SimpleGrid,
  Skeleton,
  Stack,
} from "@mantine/core";
import React from "react";
import { useMutation, useQuery, useQueryClient } from "react-query";
import { notifications } from "@mantine/notifications";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCheck, faXmark } from "@fortawesome/pro-solid-svg-icons";
import { useForm, yupResolver } from "@mantine/form";
import { revenueCodeTableStore } from "@/stores/AccountingStores";
import {
  getGLAccounts,
  getRevenueCodeDetail,
} from "@/services/AccountingRequestService";
import {
  GeneralLedgerAccount,
  RevenueCode,
  RevenueCodeFormValues,
} from "@/types/accounting";
import { TChoiceProps } from "@/types";
import { useFormStyles } from "@/assets/styles/FormStyles";
import axios from "@/helpers/AxiosConfig";
import { APIError } from "@/types/server";
import { revenueCodeSchema } from "@/helpers/schemas/AccountingSchema";
import { ValidatedTextInput } from "@/components/common/fields/TextInput";
import { ValidatedTextArea } from "@/components/common/fields/TextArea";
import { SelectInput } from "@/components/common/fields/SelectInput";

type EditRCModalFormProps = {
  revenueCode: RevenueCode;
  selectGlAccountData: TChoiceProps[];
};

export function EditRCModalForm({
  revenueCode,
  selectGlAccountData,
}: EditRCModalFormProps) {
  const { classes } = useFormStyles();
  const [loading, setLoading] = React.useState<boolean>(false);
  const queryClient = useQueryClient();

  const mutation = useMutation(
    (values: RevenueCodeFormValues) =>
      axios.put(`/revenue_codes/${revenueCode.id}/`, values),
    {
      onSuccess: () => {
        queryClient
          .invalidateQueries({
            queryKey: ["revenue-code-table-data"],
          })
          .then(() => {
            queryClient
              .invalidateQueries({
                queryKey: ["revenueCode", revenueCode?.id],
              })
              .then(() => {
                notifications.show({
                  title: "Success",
                  message: "Revenue Code updated successfully",
                  color: "green",
                  withCloseButton: true,
                  icon: <FontAwesomeIcon icon={faCheck} />,
                });
                revenueCodeTableStore.set("editModalOpen", false);
              });
          });
      },
      onError: (error: any) => {
        const { data } = error.response;
        if (data.type === "validation_error") {
          data.errors.forEach((e: APIError) => {
            form.setFieldError(e.attr, e.detail);
            if (e.attr === "non_field_errors") {
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

  const form = useForm<RevenueCodeFormValues>({
    validate: yupResolver(revenueCodeSchema),
    initialValues: {
      code: revenueCode.code,
      description: revenueCode.description,
      expenseAccount: revenueCode.expenseAccount || "",
      revenueAccount: revenueCode.revenueAccount || "",
    },
  });

  const submitForm = (values: RevenueCodeFormValues) => {
    setLoading(true);
    mutation.mutate(values);
  };

  return (
    <form onSubmit={form.onSubmit((values) => submitForm(values))}>
      <Box className={classes.div}>
        <Box>
          <ValidatedTextInput<RevenueCodeFormValues>
            form={form}
            className={classes.fields}
            name="code"
            label="Code"
            placeholder="Code"
            variant="filled"
            withAsterisk
          />
          <ValidatedTextArea<RevenueCodeFormValues>
            form={form}
            className={classes.fields}
            name="description"
            label="Description"
            placeholder="Description"
            variant="filled"
            withAsterisk
          />
          <SimpleGrid cols={2} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
            <SelectInput<RevenueCodeFormValues>
              form={form}
              data={selectGlAccountData}
              className={classes.fields}
              name="expenseAccount"
              label="Expense Account"
              placeholder="Expense Account"
              variant="filled"
              clearable
            />
            <SelectInput<RevenueCodeFormValues>
              form={form}
              data={selectGlAccountData}
              className={classes.fields}
              name="revenueAccount"
              label="Revenue Account"
              placeholder="Revenue Account"
              variant="filled"
              clearable
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

export function EditRCModal(): React.ReactElement {
  const [showEditModal, setShowEditModal] =
    revenueCodeTableStore.use("editModalOpen");
  const [revenueCode] = revenueCodeTableStore.use("selectedRecord");
  const queryClient = useQueryClient();

  const { data: glAccountData, isLoading: isGLAccountDataLoading } = useQuery({
    queryKey: "gl-account-data",
    queryFn: () => getGLAccounts(),
    enabled: showEditModal,
    initialData: () => queryClient.getQueryData("gl-account"),
    staleTime: Infinity,
  });

  const selectGlAccountData =
    glAccountData?.map((glAccount: GeneralLedgerAccount) => ({
      value: glAccount.id,
      label: glAccount.accountNumber,
    })) || [];

  const { data: revenueCodeData, isLoading: isRevenueCodeDataLoading } =
    useQuery({
      queryKey: ["revenueCode", revenueCode?.id],
      queryFn: () => {
        if (!revenueCode) {
          return Promise.resolve(null);
        }
        return getRevenueCodeDetail(revenueCode.id);
      },
      enabled: showEditModal,
      initialData: () =>
        queryClient.getQueryData(["revenueCode", revenueCode?.id]),
      staleTime: Infinity,
    });

  const isDataLoading = isRevenueCodeDataLoading || isGLAccountDataLoading;

  return (
    <Modal.Root opened={showEditModal} onClose={() => setShowEditModal(false)}>
      <Modal.Overlay />
      <Modal.Content>
        <Modal.Header>
          <Modal.Title>Edit Revenue Code</Modal.Title>
          <Modal.CloseButton />
        </Modal.Header>
        <Modal.Body>
          {isDataLoading ? (
            <Stack>
              <Skeleton height={400} />
            </Stack>
          ) : (
            <>
              {revenueCodeData && (
                <EditRCModalForm
                  revenueCode={revenueCodeData}
                  selectGlAccountData={selectGlAccountData}
                />
              )}
            </>
          )}
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
}
