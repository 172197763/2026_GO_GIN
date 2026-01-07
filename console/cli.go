package main

import (
	"gin_test/console/cmd"
	"log"
)

func main() {
	// 执行命令
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
