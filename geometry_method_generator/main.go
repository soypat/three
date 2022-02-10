// The following directive is necessary to make the package coherent:
// +build ignore

package main

import (
	_ "embed"
	"flag"
	"log"

	"github.com/soypat/three/generator"
)

//go:embed _template.go
var geometryTemplate string

var (
	geometryType = flag.String("geometryType", "", "Name of class that extends Geometry e.g. CircleGeometry")
	geometrySlug = flag.String("geometrySlug", "", "Slugified name of class e.g. circle_geometry")
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
		Slug:       *geometrySlug,
		Type:       *geometryType,
	}
	err := generator.Execute(p)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Generated file: %s", p.Filename())
}
