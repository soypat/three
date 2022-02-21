// The following directive is necessary to make the package coherent:
//go:build ignore
// +build ignore

package main

import (
	_ "embed"
	"flag"
	"log"
	"strings"

	"github.com/soypat/three/codegen/generator"
)

//go:embed _template.go
var cameraTemplate string

var (
	typeName = flag.String("typeName", "", "Name of class that extends Camera e.g. PerspectiveCamera")
	typeSlug = flag.String("typeSlug", "", "Slugified name of class e.g. perspective_camera")
)

func main() {
	flag.Parse()

	if *typeName == "" {
		log.Fatal("a type name argument must be provided (e.g. -typeName PerspectiveCamera)")
	}
	if *typeSlug == "" {
		log.Fatal("a type slug argument must be provided (e.g. -typeSlug perspective_camera)")
	}
	p := generator.Parameters{
		FilePrefix: "gen_camera",
		Template:   cameraTemplate,
		Slug:       strings.TrimSuffix(*typeSlug, "_camera"),
		Type:       *typeName,
	}
	err := generator.Execute(p)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Generated file: %s", p.Filename())
}
