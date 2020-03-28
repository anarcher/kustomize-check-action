package config

import (
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

type Inputs struct {
	Paths        []string `env:"INPUT_PATHS" envSeparator:"\n"`
	CommitBefore string   `env:"INPUT_GITHUB_SHA_BEFORE"`
}

type Config struct {
	GitHub GitHub
	Inputs Inputs
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

	return &Config{
		GitHub: github,
		Inputs: inputs,
	}, nil
}
