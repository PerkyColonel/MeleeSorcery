package Window

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/PerkyColonel/MeleeSorcery/tree/main/drawing"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// const (
// 	vertexShaderSource = `
// 	#version 410
// 	in vec3 vp;
// 	void main() {
// 		gl_Position = vec4(vp, 1.0);
// 	}
// ` + "\x00"

// 	fragmentShaderSource = `
// 	#version 410
// 	out vec4 frag_colour;
// 	void main() {
// 		frag_colour = vec4(1, 1, 1, 1);
// 	}
// ` + "\x00"
// )

var (
	vertexShaderSource   string
	fragmentShaderSource string
)

func Startup() {
	openShaders()
	setupWindowAndContext()
}

func openShaders() {

	fragmentShaderSource = openfile("./shaders/fragmentshaders/basicfragshader.frag")

	vertexShaderSource = openfile("./shaders/vertexshaders/basicvertshader.vert")

	fmt.Println(vertexShaderSource)
}

func openfile(path string) string {
	shader, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return string(shader)
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

var (
	triangle = []float32{
		0, 0.2, 0, // top
		-0.2, -0.2, 0, // left
		0.2, -0.2, 0, // right
	}
)

func setupWindowAndContext() {
	runtime.LockOSThread()

	window := initGlfw()
	defer glfw.Terminate()

	program := initOpenGL()

	// allTestShapes := [][]float32{triangle}

	drawing.SetWindow(window)
	drawing.SetProgram(program)

	drawing.AddQuad(0, 0, 0, 0)
	// drawing.AddQuad(2, 2, 1, 1)

	// vao := drawing.MakeVao(triangle)

	for !window.ShouldClose() {

		drawing.DrawAllQuads()

	}
}

func initGlfw() *glfw.Window {

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	width := 800
	height := 600

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Cliffhanger_enjin", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
