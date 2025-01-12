package main

import (
	"fmt"
	"os"

	"github.com/jarededwards/goop/internal/kubefirst/config"
	"github.com/jarededwards/goop/internal/kubefirst/generate"
)

func main() {

	config, err := config.ReadPlatformConfig()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)

	}

	err = generate.GenerateExternalDNSApp(config)
	if err != nil {
		fmt.Printf("Error generating external dns app: %v\n", err)
		os.Exit(1)
	}

	//! decorate annotations for ingress
	//! decorate annotations for sync wave ordering
	//! should add a kubefirst.konstruct.io/$somehting annotation for tracking?

}
