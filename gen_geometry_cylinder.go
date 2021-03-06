package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// geometry_method_generator -typeName CylinderGeometry -typeSlug cylinder

import "syscall/js"

// Compile-time check that this type implements Geometry interface.
var _ Geometry = CylinderGeometry{}

func (g CylinderGeometry) ApplyMatrix4(matrix Matrix4) {
	g.Value.Call("applyMatrix4", matrix)
}

func (g CylinderGeometry) RotateX() {
	g.Value.Call("rotateX")
}

func (g CylinderGeometry) RotateY() {
	g.Value.Call("rotateY")
}

func (g CylinderGeometry) RotateZ() {
	g.Value.Call("rotateZ")
}

func (g CylinderGeometry) Translate() {
	g.Value.Call("translate")
}

func (g CylinderGeometry) Scale() {
	g.Value.Call("scale")
}

func (g CylinderGeometry) LookAt() {
	g.Value.Call("lookAt")
}

func (g CylinderGeometry) FromBufferGeometry(geometry Geometry) {
	g.Value.Call("fromBufferGeometry")
}

func (g CylinderGeometry) Center() {
	g.Value.Call("center")
}

func (g CylinderGeometry) Normalize() CylinderGeometry {
	g.Value.Call("normalize")
	return g
}

func (g CylinderGeometry) ComputeFaceNormals() {
	g.Value.Call("computeFaceNormals")
}

func (g CylinderGeometry) ComputeVertexNormals(areaWeighted bool) {
	g.Value.Call("computeVertexNormals", areaWeighted)
}

func (g CylinderGeometry) ComputeFlatVertexNormals() {
	g.Value.Call("computeFlatVertexNormals")
}

func (g CylinderGeometry) ComputeMorphNormals() {
	g.Value.Call("computeMorphNormals")
}

func (g CylinderGeometry) ComputeLineDistances() {
	g.Value.Call("computeLineDistances")
}

func (g CylinderGeometry) ComputeBoundingBox() {
	g.Value.Call("computeBoundingBox")
}

func (g CylinderGeometry) ComputeBoundingSphere() {
	g.Value.Call("computeBoundingSphere")
}

func (g CylinderGeometry) Merge(geometry Geometry, matrix Matrix4, materialIndexOffset float64) {
	g.Value.Call("merge", geometry.getInternalObject(), matrix.Value, materialIndexOffset)
}

func (g CylinderGeometry) MergeMesh(mesh Mesh) {
	g.Value.Call("mergeMesh", mesh.getInternalObject())
}

func (g CylinderGeometry) MergeVertices() {
	g.Value.Call("mergeVertices")
}

func (g CylinderGeometry) SortFacesByMaterialIndex() {
	g.Value.Call("sortFacesByMaterialIndex")
}

func (g CylinderGeometry) ToJSON() js.Value {
	return g.Value.Call("toJSON")
}

// func (g CylinderGeometry) Clone() CylinderGeometry {
// 	return g.Value.Call("clone")
// }

func (g CylinderGeometry) Copy(source Object3D, recursive bool) CylinderGeometry {
	return CylinderGeometry{Value: g.getInternalObject().Call("copy", source.getInternalObject(), recursive)}
}

func (g CylinderGeometry) Dispose() {
	g.Value.Call("dispose")
}

func (g CylinderGeometry) getInternalObject() js.Value {
	return g.Value
}
