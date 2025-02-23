package main

import (
	"os"

	"migadu-bridge/internal/migadubridge"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := migadubridge.NewMigaduBridgeCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
