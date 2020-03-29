package finder

import (
	"reflect"
	"testing"

	"github.com/anarcher/kustomize-check-action/pkg/command"
	"github.com/anarcher/kustomize-check-action/pkg/config"
)

func TestMatchPaths(t *testing.T) {
	cmd := command.NewOSExec()

	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	f, err := NewGitChanged("", cfg, cmd)
	if err != nil {
		t.Fatal(err)
	}

	paths := []string{"test/base/t2"}
	cfg.Inputs.Paths = []string{"test/"}

	ps, err := f.matchPaths(paths)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(paths, ps) {
		t.Logf("paths=%v,ps=%v", paths, ps)
		t.Error("paths != ps")
	}
}
