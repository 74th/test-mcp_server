from mcp.server.fastmcp import FastMCP

def build_mcp_server() -> FastMCP:
    mcp = FastMCP("my_name")

    @mcp.resource(
            uri="users://{user_id}/name}",
            description="本名を調べる")

    async def name(user_id: str) -> str:
        return "Atsushi Morimoto"

    return mcp


def main() -> None:
    sv = build_mcp_server()

    sv.run()

if __name__ == "__main__":
    main()