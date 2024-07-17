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



import { useUserOrganization } from "@/hooks/useQueries";
import { faChevronDown } from "@fortawesome/pro-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { Skeleton } from "../ui/skeleton";

export function OrganizationLogo() {
  const { data, isLoading } = useUserOrganization();

  if (isLoading) {
    return <Skeleton className="h-14" />;
  }

  const initial = data?.name?.charAt(0);

  const cityState = `${data?.city}, ${data?.state?.abbreviation}`;

  return (
    <div className="hover:bg-muted group col-span-full flex w-full items-center gap-x-4 rounded-lg p-1 hover:cursor-pointer">
      <Avatar className="bg-muted/50 group-hover:bg-muted-foreground/20 border-muted size-14 flex-none rounded-lg border">
        <AvatarImage
          src={data?.logoUrl || ""}
          alt={"Trenova Logo"}
          className="size-14 flex-none rounded-lg p-2"
        />
        <AvatarFallback className="rounded-none" delayMs={600}>
          {initial}
        </AvatarFallback>
      </Avatar>
      <div className="flex flex-1 flex-col">
        <div className="flex items-center">
          <h2 className="text-foreground mr-2 w-36 truncate text-lg font-semibold leading-7">
            {data?.name || ""}
          </h2>
          <FontAwesomeIcon
            icon={faChevronDown}
            className="text-muted-foreground size-3"
          />
        </div>
        <p className="text-muted-foreground text-sm">{cityState}</p>
      </div>
    </div>
  );
}

export function MiniOrganizationLogo() {
  const { data, isLoading } = useUserOrganization();

  if (isLoading) {
    return <Skeleton className="h-14" />;
  }

  const initial = data?.name?.charAt(0);

  return (
    <div className="hover:bg-muted rounded-lghover:cursor-pointer group col-span-full flex w-full items-center gap-x-4">
      <Avatar className="bg-muted/50 group-hover:bg-muted-foreground/20 border-muted size-9 flex-none rounded-lg border">
        <AvatarImage
          src={data?.logoUrl || ""}
          alt={"Trenova Logo"}
          className="size-9 flex-none rounded-lg p-2"
        />
        <AvatarFallback className="rounded-none" delayMs={600}>
          {initial}
        </AvatarFallback>
      </Avatar>
    </div>
  );
}
