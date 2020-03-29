package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/caarlos0/env/v6"
)

type GitHub struct {
	Workflow   string `env:"GITHUB_WORKFLOW"`
	Action     string `env:"GITHUB_ACTION"`
	Actor      string `env:"GITHUB_ACTOR"`
	Repository string `env:"GITHUB_REPOSITORY"`
	Commit     string `env:"GITHUB_SHA"`
	EventName  string `env:"GITHUB_EVENT_NAME"`
	EventPath  string `env:"GITHUB_EVENT_PATH"`
	Ref        string `env:"GITHUB_REF"`
}

type GithubEvent struct {
	PullRequest struct {
		Commits int `json:"commits"`
	} `json:"pull_request,omitempty"`
	Commits []interface{} `json:"commits,omitempty"`
}

func (e GithubEvent) TotalCommits() int {
	return len(e.Commits) + e.PullRequest.Commits
}

type Inputs struct {
	Paths []string `env:"INPUT_PATHS" envSeparator:"\n"`
}

type Config struct {
	GitHub      GitHub
	GithubEvent GithubEvent
	Inputs      Inputs
}

func (c Config) Print() {
	fmt.Printf("GITHUB_SHA:%s\n", c.GitHub.Commit)
	fmt.Printf("INPUT_PATHS:%s\n", c.Inputs.Paths)
	fmt.Printf("TotalCommits: %v\n", c.GithubEvent.TotalCommits())
}

func (i *Inputs) trimPaths() {
	var paths []string
	for _, p := range i.Paths {
		path := strings.TrimSpace(p)
		if path != "" {
			paths = append(paths, path)
		}
	}
	i.Paths = paths
}

func loadGitHubEvent(path string) (GithubEvent, error) {
	e := GithubEvent{}

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return e, err
	}

	if err := json.Unmarshal(bs, &e); err != nil {
		return e, err
	}

	return e, nil
}

func LoadConfigFromEnv() (*Config, error) {
	github := GitHub{}
	if err := env.Parse(&github); err != nil {
		return nil, err
	}

	inputs := Inputs{}
	if err := env.Parse(&inputs); err != nil {
		return nil, err
	}
	inputs.trimPaths()

	githubEvent, err := loadGitHubEvent(github.EventPath)
	if err != nil {
		return nil, err
	}

	return &Config{
		GitHub:      github,
		GithubEvent: githubEvent,
		Inputs:      inputs,
	}, nil
}
