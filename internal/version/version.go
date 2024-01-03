package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

// Info contains version information.
// Borrowed from https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/version/types.go.
type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// JSON returns info as a full version string with json format
func (info Info) JSON() string {
	bytes, err := json.Marshal(info)
	if err != nil {
		logrus.Fatalln(err)
	}
	return string(bytes)
}

// Short returns info as a human-friendly version string.
func (info Info) Short() string {
	return fmt.Sprintf("%s (%s)", info.GitVersion, info.GitCommit)
}

func (info Info) String() string {
	return fmt.Sprintf("Version: %s (%s), Build date: %s", info.GitVersion, info.GitCommit, info.BuildDate)
}

var (
	gitVersion   = "dev"
	gitCommit    string
	gitTreeState string
	buildDate    string

	info = Info{
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
)

func GetInfo() Info {
	return info
}

func GitVersion() string {
	return gitVersion
}

func Short() string {
	return info.Short()
}
