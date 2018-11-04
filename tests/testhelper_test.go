package tests

import (
	"bytes"
	"github.com/alecthomas/kong"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/octo-cli/octo-cli/generated"
	"github.com/octo-cli/octo-cli/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	tkn, ok := os.LookupEnv("TESTUSER_TOKEN")
	if !ok {
		tkn = "deadbeef"
		recorderMode = recorder.ModeReplaying
	} else {
		recorderMode = recorder.ModeRecording
	}
	_ = os.Setenv("GITHUB_TOKEN", tkn)
}

var recorderMode recorder.Mode

func startVCR(t *testing.T, recPath string) func() {
	t.Helper()
	var err error
	rec, err := recorder.NewAsMode(recPath, recorderMode, nil)
	rec.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		var b bytes.Buffer
		if r.Body != nil {
			if _, err := b.ReadFrom(r.Body); err != nil {
				return false
			}
		}
		r.Body = ioutil.NopCloser(&b)
		return cassette.DefaultMatcher(r, i) &&
			(b.String() == "" || b.String() == i.Body) &&
			r.Header.Get("Accept") == i.Headers.Get("Accept")
	})
	require.Nil(t, err)
	internal.TransportWrapper = rec
	return func() {
		t.Helper()
		require.Nil(t, rec.Stop())
	}
}

type cmdLine []string

func newCmdLine(cmd string, args ...string) cmdLine {
	return append(cmdLine{cmd}, args...)
}

func (c cmdLine) test(t *testing.T, fixtureName, wantStdout, wantStderr string, wantErr bool) {
	t.Helper()
	stdout, stderr, err := runCmdLine(t, fixtureName, &generated.CLI{}, c...)
	if wantErr {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}
	assert.Equal(t, wantStdout, stdout.String())
	assert.Equal(t, wantStderr, stderr.String())
}

func runCmdLine(t *testing.T, fixtureName string, cmdStruct interface{}, cmdline ...string) (sout bytes.Buffer, serr bytes.Buffer, err error) {
	t.Helper()
	recCleanup := startVCR(t, filepath.Join("testdata", "fixtures", fixtureName))
	defer recCleanup()
	p, e := kong.New(cmdStruct)

	require.NoError(t, e)
	internal.Stdout = &sout

	k, e := p.Parse(cmdline)
	require.NoError(t, e)
	valueIsSetMap := map[string]bool{}
	for _, flag := range k.Flags() {
		valueIsSetMap[flag.Name] = flag.Set
	}
	err = k.Run(valueIsSetMap)
	return
}
