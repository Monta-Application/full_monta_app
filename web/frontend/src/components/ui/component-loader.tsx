/**
 * COPYRIGHT(c) 2024 Trenova
 *
 * This file is part of Trenova.
 *
 * The Trenova software is licensed under the Business Source License 1.1. You are granted the right
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

import { cn } from "@/lib/utils";
import { faSpinner } from "@fortawesome/pro-duotone-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

export function ComponentLoader({ className }: { className?: string }) {
  return (
    <div
      className={cn("flex flex-col items-center justify-center p-2", className)}
    >
      <FontAwesomeIcon
        icon={faSpinner}
        size="1x"
        className="text-primary motion-safe:animate-spin"
      />
      <p className="text-foreground mt-2 text-sm">Loading data...</p>
      <p className="text-muted-foreground mt-2 text-sm">
        If this takes too long, please refresh the page.
      </p>
    </div>
  );
}
