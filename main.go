/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"oreonproject/basalt/cmd"
	"oreonproject/basalt/utils"
)

func main() {
	utils.LogSetup("logs/main.log")
	log.Print("Logger was Setup")
	cmd.Execute()
	log.Print("Executed root command")

}
