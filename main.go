package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/emerald-ci/test-runner/project"
)

func main() {
	var projectName string
	var configFile string
	flag.StringVar(&projectName, "project", "project", "Prefix for the container names")
	flag.StringVar(&configFile, "file", ".emerald.yml", "Config file (default `.emerald.yml`)")
	flag.Parse()

	var configContent = []byte{}

	stat, err := os.Stdin.Stat()
	if err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		configContent, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	if len(configContent) == 0 {
		fileName, _ := filepath.Abs(configFile)
		configContent, err = ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
	}

	buildConfig, err := project.BuildProjectConfig(configContent)
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
