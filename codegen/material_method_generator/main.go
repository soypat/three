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
var materialTemplate string

var (
	materialName = flag.String("typeName", "", "Name of class that extends Material e.g. MeshBasicMaterial")
	materialSlug = flag.String("typeSlug", "", "Slugified name of class e.g. mesh_basic_material")
)

func main() {
	flag.Parse()

	if *materialName == "" {
		log.Fatal("a material name argument must be provided (e.g. -typeName MeshBasicMaterial)")
	}
	if *materialSlug == "" {
		log.Fatal("a material slug argument must be provided (e.g. -typeSlug mesh_basic_material)")
	}
	p := generator.Parameters{
		FilePrefix: "gen_material",
		Template:   materialTemplate,
		Slug:       strings.TrimSuffix(*materialSlug, "_material"),
		Type:       *materialName,
	}
	err := generator.Execute(p)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Generated file: %s", p.Filename())
}
