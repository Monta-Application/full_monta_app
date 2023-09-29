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
import { useQuery, useQueryClient } from "react-query";
import { useForm, yupResolver } from "@mantine/form";
import { notifications } from "@mantine/notifications";
import { divisionCodeTableStore as store } from "@/stores/AccountingStores";
import { getGLAccounts } from "@/services/AccountingRequestService";
import {
  DivisionCode,
  DivisionCodeFormValues as FormValues,
  GeneralLedgerAccount,
} from "@/types/accounting";
import { TChoiceProps } from "@/types";
import { useFormStyles } from "@/assets/styles/FormStyles";
import { divisionCodeSchema } from "@/lib/schemas/AccountingSchema";
import { SelectInput } from "@/components/common/fields/SelectInput";
import { statusChoices } from "@/lib/constants";
import { ValidatedTextInput } from "@/components/common/fields/TextInput";
import { ValidatedTextArea } from "@/components/common/fields/TextArea";
import { useCustomMutation } from "@/hooks/useCustomMutation";
import { TableStoreProps } from "@/types/tables";

function CreateDCModalForm({
  selectGlAccountData,
}: {
  selectGlAccountData: TChoiceProps[];
}) {
  const { classes } = useFormStyles();
  const [loading, setLoading] = React.useState<boolean>(false);

  const form = useForm<FormValues>({
    validate: yupResolver(divisionCodeSchema),
    initialValues: {
      status: "A",
      code: "",
      description: "",
      apAccount: "",
      cashAccount: "",
      expenseAccount: "",
    },
  });

  const mutation = useCustomMutation<
    FormValues,
    Omit<TableStoreProps<DivisionCode>, "drawerOpen">
  >(
    form,
    store,
    notifications,
    {
      method: "POST",
      path: "/division_codes/",
      successMessage: "Division Code created successfully.",
      queryKeysToInvalidate: ["division-code-table-data"],
      additionalInvalidateQueries: ["divisionCodes"],
      closeModal: true,
      errorMessage: "Failed to create division code.",
    },
    () => setLoading(false),
  );

  const submitForm = (values: FormValues) => {
    setLoading(true);
    mutation.mutate(values);
  };

  return (
    <form onSubmit={form.onSubmit((values) => submitForm(values))}>
      <Box className={classes.div}>
        <Box>
          <SimpleGrid cols={2} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
            <SelectInput<FormValues>
              form={form}
              data={statusChoices}
              name="status"
              label="Status"
              placeholder="Status"
              variant="filled"
              withAsterisk
            />
            <ValidatedTextInput<FormValues>
              form={form}
              name="code"
              label="Code"
              placeholder="Code"
              variant="filled"
              withAsterisk
              maxLength={4}
            />
          </SimpleGrid>
          <ValidatedTextArea<FormValues>
            form={form}
            name="description"
            label="Description"
            placeholder="Description"
            variant="filled"
            withAsterisk
          />
          <SimpleGrid cols={2} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
            <SelectInput<FormValues>
              form={form}
              data={selectGlAccountData}
              name="apAccount"
              label="AP Account"
              placeholder="AP Account"
              variant="filled"
              clearable
            />
            <SelectInput<FormValues>
              form={form}
              data={selectGlAccountData}
              name="cashAccount"
              label="Cash Account"
              placeholder="Cash Account"
              variant="filled"
              clearable
            />
          </SimpleGrid>
          <SelectInput<FormValues>
            form={form}
            data={selectGlAccountData}
            name="expenseAccount"
            label="Expense Account"
            placeholder="Expense Account"
            variant="filled"
            clearable
          />
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

export function CreateDCModal(): React.ReactElement {
  const [showCreateModal, setShowCreateModal] = store.use("createModalOpen");
  const queryClient = useQueryClient();

  const { data: glAccountData, isLoading: isGLAccountDataLoading } = useQuery({
    queryKey: "gl-account-data",
    queryFn: () => getGLAccounts(),
    enabled: showCreateModal,
    initialData: () => queryClient.getQueryData("gl-account"),
    staleTime: Infinity,
  });

  const selectGlAccountData =
    glAccountData?.map((glAccount: GeneralLedgerAccount) => ({
      value: glAccount.id,
      label: glAccount.accountNumber,
    })) || [];

  return (
    <Modal.Root
      opened={showCreateModal}
      onClose={() => setShowCreateModal(false)}
    >
      <Modal.Overlay />
      <Modal.Content>
        <Modal.Header>
          <Modal.Title>Create Division Code</Modal.Title>
          <Modal.CloseButton />
        </Modal.Header>
        <Modal.Body>
          {isGLAccountDataLoading ? (
            <Stack>
              <Skeleton height={400} />
            </Stack>
          ) : (
            <CreateDCModalForm selectGlAccountData={selectGlAccountData} />
          )}
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
}
