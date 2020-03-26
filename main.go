package main

import (
	"log"

	"github.com/anarcher/kustomize-check-action/pkg/checker"
	"github.com/anarcher/kustomize-check-action/pkg/command"
	"github.com/anarcher/kustomize-check-action/pkg/config"
	"github.com/anarcher/kustomize-check-action/pkg/finder"
)

func main() {
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	base := ""
	cmd := command.NewOSExec()

	f, err := finder.NewGitChanged(base, cfg, cmd)
	if err != nil {
		log.Fatal(err)
	}

	checker := checker.NewKustBuildAndEval(cfg, f, cmd)
	if err := checker.Check(); err != nil {
		log.Fatal(err)
	}
}
