package build

import (
	_ "embed"
	"strings"
)

//go:generate bash version.sh
//go:embed version.txt
var version string //nolint:gochecknoglobals

func Version() string {
	return strings.TrimSpace(version)
}
