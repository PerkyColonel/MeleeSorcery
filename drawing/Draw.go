package drawing

import (
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	Window  *glfw.Window
	Program uint32
)

func SetWindow(window *glfw.Window) {
	Window = window
}

func SetProgram(program uint32) {
	Program = program
}

func Gldraw(vao uint32, shapes [][]float32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(Program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(shapes[0])/3))

	glfw.PollEvents()
	Window.SwapBuffers()
}

func MakeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	var f float32
	vertexSize := int32(unsafe.Sizeof(f) * 3)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, vertexSize, nil)

	return vao
}

func MakeEbo(points []float32) uint32 {
	var ibo uint32
	gl.GenBuffers(1, &ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	// 	var vao uint32
	// 	gl.GenVertexArrays(1, &vao)
	// 	gl.BindVertexArray(vao)
	// 	gl.EnableVertexAttribArray(0)
	// 	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	// 	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return ibo
}
