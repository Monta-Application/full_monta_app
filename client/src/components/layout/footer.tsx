/*
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

import { Button } from "../ui/button";

export function Footer() {
  return (
    <footer className="mt-10 flex h-10 items-center justify-between border-t px-10 py-2 font-semibold">
      {/* Copyright */}
      <div className="text-foreground text-xs">
        <span className="me-1">&copy; {new Date().getFullYear()}</span>
        <a href="#" target="_blank" rel="noopener noreferrer">
          Trenova Technologies
        </a>
      </div>

      <Button
        size="xs"
        className="hover:bg-muted inline-flex gap-x-1.5 bg-transparent px-2.5 py-0.5 text-xs font-bold text-blue-600 transition-colors"
      >
        <span className="relative mb-0.5 flex h-2 w-2">
          <span className="absolute inline-flex h-full w-full rounded-full bg-blue-400 opacity-75 motion-safe:animate-ping"></span>
          <span className="relative inline-flex h-2 w-2 rounded-full bg-blue-600"></span>
        </span>
        All Systems Operational
      </Button>
    </footer>
  );
}
