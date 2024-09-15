package Drawing

import "github.com/go-gl/gl/v4.6-compatibility/gl"

var (
	vertices  = []float32{1, 1, -1, 1, -1, -1, 1, -1}
	indeces   = []float32{0, 1, 2, 2, 3, 0}
	texCoords = []float32{0.0, 0.0, 1.0, 0.0, 1.0, 1.0, 0.0, 1.0}
)

func DrawQuad(x float32, y float32, width float32, height float32) {
	
}
