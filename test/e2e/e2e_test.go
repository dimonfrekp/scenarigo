// +build !race

package scenarigo

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/sergi/go-diff/diffmatchpatch"
	"google.golang.org/grpc"

	"github.com/zoncoen/scenarigo"
	"github.com/zoncoen/scenarigo/context"
	"github.com/zoncoen/scenarigo/internal/testutil"
	"github.com/zoncoen/scenarigo/reporter"
	"github.com/zoncoen/scenarigo/schema"
	"github.com/zoncoen/scenarigo/testdata/gen/pb/test"
)

func TestE2E(t *testing.T) {
	dir := "testdata/testcases"
	infos, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	files := []string{}
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		if strings.HasSuffix(info.Name(), ".yaml") {
			files = append(files, filepath.Join(dir, info.Name()))
		}
	}

	teardown := startGRPCServer(t)
	defer teardown()

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			f, err := os.Open(file)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			var tc TestCase
			if err := yaml.NewDecoder(f).Decode(&tc); err != nil {
				t.Fatal(err)
			}

			for _, scenario := range tc.Scenarios {
				t.Run(scenario.Filename, func(t *testing.T) {
					config := &schema.Config{
						PluginDirectory: "testdata/gen/plugins",
						Plugins:         map[string]schema.PluginConfig{},
					}
					if scenario.Mocks != "" {
						f := filepath.Join(dir, "mocks", scenario.Mocks)
						config.Plugins["mock.so"] = schema.PluginConfig{
							Setup: fmt.Sprintf("{{SetupServerFunc(%q, %t)}}", f, !scenario.Success),
						}
					}

					r, err := scenarigo.NewRunner(
						scenarigo.WithConfig(config),
						scenarigo.WithScenarios(filepath.Join(dir, "scenarios", scenario.Filename)),
					)
					if err != nil {
						t.Fatal(err)
					}

					var b bytes.Buffer
					ok := reporter.Run(func(rptr reporter.Reporter) {
						r.Run(context.New(rptr))
					}, reporter.WithWriter(&b))
					if ok != scenario.Success {
						t.Errorf("expect %t but got %t", scenario.Success, ok)
					}

					f, err := os.Open(filepath.Join(dir, "stdout", scenario.Output.Stdout))
					if err != nil {
						t.Fatal(err)
					}
					defer f.Close()

					stdout, err := io.ReadAll(f)
					if err != nil {
						t.Fatal(err)
					}

					if got, expect := testutil.ReplaceOutput(b.String()), string(stdout); got != expect {
						dmp := diffmatchpatch.New()
						diffs := dmp.DiffMain(expect, got, false)
						t.Errorf("stdout differs:\n%s", dmp.DiffPrettyText(diffs))
					}
				})
			}
		})
	}
}

type TestCase struct {
	Tilte     string         `yaml:"title"`
	Scenarios []TestScenario `yaml:"scenarios"`
}

type TestScenario struct {
	Filename string       `yaml:"filename"`
	Mocks    string       `yaml:"mocks"`
	Success  bool         `yaml:"success"`
	Output   ExpectOutput `yaml:"output"`
}

type ExpectOutput struct {
	Stdout string `yaml:"stdout"`
}

func startGRPCServer(t *testing.T) func() {
	t.Helper()

	token := "XXXXX"
	testServer := &testGRPCServer{
		users: map[string]string{
			token: "test user",
		},
	}
	s := grpc.NewServer()
	test.RegisterTestServer(s, testServer)

	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if err := os.Setenv("TEST_GRPC_SERVER_ADDR", ln.Addr().String()); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := os.Setenv("TEST_TOKEN", token); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	go func() {
		_ = s.Serve(ln)
	}()

	return func() {
		s.Stop()
		os.Unsetenv("TEST_GRPC_SERVER_ADDR")
		os.Unsetenv("TEST_TOKEN")
	}
}
