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

# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2023 MONTA                                                                         -
#                                                                                                  -
#  This file is part of Monta.                                                                     -
#                                                                                                  -
#  Monta is free software: you can redistribute it and/or modify                                   -
#  it under the terms of the GNU General Public License as published by                            -
#  the Free Software Foundation, either version 3 of the License, or                               -
#  (at your option) any later version.                                                             -
#                                                                                                  -
#  Monta is distributed in the hope that it will be useful,                                        -
#  but WITHOUT ANY WARRANTY; without even the implied warranty of                                  -
#  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the                                   -
#  GNU General Public License for more details.                                                    -
#                                                                                                  -
#  You should have received a copy of the GNU General Public License                               -
#  along with Monta.  If not, see <https://www.gnu.org/licenses/>.                                 -
# --------------------------------------------------------------------------------------------------

from typing import Any

from django.db.models import QuerySet
from rest_framework import status, permissions
from rest_framework.authtoken.views import ObtainAuthToken
from rest_framework.generics import UpdateAPIView
from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.views import APIView

from accounts import models, serializers
from utils.exceptions import InvalidTokenException
from utils.views import OrganizationMixin


class TokenProvisionView(ObtainAuthToken):
    throttle_scope = "auth"
    permission_classes = (permissions.AllowAny,)
    serializer_class = serializers.TokenProvisionSerializer

    def post(self, request: Request, *args: Any, **kwargs: Any) -> Response:
        serializer = self.serializer_class(data=request.data)
        serializer.is_valid(raise_exception=True)
        user = serializer.validated_data["user"]
        token, _ = models.Token.objects.get_or_create(user=user)
        if token.is_expired:
            token.delete()
            token = models.Token.objects.create(user=user)

        return Response(
            {
                "id": user.id,
                "first_name": user.profile.first_name,
                "last_name": user.profile.last_name,
                "email": user.email,
                "organization_id": user.organization.id,
                "department_id": user.department.id if user.department else None,
                "token": token.key,
            }
        )


class UserViewSet(OrganizationMixin):
    """
    User ViewSet to manage requests to the user endpoint
    """

    serializer_class = serializers.UserSerializer
    queryset = models.User.objects.all()
    filterset_fields = ["department__name", "is_staff"]

    def get_queryset(self) -> QuerySet[models.User]:  # type: ignore
        """Filter the queryset to only include the current user

        Returns:
            QuerySet[models.User]: Filtered queryset
        """

        return self.queryset.filter(organization=self.request.user.organization).select_related(  # type: ignore
            "organization",
            "profiles",
            "profiles__title",
            "profiles__user",
            "department",
        )


class UpdatePasswordView(UpdateAPIView):
    """
    An endpoint for changing password.
    """

    throttle_scope = "auth"
    serializer_class = serializers.ChangePasswordSerializer

    def update(self, request: Request, *args: Any, **kwargs: Any) -> Response:
        """Handle update requests

        Args:
            request (Request): Request object
            *args (Any): Arguments
            **kwargs (Any): Keyword Arguments

        Returns:
            Response: Response of the updated user
        """

        serializer = self.get_serializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        serializer.save()
        return Response(
            "Password updated successfully",
            status=status.HTTP_200_OK,
        )


class JobTitleViewSet(OrganizationMixin):
    """
    Job Title ViewSet to manage requests to the job title endpoint
    """

    serializer_class = serializers.JobTitleSerializer
    queryset = models.JobTitle.objects.all()
    filterset_fields = ["is_active", "name"]

    def get_queryset(self) -> QuerySet[models.JobTitle]:
        """Filter the queryset to only include the current user

        Returns:
            QuerySet[models.JobTitle]: Filtered queryset
        """

        return self.queryset.filter(
            organization=self.request.user.organization  # type: ignore
        ).select_related("organization")


class TokenVerifyView(APIView):
    """
    Rest API endpoint for users can verify a token
    """

    permission_classes: list[Any] = []
    serializer_class = serializers.VerifyTokenSerializer

    def post(self, request: Request, *args: Any, **kwargs: Any) -> Response:
        """Handle Post requests
        Args:
            request (Request): Request object
            *args (Any): Arguments
            **kwargs (Any): Keyword Arguments
        Returns:
            Response: Response of token and user id
        """

        serializer = self.serializer_class(data=request.data)
        serializer.is_valid(raise_exception=True)

        token = serializer.data.get("token")

        try:
            token = models.Token.objects.get(key=token)
        except models.Token.DoesNotExist:
            raise InvalidTokenException("Token is invalid")

        user = (
            models.User.objects.select_related("profiles", "organization", "department")
            .only(
                "id",
                "username",
                "email",
                "last_login",
                "is_staff",
                "is_superuser",
                "organization__id",
                "department__id",
                "profiles__first_name",
                "profiles__last_name",
            )
            .get(id=token.user.id)
        )

        return Response(
            {
                "token": token.key,
                "id": user.id,
                "organization_id": user.organization.id,
                "department_id": user.department.id if user.department else None,
                "username": user.username,
                "first_name": user.profile.first_name,
                "last_name": user.profile.last_name,
                "full_name": f"{user.profile.first_name} {user.profile.last_name}",
                "email": user.email,
            }
        )
