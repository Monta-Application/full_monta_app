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

import { StatusChoiceProps } from "@/types/index";
import {
  AccountClassificationChoiceProps,
  AccountSubTypeChoiceProps,
  AccountTypeChoiceProps,
  CashFlowTypeChoiceProps,
} from "@/lib/choices";
import { BaseModel } from "@/types/organization";

/** Types for Division Codes */
export interface DivisionCode extends BaseModel {
  id: string;
  status: StatusChoiceProps;
  code: string;
  description: string;
  apAccount?: string | null;
  cashAccount?: string | null;
  expenseAccount?: string | null;
}

export type DivisionCodeFormValues = Omit<
  DivisionCode,
  "id" | "organization" | "created" | "modified"
>;

/** Types for General Ledger Accounts */
export interface GeneralLedgerAccount extends BaseModel {
  id: string;
  status: StatusChoiceProps;
  accountNumber: string;
  description: string;
  accountType: AccountTypeChoiceProps | "";
  cashFlowType?: CashFlowTypeChoiceProps | "" | null;
  accountSubType?: AccountSubTypeChoiceProps | "" | null;
  accountClassification?: AccountClassificationChoiceProps | "" | null;
}

export type GLAccountFormValues = Omit<
  GeneralLedgerAccount,
  "organization" | "created" | "modified" | "id"
>;

/** Types for Revenue Codes */
export interface RevenueCode extends BaseModel {
  id: string;
  code: string;
  description: string;
  expenseAccount?: string | null;
  revenueAccount?: string | null;
  revAccountNum?: string | null;
  expAccountNum?: string | null;
}

export type RevenueCodeFormValues = Omit<
  RevenueCode,
  | "id"
  | "organization"
  | "revAccountNum"
  | "expAccountNum"
  | "created"
  | "modified"
>;
