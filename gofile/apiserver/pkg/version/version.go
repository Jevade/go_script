package version

import (
	"fmt"
	"runtime"
)

// Info contains versioning information.
type Info struct {
	GitTag       string `json:"GitTag"`
	GitCommit    string `json:"GitCommit"`
	GitTreeState string `json:"GitTreeState"`
	BuildDate    string `json:"BuildDate"`
	GoVersion    string `json:"GoVersion"`
	Compiler     string `json:"Compiler"`
	Platform     string `json:"Platform"`
}

//String return gittag
func (info *Info) String() string {
	return info.GitTag
}

func PrintStr() {
	fmt.Println(GitTag)
}

//Get is
func Get() Info {
	return Info{
		GitTag:       GitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
