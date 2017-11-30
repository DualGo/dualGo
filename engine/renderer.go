package engine

import (
	"github.com/DualGo/glHelper"
)

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


type Renderer struct{
	system *System
	width, height  int
	shader         Shader
	//camera         camera.Camera2D
	vao            uint32
	vertAttrib     uint32
	texCoordAttrib uint32
	//fonts          []texture.Font
	//fontIndex      int
}

func (renderer *Renderer) Init(vertexShader, fragmentShader string){
	renderer.system = &System{}
	renderer.system.Init = func() {
		renderer.width, renderer.height = 800, 600
		//renderer.fontIndex = 0
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
		//renderer.camera.Init(&renderer.shader, renderer.width, renderer.height)
		glHelper.Context.UseProgram(0)
		println("Renderer init")
	}

	renderer.system.Update = func() {
		println("renderer update")
		//dessine la liste <s
	}
}

func (renderer Renderer) initVao(){

}

func (renderer Renderer) GetSystem() *System{
	return renderer.system
}

func (renderer *Renderer) Draw(){
	//append dans une liste d'element a dessiner
}

//function with drawable parameter which is going to be draw
func(renderer *Renderer) draw(){
	//drawable.GetShader().Use()
	//renderer.camera.Update(drawable.GetShader())
	//gl.Enable(gl.BLEND)
	//gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	//drawable.Push()
	//defer drawable.Pop()
	////position
	//gl.BindVertexArray(renderer.vao)
	//gl.EnableVertexAttribArray(renderer.vertAttrib)
	//gl.EnableVertexAttribArray(renderer.texCoordAttrib)
	//defer gl.DisableVertexAttribArray(renderer.vertAttrib)
	//defer gl.DisableVertexAttribArray(renderer.texCoordAttrib)
	//
	//gl.DrawArrays(gl.TRIANGLES, 0, 6)
	//gl.BindVertexArray(0)
	//gl.UseProgram(0)
	//gl.Disable(gl.BLEND)
}
