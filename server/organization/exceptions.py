# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2024 MONTA                                                                         -
#                                                                                                  -
#  This file is part of Monta.                                                                     -
#                                                                                                  -
#  The Monta software is licensed under the Business Source License 1.1. You are granted the right -
#  to copy, modify, and redistribute the software, but only for non-production use or with a total -
#  of less than three server instances. Starting from the Change Date (November 16, 2026), the     -
#  software will be made available under version 2 or later of the GNU General Public License.     -
#  If you use the software in violation of this license, your rights under the license will be     -
#  terminated automatically. The software is provided "as is," and the Licensor disclaims all      -
#  warranties and conditions. If you use this license's text or the "Business Source License" name -
#  and trademark, you must comply with the Licensor's covenants, which include specifying the      -
#  Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use     -
#  Grant, and not modifying the license in any other way.                                          -
# --------------------------------------------------------------------------------------------------


class CacheManagerError(Exception):
    """Base exception class for CacheManager-related errors."""

    pass


class CacheKeyError(CacheManagerError):
    """Raised when there is an issue with a cache key."""

    pass


class CacheConnectionError(CacheManagerError):
    """Raised when there is an issue connecting to the cache."""

    pass


class CacheOperationError(CacheManagerError):
    """Raised when there is an issue performing a cache operation."""

    pass


class InvalidEmailProtocol(Exception):
    """Raised when an invalid email protocol is used."""

    pass


class InvalidEmailProfile(Exception):
    """Raised when an invalid email profile is used."""

    pass


class ConditionalStructureError(Exception):
    """
    Exception raised when the structure of the conditional logic is not valid
    """

    pass


class InvalidOperationError(Exception):
    """
    Exception raised when the conditional logic operation is not in the list of available operations
    """

    pass
