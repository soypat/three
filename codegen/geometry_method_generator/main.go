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
var geometryTemplate string

var (
	geometryType = flag.String("typeName", "", "Name of class that extends Geometry e.g. CircleGeometry")
	geometrySlug = flag.String("typeSlug", "", "Slugified name of class e.g. circle_geometry")
)

func main() {
	flag.Parse()

	if *geometryType == "" {
		log.Fatal("a geometry name argument must be provided (e.g. -geometryType CircleGeometry)")
	}
	if *geometrySlug == "" {
		log.Fatal("a geometry slug argument must be provided (e.g. -geometrySlug circle_geometry)")
	}
	p := generator.Parameters{
		FilePrefix: "gen_geometry",
		Template:   geometryTemplate,
		Slug:       strings.TrimSuffix(*geometrySlug, "_geometry"),
		Type:       *geometryType,
	}
	err := generator.Execute(p)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Generated file: %s", p.Filename())
}
