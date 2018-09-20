package platform

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/thegtproject/gravity/internal/gravitygl"
)

// Platform ...
type Platform interface {
	Run(f func())
	Running() bool
	Update()
	SetClearColor(mgl32.Vec4)
	Stop()
	SetTitle(val string)
}

// New ...
func New(title string, width int, height int, vsync bool) Platform {
	println("Platform.New()")
	platform := newPlatform(title, width, height, vsync)
	gravitygl.Init()
	println("OpenGL version " + gravitygl.GetGLVersion())
	return platform
}
