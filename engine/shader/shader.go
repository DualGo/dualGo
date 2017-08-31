package shader

import (
	"fmt"
	"strings"

	"github.com/DualGo/gl/v4.1-core/gl"
)

const (
	vertexShaderSource = `
	#version 330
	uniform mat4 projection;
	uniform mat4 camera;
	uniform mat4 model;
	in vec3 vert;
	in vec2 vertTexCoord;
	out vec2 fragTexCoord;
	void main() {
		fragTexCoord = vertTexCoord;
		gl_Position = projection * camera * model * vec4(vert, 1);
	}
	` + "\x00"

	fragmentShaderSource = `
	#version 330
	uniform sampler2D tex;
	in vec2 fragTexCoord;
	out vec4 outputColor;
	void main() {
		outputColor = texture(tex, fragTexCoord);
	}
	` + "\x00"
)

type Shader struct {
	vertexShader   uint32
	fragmentShader uint32
	prog           uint32
}

func (this *Shader) Init(vertexSource, fragmentSource string) error {
	var err error
	if vertexSource != "" {
		this.vertexShader, err = compileShader(vertexSource, gl.VERTEX_SHADER)
	} else {
		this.vertexShader, err = compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	}
	if err != nil {
		panic(err)
	}
	if fragmentSource != "" {
		this.fragmentShader, err = compileShader(fragmentSource, gl.FRAGMENT_SHADER)
	} else {
		this.fragmentShader, err = compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	}
	if err != nil {
		panic(err)
	}
	this.prog = gl.CreateProgram()
	gl.AttachShader(this.prog, this.vertexShader)
	gl.AttachShader(this.prog, this.fragmentShader)
	gl.LinkProgram(this.prog)
	gl.DeleteShader(this.vertexShader)
	gl.DeleteShader(this.fragmentShader)
	return nil
}

func (this *Shader) Use() {
	gl.UseProgram(this.prog)
}

func (this *Shader) GetProgram() uint32 {
	return this.prog
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
