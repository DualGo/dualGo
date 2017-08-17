package renderer

import (
	"unsafe"

	"github.com/DualGo/dualGo/engine/camera"
	"github.com/DualGo/dualGo/engine/graphics/d2d"
	"github.com/DualGo/dualGo/engine/shader"

	"github.com/DualGo/gl/v4.1-core/gl"
)

type Renderer interface {
	Init(width, height int)
	GetShader() shader.Shader
}

//2D Renderer
type Renderer2D struct {
	width, height  int
	shader         shader.Shader
	camera         camera.Camera2D
	vao            uint32
	vertAttrib     uint32
	texCoordAttrib uint32
}

const (
	vertexShaderSource = `
	#version 330 core
	uniform mat4 projection; 
	uniform mat4 model;
	in vec3 vertexPosition;
	in vec2 vertTexCoord;
	out vec2 fragTexCoord;
	void main(){
		fragTexCoord = vertTexCoord;
		gl_Position = projection*model*vec4(vertexPosition, 1);
	
	}
	` + "\x00"
	fragmentShaderSource = `
	#version 330 core
	uniform sampler2D tex;
	in vec2 fragTexCoord;
	out vec4 color;
	void main(){
  		color = texture(tex, fragTexCoord);
	}
	` + "\x00"
)

func (renderer *Renderer2D) Init(width, height int, vertexShader, fragmentShader string) {
	renderer.width, renderer.height = width, height
	if vertexShader != "" {
		if fragmentShader != "" {
			renderer.shader.Init(vertexShader, fragmentShader)
		} else {
			renderer.shader.Init(vertexShader, fragmentShaderSource)
		}
	} else {
		renderer.shader.Init(vertexShaderSource, fragmentShaderSource)
	}

	renderer.shader.Use()
	renderer.initVao()
	renderer.camera.Init(&renderer.shader, renderer.width, renderer.height)

}

func (renderer *Renderer2D) initVao() {
	gl.GenVertexArrays(1, &renderer.vao)
	gl.BindVertexArray(renderer.vao)
	defer gl.BindVertexArray(0)

	vertices := []float32{

		0, 0, 0, 0.0, 0.0,
		1, 0, 0, 1.0, 0.0,
		0, 1, 0, 0.0, 1.0,
		0, 1, 0, 0.0, 1.0,
		1, 1, 0, 1.0, 1.0,
		1, 0, 0, 1.0, 0.0,
	}

	//creation du vbo
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	defer gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(unsafe.Sizeof(vertices[0])), gl.Ptr(vertices), gl.STATIC_DRAW)

	renderer.vertAttrib = uint32(gl.GetAttribLocation(renderer.shader.GetProgram(), gl.Str("vertexPosition\x00")))
	gl.EnableVertexAttribArray(renderer.vertAttrib)
	defer gl.DisableVertexAttribArray(renderer.vertAttrib)
	gl.VertexAttribPointer(renderer.vertAttrib, 3, gl.FLOAT, false, int32(5*int(unsafe.Sizeof(vertices[0]))), gl.PtrOffset(0))

	renderer.texCoordAttrib = uint32(gl.GetAttribLocation(renderer.shader.GetProgram(), gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(renderer.texCoordAttrib)
	defer gl.DisableVertexAttribArray(renderer.texCoordAttrib)
	gl.VertexAttribPointer(renderer.texCoordAttrib, 2, gl.FLOAT, false, int32(5*int(unsafe.Sizeof(vertices[0]))), gl.PtrOffset(3*int(unsafe.Sizeof(vertices[0]))))

}

func (renderer *Renderer2D) Draw(drawable d2d.Drawable2D) {
	renderer.shader.Use()
	renderer.camera.Update(&renderer.shader)
	drawable.Push(&renderer.shader)
	defer drawable.Pop()
	//position
	gl.BindVertexArray(renderer.vao)
	gl.EnableVertexAttribArray(renderer.vertAttrib)
	gl.EnableVertexAttribArray(renderer.texCoordAttrib)
	defer gl.DisableVertexAttribArray(renderer.vertAttrib)
	defer gl.DisableVertexAttribArray(renderer.texCoordAttrib)

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
	gl.BindVertexArray(0)

}

func (renderer Renderer2D) GetCamera() camera.Camera2D {
	return renderer.camera
}

func (renderer Renderer2D) GetShader() *shader.Shader {
	return &renderer.shader
}
