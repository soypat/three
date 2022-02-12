package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// geometry_method_generator -geometryType BoxGeometry -geometrySlug box_geometry

import "syscall/js"

// Compile-time check that this type implements Geometry interface.
var _ Geometry = BoxGeometry{}

func (g BoxGeometry) ApplyMatrix4(matrix *Matrix4) {
	g.Value.Call("applyMatrix4", matrix)
}

func (g BoxGeometry) RotateX() {
	g.Value.Call("rotateX")
}

func (g BoxGeometry) RotateY() {
	g.Value.Call("rotateY")
}

func (g BoxGeometry) RotateZ() {
	g.Value.Call("rotateZ")
}

func (g BoxGeometry) Translate() {
	g.Value.Call("translate")
}

func (g BoxGeometry) Scale() {
	g.Value.Call("scale")
}

func (g BoxGeometry) LookAt() {
	g.Value.Call("lookAt")
}

func (g BoxGeometry) FromBufferGeometry(geometry Geometry) {
	g.Value.Call("fromBufferGeometry")
}

func (g BoxGeometry) Center() {
	g.Value.Call("center")
}

func (g BoxGeometry) Normalize() BoxGeometry {
	g.Value.Call("normalize")
	return g
}

func (g BoxGeometry) ComputeFaceNormals() {
	g.Value.Call("computeFaceNormals")
}

func (g BoxGeometry) ComputeVertexNormals(areaWeighted bool) {
	g.Value.Call("computeVertexNormals", areaWeighted)
}

func (g BoxGeometry) ComputeFlatVertexNormals() {
	g.Value.Call("computeFlatVertexNormals")
}

func (g BoxGeometry) ComputeMorphNormals() {
	g.Value.Call("computeMorphNormals")
}

func (g BoxGeometry) ComputeLineDistances() {
	g.Value.Call("computeLineDistances")
}

func (g BoxGeometry) ComputeBoundingBox() {
	g.Value.Call("computeBoundingBox")
}

func (g BoxGeometry) ComputeBoundingSphere() {
	g.Value.Call("computeBoundingSphere")
}

func (g BoxGeometry) Merge(geometry Geometry, matrix Matrix4, materialIndexOffset float64) {
	g.Value.Call("merge", geometry, matrix, materialIndexOffset)
}

func (g BoxGeometry) MergeMesh(mesh Mesh) {
	g.Value.Call("mergeMesh", mesh.getInternalObject())
}

func (g BoxGeometry) MergeVertices() {
	g.Value.Call("mergeVertices")
}

func (g BoxGeometry) SortFacesByMaterialIndex() {
	g.Value.Call("sortFacesByMaterialIndex")
}

func (g BoxGeometry) ToJSON() js.Value {
	return g.Value.Call("toJSON")
}

// func (g BoxGeometry) Clone() BoxGeometry {
// 	return g.Value.Call("clone")
// }

func (g BoxGeometry) Copy(source Object3D, recursive bool) *BoxGeometry {
	return &BoxGeometry{Value: g.getInternalObject().Call("copy", source.getInternalObject(), recursive)}
}

func (g BoxGeometry) Dispose() {
	g.Value.Call("dispose")
}

func (g BoxGeometry) getInternalObject() js.Value {
	return g.Value
}
