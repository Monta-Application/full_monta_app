"""
COPYRIGHT 2022 MONTA

This file is part of Monta.

Monta is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Monta is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Monta.  If not, see <https://www.gnu.org/licenses/>.
"""

from typing import Any, OrderedDict

from rest_framework import serializers


class ValidatedSerializer(serializers.ModelSerializer):
    """
    Serializer to enforce calling full_clean() on the serializer
    """

    def validate(self, attrs: OrderedDict[str, Any]) -> dict[str, Any]:
        """
        Validate the serializer
        """

        if self.instance is None:
            instance = self.Meta.model(**attrs)  # type: ignore
        else:
            instance = self.instance
            for k, v in attrs.items():
                setattr(self.instance, k, v)
        instance.full_clean()

        return attrs
