package finder

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/anarcher/kustomize-check-action/pkg/command"
	"github.com/anarcher/kustomize-check-action/pkg/config"
	"github.com/anarcher/kustomize-check-action/pkg/utils"
)

type GitChanged struct {
	base string
	cfg  *config.Config
	cmd  command.Commander
}

const (
	gitDiffTreeFmt    = "diff-tree --no-commit-id --name-only -r HEAD HEAD~%d"
	kustomizeFileName = "kustomization.yaml"
)

func NewGitChanged(base string, cfg *config.Config, cmd command.Commander) (*GitChanged, error) {

	pf := &GitChanged{
		base: base,
		cfg:  cfg,
		cmd:  cmd,
	}

	return pf, nil
}

func (f *GitChanged) PathFind() ([]string, error) {

	paths, err := f.gitDiffTree()
	if err != nil {
		return nil, err
	}

	paths = f.ensureDir(paths)
	paths = f.filterKustomize(paths)
	paths, err = f.matchPaths(paths)
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (f *GitChanged) gitDiffTree() ([]string, error) {
	commits := f.cfg.GithubEvent.TotalCommits()
	args := fmt.Sprintf(gitDiffTreeFmt, commits)

	out, err := f.cmd.Run("git", strings.Split(args, " ")...)
	if err != nil {
		return nil, err
	}
	paths := strings.Split(out, "\n")
	return paths, nil
}

func (f *GitChanged) ensureDir(paths []string) []string {
	pm := make(map[string]bool)
	for _, p := range paths {
		if len(p) == 0 {
			continue
		}
		d := filepath.Dir(p)
		if ok := pm[d]; !ok {
			pm[d] = true
		}
	}

	ps := make([]string, 0, len(pm))
	for p := range pm {
		ps = append(ps, p)
	}
	return ps
}

func (f *GitChanged) filterKustomize(paths []string) []string {
	var ps []string
	for _, p := range paths {
		fn := fmt.Sprintf("%s/%s", p, kustomizeFileName)
		if utils.FileExists(fn) {
			ps = append(ps, p)
		}
	}
	return ps
}

func (f *GitChanged) matchPaths(paths []string) ([]string, error) {
	patterns := f.cfg.Inputs.Paths
	if len(patterns) == 0 {
		return paths, nil
	}

	var ps []string

	for _, p := range paths {
		for _, pat := range patterns {
			matched := strings.HasPrefix(p, pat)
			if matched {
				ps = append(ps, p)
				break
			}
		}
	}
	return ps, nil
}
