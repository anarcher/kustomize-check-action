package checker

import (
	"fmt"

	"github.com/anarcher/kustomize-check-action/pkg/command"
	"github.com/anarcher/kustomize-check-action/pkg/config"
	"github.com/anarcher/kustomize-check-action/pkg/finder"
)

const (
	kustBuildAndEvalCmdFmt = "set -o pipefail ; kustomize build %s | kubeval --ignore-missing-schemas "
)

type KustBuildAndEval struct {
	cfg    *config.Config
	finder finder.PathFinder
	cmd    command.Commander
}

func NewKustBuildAndEval(cfg *config.Config, f finder.PathFinder, cmd command.Commander) *KustBuildAndEval {
	k := &KustBuildAndEval{
		cfg:    cfg,
		finder: f,
		cmd:    cmd,
	}
	return k
}

func (k *KustBuildAndEval) Check() error {
	paths, err := k.finder.PathFind()
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		fmt.Println("kustomization not found")
		return nil
	}
	for _, p := range paths {
		fmt.Printf("PATH: %s\n", p)
		fmt.Println("CMD:", fmt.Sprintf(kustBuildAndEvalCmdFmt, p))
		_, err := k.cmd.Run("bash", "-c", fmt.Sprintf(kustBuildAndEvalCmdFmt, p))
		fmt.Println("err::::", err)
		if err != nil {
			return err
		}
	}
	return nil
}
