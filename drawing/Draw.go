package Drawing

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.0/glfw"
)

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}
