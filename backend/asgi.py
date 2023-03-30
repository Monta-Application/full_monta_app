# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2023 MONTA                                                                         -
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

import os

from art import *
from channels.routing import ProtocolTypeRouter, URLRouter
from django.core.asgi import get_asgi_application
from rich.console import Console
from organization.routing import websocket_urlpatterns
import warnings

warnings.warn("Running on PyPy is not fully supported. Be aware some features may not work as expected.", RuntimeWarning)

console = Console()
logo = text2art("MONTA", font="Larry 3D")  # type: ignore
console.print(logo, style="bold purple")

os.environ.setdefault("DJANGO_SETTINGS_MODULE", "backend.settings")

django_asgi_app = get_asgi_application()


application = ProtocolTypeRouter(
    {"http": django_asgi_app, "websocket": URLRouter(websocket_urlpatterns)}
)
