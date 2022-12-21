//go:build mage

package main

import (
	"errors"
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"strconv"
)

const dockerTag = "discgolfapi"

// Builds a dockerfile
func Docker() error {
	if err := ensureDocker(); err != nil {
		return err
	}

	if err := sh.RunV("docker", "build", "-t", dockerTag, "."); err != nil {
		return fmt.Errorf("running build: %w", err)
	}

	return nil
}

// Runs the server locally
func Local() error {
	mg.Deps(Docker)

	if err := ensureDocker(); err != nil {
		return err
	}

	port := 8080
	if err := sh.RunV("docker", "run",
		"-p", fmt.Sprintf("%[1]d:%[1]d", port),
		"-t", dockerTag,
		"./disc-golf-api",
		"server",
		"-p", strconv.Itoa(port),
	); err != nil {
		return err
	}

	return nil
}

// Generates the api code from the open api spec
func Openapi() error {
	if err := ensureDocker(); err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working directory: %w", err)
	}

	if err := sh.RunV("docker", "run",
		"--rm",
		"-v", fmt.Sprintf("/%s:/disc-golf-api", wd),
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
