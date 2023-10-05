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

import { Button, Drawer, Group } from "@mantine/core";
import React from "react";
import { notifications } from "@mantine/notifications";
import { useForm, yupResolver } from "@mantine/form";
import { jobTitleTableStore as store } from "@/stores/UserTableStore";
import { JobTitle, JobTitleFormValues as FormValues } from "@/types/accounts";
import { useFormStyles } from "@/assets/styles/FormStyles";
import { jobTitleSchema } from "@/lib/schemas/AccountsSchema";
import { useCustomMutation } from "@/hooks/useCustomMutation";
import { TableStoreProps } from "@/types/tables";
import { JobTitleForm } from "./CreateJobTitleModal";

export function EditJobTitleModalForm({
  jobTitle,
  onCancel,
}: {
  jobTitle: JobTitle;
  onCancel: () => void;
}) {
  const { classes } = useFormStyles();
  const [loading, setLoading] = React.useState<boolean>(false);

  const form = useForm<FormValues>({
    validate: yupResolver(jobTitleSchema),
    initialValues: {
      status: jobTitle.status,
      name: jobTitle.name,
      description: jobTitle.description || "",
      jobFunction: jobTitle.jobFunction,
    },
  });

  const mutation = useCustomMutation<FormValues, TableStoreProps<JobTitle>>(
    form,
    notifications,
    {
      method: "PUT",
      path: `/job_titles/${jobTitle.id}/`,
      successMessage: "Job Title updated successfully.",
      queryKeysToInvalidate: ["job-title-table-data"],
      additionalInvalidateQueries: ["jobTitles"],
      closeModal: true,
      errorMessage: "Failed to update job title.",
    },
    () => setLoading(false),
    store,
  );

  const submitForm = (values: FormValues) => {
    setLoading(true);
    mutation.mutate(values);
  };

  return (
    <form onSubmit={form.onSubmit((values) => submitForm(values))}>
      <JobTitleForm form={form} />
      <Group position="right" mt="md">
        <Button
          variant="subtle"
          onClick={onCancel}
          color="gray"
          type="button"
          className={classes.control}
        >
          Cancel
        </Button>
        <Button
          color="white"
          type="submit"
          className={classes.control}
          loading={loading}
        >
          Submit
        </Button>
      </Group>
    </form>
  );
}

export function JobTitleDrawer(): React.ReactElement {
  const [drawerOpen, setDrawerOpen] = store.use("drawerOpen");
  const [jobTitle] = store.use("selectedRecord");
  const onCancel = () => setDrawerOpen(false);

  return (
    <Drawer.Root
      position="right"
      opened={drawerOpen}
      onClose={() => setDrawerOpen(false)}
    >
      <Drawer.Overlay />
      <Drawer.Content>
        <Drawer.Header>
          <Drawer.Title>
            Edit Job Title: {jobTitle && jobTitle.name}
          </Drawer.Title>
          <Drawer.CloseButton />
        </Drawer.Header>
        <Drawer.Body>
          {jobTitle && (
            <EditJobTitleModalForm jobTitle={jobTitle} onCancel={onCancel} />
          )}
        </Drawer.Body>
      </Drawer.Content>
    </Drawer.Root>
  );
}
