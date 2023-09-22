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

import { Box, Button, Group, Modal, SimpleGrid, Skeleton } from "@mantine/core";
import React, { Suspense } from "react";
import { useMutation, useQueryClient } from "react-query";
import { notifications } from "@mantine/notifications";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCheck, faXmark } from "@fortawesome/pro-solid-svg-icons";
import { useForm, yupResolver } from "@mantine/form";
import { generalLedgerTableStore as store } from "@/stores/AccountingStores";
import { useFormStyles } from "@/assets/styles/FormStyles";
import { GLAccountFormValues } from "@/types/accounting";
import axios from "@/helpers/AxiosConfig";
import { APIError } from "@/types/server";
import { glAccountSchema } from "@/helpers/schemas/AccountingSchema";
import { SelectInput } from "@/components/common/fields/SelectInput";
import { statusChoices } from "@/helpers/constants";
import { ValidatedTextInput } from "@/components/common/fields/TextInput";
import { ValidatedTextArea } from "@/components/common/fields/TextArea";
import {
  accountClassificationChoices,
  accountSubTypeChoices,
  accountTypeChoices,
  cashFlowTypeChoices,
} from "@/helpers/choices";

function CreateGLAccountModalForm(): React.ReactElement {
  const { classes } = useFormStyles();
  const [loading, setLoading] = React.useState<boolean>(false);
  const queryClient = useQueryClient();

  const mutation = useMutation(
    (values: GLAccountFormValues) => axios.post("/gl_accounts/", values),
    {
      onSuccess: () => {
        queryClient
          .invalidateQueries({
            queryKey: ["gl-account-table-data"],
          })
          .then(() => {
            notifications.show({
              title: "Success",
              message: "GL Account created successfully",
              color: "green",
              withCloseButton: true,
              icon: <FontAwesomeIcon icon={faCheck} />,
            });
            store.set("createModalOpen", false);
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

  const form = useForm<GLAccountFormValues>({
    validate: yupResolver(glAccountSchema),
    initialValues: {
      status: "A",
      accountNumber: "0000-0000-0000-0000",
      description: "",
      accountType: "",
      cashFlowType: "",
      accountSubType: "",
      accountClassification: "",
    },
  });

  const submitForm = (values: GLAccountFormValues) => {
    setLoading(true);
    mutation.mutate(values);
  };

  return (
    <form onSubmit={form.onSubmit((values) => submitForm(values))}>
      <Box className={classes.div}>
        <SimpleGrid cols={2} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
          <SelectInput<GLAccountFormValues>
            form={form}
            data={statusChoices}
            className={classes.fields}
            name="status"
            label="Status"
            placeholder="Status"
            variant="filled"
            withAsterisk
          />
          <ValidatedTextInput<GLAccountFormValues>
            form={form}
            className={classes.fields}
            name="accountNumber"
            label="Account Number"
            placeholder="Account Number"
            variant="filled"
            withAsterisk
          />
        </SimpleGrid>
        <ValidatedTextArea<GLAccountFormValues>
          form={form}
          className={classes.fields}
          name="description"
          label="Description"
          placeholder="Description"
          variant="filled"
          withAsterisk
        />
        <SimpleGrid cols={2} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
          <SelectInput<GLAccountFormValues>
            form={form}
            data={accountTypeChoices}
            className={classes.fields}
            name="accountType"
            label="Account Type"
            placeholder="AP Account"
            variant="filled"
            withAsterisk
            clearable
          />
          <SelectInput<GLAccountFormValues>
            form={form}
            data={cashFlowTypeChoices}
            className={classes.fields}
            name="cashFlowType"
            label="Cash Flow Type"
            placeholder="Cash Flow Type"
            variant="filled"
            clearable
          />
        </SimpleGrid>
        <SimpleGrid cols={2} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
          <SelectInput<GLAccountFormValues>
            form={form}
            data={accountSubTypeChoices}
            className={classes.fields}
            name="accountSubType"
            label="Account Sub Type"
            placeholder="Account Sub Type"
            variant="filled"
            clearable
          />
          <SelectInput<GLAccountFormValues>
            form={form}
            data={accountClassificationChoices}
            className={classes.fields}
            name="accountClassification"
            label="Account Classification"
            placeholder="Account Classification"
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
    </form>
  );
}

export function CreateGLAccountModal(): React.ReactElement {
  const [showCreateModal, setShowCreateModal] = store.use("createModalOpen");

  return (
    <Modal.Root
      opened={showCreateModal}
      onClose={() => setShowCreateModal(false)}
    >
      <Modal.Overlay />
      <Modal.Content>
        <Modal.Header>
          <Modal.Title>Create GL Account</Modal.Title>
          <Modal.CloseButton />
        </Modal.Header>
        <Modal.Body>
          <Suspense fallback={<Skeleton height={400} />}>
            <CreateGLAccountModalForm />
          </Suspense>
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
}
