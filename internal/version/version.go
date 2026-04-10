package version

import "fmt"

var Version = "dev"
var Commit = "none"
var Date = "unknown"

func String() string {
	return fmt.Sprintf("mbti-cli version %s (commit %s, date %s)", Version, Commit, Date)
}
