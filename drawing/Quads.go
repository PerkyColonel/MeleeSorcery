package drawing

type Quad struct {
	x      float32
	y      float32
	width  float32
	height float32
}

var (
	vertices = []float32{
		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		-0.5, -0.5, 0,
		//top left

		0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
		//down right
	}
	// indeces   = []float32{0, 1, 2, 2, 3, 0}
	// texCoords = []float32{0.0, 0.0, 1.0, 0.0, 1.0, 1.0, 0.0, 1.0}

	allQuads []Quad
)

func AddQuad(x float64, y float64, width float64, height float64) {
	quad := Quad{float32(x), float32(y), float32(width), float32(height)}

	allQuads = append(allQuads, quad)
}

func DrawAllQuads() {

	vao := MakeVao(QuadsToVertices())
	allVertices := [][]float32{QuadsToVertices()}
	Gldraw(vao, allVertices)

}

func QuadsToVertices() []float32 {
	var returnVertexArray []float32
	for i := 0; i < len(allQuads); i++ {
		q := allQuads[i]
		returnVertexArray = append(returnVertexArray,
			Modifier(-0.1, q.x, q.width), Modifier(0.1, q.y, q.height), 0,
			Modifier(0.1, q.x, q.width), Modifier(0.1, q.y, q.height), 0,
			Modifier(-0.1, q.x, q.width), Modifier(-0.1, q.y, q.height), 0,

			Modifier(0.1, q.x, q.width), Modifier(0.1, q.y, q.height), 0,
			Modifier(-0.1, q.x, q.width), Modifier(-0.1, q.y, q.height), 0,
			Modifier(0.1, q.x, q.width), Modifier(-0.1, q.y, q.height), 0,
		)
	}

	return returnVertexArray
}

func Modifier(base float32, pos float32, width float32) float32 {

	return base + (pos * (width / 2))
}
