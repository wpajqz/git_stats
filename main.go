package main

import (
	"git.links123.net/links123.com/stats/cmd/http"
	"git.links123.net/links123.com/stats/cmd/version"
	"github.com/spf13/cobra"
)

var apiVersion, gitCommit, built string

// @title stats api for the project of links123.com
// @version 1.0
// @description stats api for the project of links123.com
// @host 127.0.0.1:8080/v1
func main() {
	rootCmd := &cobra.Command{
		Use:   "stats",
		Short: "stats api for the project of links123.com",
	}

	rootCmd.AddCommand(http.RunCommand())
	rootCmd.AddCommand(version.RunCommand(apiVersion, gitCommit, built))

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
