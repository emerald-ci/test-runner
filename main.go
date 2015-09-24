package main

import (
	"flag"
	"log"
	"os"

	"github.com/emerald-ci/test-runner/project"
)

func main() {
	var projectName string
	flag.StringVar(&projectName, "project", "project", "Prefix for the container names")
	flag.Parse()

	buildConfig, err := project.LoadBuildConfig()
	if err != nil {
		log.Fatal(err)
	}

	composeProject, err := project.BuildComposeProject(projectName, buildConfig.ComposeFile)
	if err != nil {
		log.Fatal(err)
	}

	exitCode, err := composeProject.Run(buildConfig.Service, buildConfig.CommandParts())
	if err != nil {
		log.Fatal(err)
	}

	composeProject, err = project.BuildComposeProject(projectName, buildConfig.ComposeFile)
	if err != nil {
		log.Fatal(err)
	}

	err = composeProject.Delete()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(exitCode)
}
