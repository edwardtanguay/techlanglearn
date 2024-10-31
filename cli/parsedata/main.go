package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "helloworld",
		Short: "Prints Hello, World!",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
