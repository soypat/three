package generator

import (
	"fmt"
	"os"
	"text/template"
)

type Parameters struct {
	FilePrefix string
	Template   string
	Slug       string
	Type       string
}

func Execute(p Parameters) error {
	filePath := "./" + p.Filename()
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return template.Must(template.New("").Parse(p.Template)).Execute(f, p)
}

func (p Parameters) Filename() string {
	return fmt.Sprintf("%s_%s.go", p.FilePrefix, p.Slug)
}
