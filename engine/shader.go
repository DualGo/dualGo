package engine

import (
"fmt"
"strings"

"github.com/DualGo/glHelper"
)

const (
	vertexShaderDefault = `
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

	fragmentShaderDefault = `
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
	prog           int
}

func (shader *Shader) Init(vertexSource, fragmentSource string) error {
	var err error
	if vertexSource != "" {
		shader.vertexShader, err = compileShader(vertexSource, glHelper.VERTEX_SHADER)
	} else {
		shader.vertexShader, err = compileShader(vertexShaderDefault, glHelper.VERTEX_SHADER)
	}
	if err != nil {
		panic(err)
	}
	if fragmentSource != "" {
		shader.fragmentShader, err = compileShader(fragmentSource, glHelper.FRAGMENT_SHADER)
	} else {
		shader.fragmentShader, err = compileShader(fragmentShaderDefault, glHelper.FRAGMENT_SHADER)
	}
	if err != nil {
		panic(err)
	}
	glHelper.Context.CreateProgram(shader.prog)
	glHelper.Context.AttachShader(shader.prog, shader.vertexShader)
	glHelper.Context.AttachShader(shader.prog, shader.fragmentShader)
	glHelper.Context.LinkProgram(shader.prog)
	glHelper.Context.DeleteShader(shader.vertexShader)
	glHelper.Context.DeleteShader(shader.fragmentShader)
	return nil
}

func (shader *Shader) Use() {
	glHelper.Context.UseProgram(shader.prog)
}

func (shader *Shader) GetProgram() int {
	return shader.prog
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := glHelper.Context.CreateShader(shaderType)

	csources, free := glHelper.Context.Strs(source)
	glHelper.Context.ShaderSource(shader, 1, csources, nil)
	free()
	glHelper.Context.CompileShader(shader)

	var status int32
	glHelper.Context.GetShaderiv(shader, glHelper.Context.COMPILE_STATUS, &status)
	if status == glHelper.Context.FALSE {
		var logLength int32
		glHelper.Context.GetShaderiv(shader, glHelper.Context.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		glHelper.Context.GetShaderInfoLog(shader, logLength, nil, glHelper.Context.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
