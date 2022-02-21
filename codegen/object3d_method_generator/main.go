// The following directive is necessary to make the package coherent:
//go:build ignore
// +build ignore

package main

import (
	_ "embed"
	"flag"
	"log"

	"github.com/soypat/three/codegen/generator"
)

//go:embed _template.go
var object3DTemplate string

var (
	typeName = flag.String("typeName", "", "Name of class that extends Object3D e.g. MeshBasicMaterial")
	typeSlug = flag.String("typeSlug", "", "Slugified name of class e.g. mesh_basic_material")
)

func main() {
	flag.Parse()

	if *typeName == "" {
		log.Fatal("a type name argument must be provided (e.g. -typeName MeshBasicMaterial)")
	}
	if *typeSlug == "" {
		log.Fatal("a type slug argument must be provided (e.g. -typeSlug mesh_basic_material)")
	}
	p := generator.Parameters{
		FilePrefix: "gen_object3d",
		Template:   object3DTemplate,
		Slug:       *typeSlug,
		Type:       *typeName,
	}
	err := generator.Execute(p)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Generated file: %s", p.Filename())
}
