from pathlib import Path
from mcp.server.fastmcp import FastMCP
from mcp.server.fastmcp.resources import Resource

mcp = FastMCP("my datasheets")

DATASHEET_DIR = Path(__file__).parent / "datasheets"

class Doc(Resource):
    async def read(self) -> str:
        assert self.uri.path
        file_name = self.uri.path[1:]
        file_path = DATASHEET_DIR / file_name
        with open(file_path, "r") as f:
            content = f.read()
        return content

mcp.add_resource(Doc(
    uri="file://datahseets/ch32v003-ds.txt",
    name="CH32V003 Datasheet",
    description="マイコンCH32V003のデータシート",
    mime_type="plain/text",
))

mcp.add_resource(Doc(
    uri="file://datasheets/ch32v003-rm.txt",
    name="CH32V003 Reference Manual",
    description="マイコンCH32V003のリファレンスマニュアル",
    mime_type="plain/text",
))

@mcp.tool(
    name="CH32V003 Datasheet",
    description="マイコンCH32V003のデータシート",
)
def read_datasheet() -> str:
    file_name = "ch32v003-ds.txt"
    file_path = DATASHEET_DIR / file_name
    with open(file_path, "r") as f:
        content = f.read()
    return content