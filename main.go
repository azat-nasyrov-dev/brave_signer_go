package main

import (
	"brave_signer/internal/logger"
	"errors"

	"brave_signer/cmd"
)

func main() {
	rootCmd := cmd.RootCmd()

	if err := rootCmd.Execute(); err != nil {
		logger.HaltOnErr(errors.New("cannot proceed, exiting now"), "Initial setup failed")
	}
}
