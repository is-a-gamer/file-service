package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("Execute command error:", err.Error())
		os.Exit(1)
	}
}