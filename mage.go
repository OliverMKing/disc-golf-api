//go:build mage

package main

import (
	"errors"
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
)

func GenApi() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working directory: %w", err)
	}

	if err := ensureDocker(); err != nil {
		return err
	}

	if err := sh.Run("docker", "run", "--rm", "-v", fmt.Sprintf("/%s:/disc-golf-api", wd),
		"openapitools/openapi-generator-cli",
		"generate",
		"-i", "./disc-golf-api/openapi.yaml",
		"-g", "go-server",
		"--additional-properties=outputAsLibrary=true,sourceFolder=openapi",
		"-o", "./disc-golf-api/pkg/gen",
	); err != nil {
		return fmt.Errorf("generating api: %w", err)
	}

	return nil
}

func ensureDocker() error {
	if err := sh.Run("docker", "version"); err != nil {
		return errors.New("docker not installed")
	}

	return nil
}
