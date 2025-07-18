package main

import (
	"fmt"
	"os"

	"example/web-service-gin/albums"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myapi",
		Short: "My API server using Cobra and Gin",
	}

	rootCmd.AddCommand(albums.ServeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
