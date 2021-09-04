// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	// mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build
var App = "tailscale-cloudrun"
var Image = "gcr.io/threecomma-200317/tailscale-cloudrun-reverseproxy"
var Service = "tailscale-hello"
var Region = "us-central1"

// Build handles building the app
func Build() error {
	fmt.Println("Building...")
	return sh.RunV("docker", "build", ".", "-t", Image)
}

// Push handles pushing the image
func Push() error {
	mg.Deps(Build)
	fmt.Println("Push...")
	return sh.RunV("docker", "push", Image)
}

// Deploy handles deploying to cloud run
func Deploy() error {
	return sh.Run("gcloud",
		"run",
		"deploy",
		// image name to use for deployment
		"--image",
		Image,
		"--platform",
		"managed",
		"--region",
		Region,
		// name of the service
		Service)
}

// Clean handles cleaning the project up
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll(App)
}
