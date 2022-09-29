package version

import (
	"bytes"
	_ "embed"
	"github.com/fatih/color"
	"github.com/iyear/searchx/global"
	"github.com/spf13/cobra"
	"runtime"
	"text/template"
)

//go:embed version.tmpl
var version string

var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Check the version info",
	Example: "searchx version",
	Run: func(cmd *cobra.Command, args []string) {
		buf := &bytes.Buffer{}
		_ = template.Must(template.New("version").Parse(version)).Execute(buf, map[string]interface{}{
			"Version":   global.Version,
			"Commit":    global.Commit,
			"Date":      global.CommitDate,
			"GoVersion": runtime.Version(),
			"GOOS":      runtime.GOOS,
			"GOARCH":    runtime.GOARCH,
		})
		color.Blue(buf.String())
	},
}
