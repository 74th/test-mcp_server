from mcp.server.fastmcp import FastMCP
from mcp.server.fastmcp.resources import Resource

mcp = FastMCP("my_name")

class Doc(Resource):
    async def read(self) -> str:
        return "Atsushi Morimoto"

mcp.add_resource(Doc(
    uri="doc://74th/name",
    name="74thの本名",
    description="74thの本名",
    mime_type="plain/text",
))