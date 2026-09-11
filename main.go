/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import "commits-for-version/cmd"

var (
	Version = "dev"
	Build   = "unknown"
)

func main() {
	cmd.SetVersion(Version, Build)
	cmd.Execute()
}
