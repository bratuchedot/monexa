package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
)

func StartServer(e *echo.Echo, address string) {
	err := e.Start(address)
	if err != nil {
		fmt.Fprint(os.Stderr, "⛔ ️Exit!!! Cannot run HTTP server on port "+address+"\n")
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
