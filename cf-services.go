package main

import (
	"github.com/cloudfoundry-community/go-cfenv"
)

func getServiceURI(label string) string {
	// get cloudfoundry env
	appEnv, _ := cfenv.Current()

	if appEnv != nil {
		services, _ := appEnv.Services.WithLabel("elephantsql")
		if len(services) >= 1 {
			return services[0].Credentials["uri"].(string)
		}
	}
	return ""
}
