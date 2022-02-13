package three

type Camera interface {
	objecter
	// Returns a Vector3 representing the world space direction
	// in which the camera is looking. (Note: A camera looks down its local, negative z-axis).
	GetWorldDirection(target Vector3) Vector3
}
