package cli

import (
	"html/template"
	"io"
	"os"
)

const helpTemp = `{{.Name}} ― {{ .Version}}

USAGE
  {{.Name}} [command] //somepackage:somelabel ...args

COMMANDS
  {{- range .Commands}}
  {{ .Name}}  {{.Usage}}
  {{- end}}
`

func helpStd(app *App) {
	help(os.Stdout, app)
}
func help(w io.Writer, app *App) {
	tmpl, err := template.New("help").Parse(helpTemp)
	if err != nil {
		l.Fatal(err)
	}
	tmpl.Execute(w, app)
}
