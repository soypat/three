// The following directive is necessary to make the package coherent:
//go:build ignore
// +build ignore

package main

import (
	_ "embed"
	"flag"
	"log"

	"github.com/soypat/three/generator"
)

//go:embed _template.go
var materialTemplate string

var (
	materialName = flag.String("materialName", "", "Name of class that extends Material e.g. MeshBasicMaterial")
	materialSlug = flag.String("materialSlug", "", "Slugified name of class e.g. mesh_basic_material")
)

func main() {
	flag.Parse()

	if *materialName == "" {
		log.Fatal("a material name argument must be provided (e.g. -materialName MeshBasicMaterial)")
	}
	if *materialSlug == "" {
		log.Fatal("a material slug argument must be provided (e.g. -materialSlug mesh_basic_material)")
	}
	p := generator.Parameters{
		FilePrefix: "gen_material",
		Template:   materialTemplate,
		Slug:       *materialSlug,
		Type:       *materialName,
	}
	err := generator.Execute(p)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Generated file: %s", p.Filename())
	// filePath := fmt.Sprintf("./gen_material_%s.go", *materialSlug)

	// f, err := os.Create(filePath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// err = template.Must(template.New("").Parse(materialTemplate)).Execute(f, struct {
	// 	MaterialName string
	// 	MaterialSlug string
	// }{
	// 	MaterialName: *materialName,
	// 	MaterialSlug: *materialSlug,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Generated file: %s", filePath)
}
