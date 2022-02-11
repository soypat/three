package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// geometry_method_generator -geometryType SphereGeometry -geometrySlug sphere_geometry

import "syscall/js"

// Compile-time check that this type implements Geometry interface.
var _ Geometry = SphereGeometry{}

func (g SphereGeometry) ApplyMatrix4(matrix *Matrix4) {
	g.Value.Call("applyMatrix4", matrix)
}

func (g SphereGeometry) RotateX() {
	g.Value.Call("rotateX")
}

func (g SphereGeometry) RotateY() {
	g.Value.Call("rotateY")
}

func (g SphereGeometry) RotateZ() {
	g.Value.Call("rotateZ")
}

func (g SphereGeometry) Translate() {
	g.Value.Call("translate")
}

func (g SphereGeometry) Scale() {
	g.Value.Call("scale")
}

func (g SphereGeometry) LookAt() {
	g.Value.Call("lookAt")
}

func (g SphereGeometry) FromBufferGeometry(geometry Geometry) {
	g.Value.Call("fromBufferGeometry")
}

func (g SphereGeometry) Center() {
	g.Value.Call("center")
}

func (g SphereGeometry) Normalize() SphereGeometry {
	g.Value.Call("normalize")
	return g
}

func (g SphereGeometry) ComputeFaceNormals() {
	g.Value.Call("computeFaceNormals")
}

func (g SphereGeometry) ComputeVertexNormals(areaWeighted bool) {
	g.Value.Call("computeVertexNormals", areaWeighted)
}

func (g SphereGeometry) ComputeFlatVertexNormals() {
	g.Value.Call("computeFlatVertexNormals")
}

func (g SphereGeometry) ComputeMorphNormals() {
	g.Value.Call("computeMorphNormals")
}

func (g SphereGeometry) ComputeLineDistances() {
	g.Value.Call("computeLineDistances")
}

func (g SphereGeometry) ComputeBoundingBox() {
	g.Value.Call("computeBoundingBox")
}

func (g SphereGeometry) ComputeBoundingSphere() {
	g.Value.Call("computeBoundingSphere")
}

func (g SphereGeometry) Merge(geometry Geometry, matrix Matrix4, materialIndexOffset float64) {
	g.Value.Call("merge", geometry, matrix, materialIndexOffset)
}

func (g SphereGeometry) MergeMesh(mesh Mesh) {
	g.Value.Call("mergeMesh", mesh.getInternalObject())
}

func (g SphereGeometry) MergeVertices() {
	g.Value.Call("mergeVertices")
}

func (g SphereGeometry) SortFacesByMaterialIndex() {
	g.Value.Call("sortFacesByMaterialIndex")
}

func (g SphereGeometry) ToJSON() js.Value {
	return g.Value.Call("toJSON")
}

// func (g SphereGeometry) Clone() SphereGeometry {
// 	return g.Value.Call("clone")
// }

func (g SphereGeometry) Copy(source Object3D, recursive bool) *SphereGeometry {
	return &SphereGeometry{Value: g.getInternalObject().Call("copy", source.getInternalObject(), recursive)}
}

func (g SphereGeometry) Dispose() {
	g.Value.Call("dispose")
}

func (g SphereGeometry) getInternalObject() js.Value {
	return g.Value
}