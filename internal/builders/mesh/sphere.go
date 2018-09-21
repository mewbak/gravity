package meshbuilder

import (
	"github.com/thegtproject/gravity/internal/mesh"
)

// Sphere ...
func Sphere(detail int) *mesh.Mesh {
	mb := New()

	return mb.Build()
}
