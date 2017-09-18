package d2d

import (
	"log"

	"github.com/DualGo/dualGo/engine/shader"
	"github.com/DualGo/dualGo/engine/texture"
	"github.com/DualGo/dualGo/engine/utils"

	"github.com/DualGo/gl/v4.1-core/gl"
	"github.com/DualGo/mathgl/mgl32"
)

const (
	rectangleVertexShaderSource = `
	#version 330 core
	uniform mat4 projection; 
	uniform mat4 model;
	uniform	vec4 color;
	uniform float stroke;
	in vec3 vertexPosition;
	in vec2 vertTexCoord;
	out vec4 colorIn;
	out vec4 position;
	out float strokeIn;
	void main(){
		colorIn = color;
		gl_Position = projection*model*vec4(vertexPosition, 1);
		position = vec4(vertexPosition, 1);
		strokeIn = stroke;
	}
	` + "\x00"
	rectangleFragmentShaderSource = `
	#version 330 core
	in vec4 colorIn;
	in vec4 position;
	in float strokeIn;
	out vec4 color;
	void main(){
		if(position.y > (1-strokeIn) || position.y < (0+strokeIn)){
			color = colorIn;
				
		}
		else if(position.x > (1-strokeIn)|| position.x < (0+strokeIn)){
			color = colorIn;
		}
  		
	}
	` + "\x00"

	spriteVertexShaderSource = `
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
	spriteFragmentShaderSource = `
	#version 330 core
	uniform sampler2D tex;
	in vec2 fragTexCoord;
	out vec4 color;
	void main(){
			color = texture(tex, fragTexCoord);
	}
	` + "\x00"

	circleVertexShaderSource = `
	#version 330 core
	uniform mat4 projection; 
	uniform mat4 model;
	uniform	vec4 color;
	uniform float stroke;
	uniform vec2 center;
	uniform	float radius;
	in vec3 vertexPosition;
	in vec2 vertTexCoord;
	out vec4 colorIn;
	out vec4 position;
	out float strokeIn;
	out float radiusIn;
	out vec2  centerIn;
	void main(){
		colorIn = color;
		gl_Position = projection*model*vec4(vertexPosition, 1);
		position = vec4(vertexPosition, 1);
		strokeIn = stroke;
		radiusIn = radius;
		centerIn = vec2(vertexPosition.x-radius/2, vertexPosition.y-radius/2);
	}
	` + "\x00"
	circleFragmentShaderSource = `
	#version 330 core
	in vec4 colorIn;
	in vec4 position;
	in vec2 centerIn;
	in float radiusIn;
	in float strokeIn;
	out vec4 color;
	
	float circle(in vec2 _st, in float _radius){
		vec2 dist = _st-vec2(0.5);
		return 1.-smoothstep(_radius-(_radius*0.01),
							 _radius+(_radius*0.01),
							 dot(dist,dist)*4.0);
	}
	
	void main(){
		vec2 st = centerIn.xy/vec2(0.2,0.2);
		color = vec4(vec3(circle(st, radiusIn)), 1.0);
	}
	` + "\x00"
)

type Drawable2D interface {
	Push()
	Pop()
	GetPosition() mgl32.Vec2
	GetSize() mgl32.Vec2
	GetShader() *shader.Shader
}

type Rectangle struct {
	position       mgl32.Vec2
	size           mgl32.Vec2
	origin         mgl32.Vec2
	scale          float32
	angle          float32
	color          mgl32.Vec4
	stroke         float32
	scaleMat       mgl32.Mat4
	rotationMat    mgl32.Mat4
	translationMat mgl32.Mat4
	shader         shader.Shader
	err            error
}

func (rectangle *Rectangle) Init(position, size mgl32.Vec2) {
	rectangle.shader.Init(rectangleVertexShaderSource, rectangleFragmentShaderSource)
	rectangle.shader.Use()
	rectangle.position = position
	rectangle.size = size
	rectangle.origin = mgl32.Vec2{0, 0}
	rectangle.scale = 1
	rectangle.angle = 0
	rectangle.color = constants.Param.DefaultColor
	rectangle.stroke = 0.01
	gl.UseProgram(0)
}

func (rectangle Rectangle) Push() {
	model := mgl32.Translate3D(rectangle.position.X(), rectangle.position.Y(), 0)
	model = model.Mul4(mgl32.Translate3D(rectangle.origin.X(), rectangle.origin.Y(), 0))
	model = model.Mul4(mgl32.HomogRotate3D(rectangle.angle, mgl32.Vec3{0, 0, 1}))
	model = model.Mul4(mgl32.Translate3D(-rectangle.origin.X(), -rectangle.origin.Y(), 0))
	model = model.Mul4(mgl32.Scale3D(rectangle.size.X(), rectangle.size.Y(), 1))

	modelUniform := gl.GetUniformLocation(rectangle.shader.GetProgram(), gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	colorUniform := gl.GetUniformLocation(rectangle.shader.GetProgram(), gl.Str("color\x00"))
	gl.Uniform4f(colorUniform, rectangle.color.X(), rectangle.color.Y(), rectangle.color.Z(), rectangle.color.W())

	stokeUniform := gl.GetUniformLocation(rectangle.shader.GetProgram(), gl.Str("stroke\x00"))
	gl.Uniform1f(stokeUniform, rectangle.stroke)
}

func (rectangle Rectangle) Pop() {

}

func (rectangle *Rectangle) Move(x, y float32) {
	rectangle.SetPosition(mgl32.Vec2{rectangle.position.X() + x, rectangle.position.Y() + y})
}

func (rectangle *Rectangle) SetScale(scale float32) {
	rectangle.scale = scale
}

func (rectangle Rectangle) GetScale() float32 {
	return rectangle.scale
}

func (rectangle *Rectangle) SetAngle(angle float32) {
	rectangle.angle = angle
}

func (rectangle Rectangle) GetAngle() float32 {
	return rectangle.angle
}

func (rectangle *Rectangle) SetPosition(position mgl32.Vec2) {
	rectangle.position = position
}

func (rectangle Rectangle) GetPosition() mgl32.Vec2 {
	return rectangle.position
}

func (rectangle *Rectangle) SetSize(size mgl32.Vec2) {
	rectangle.size = size
}

func (rectangle Rectangle) GetSize() mgl32.Vec2 {
	return rectangle.size
}

func (rectangle *Rectangle) SetOrigin(origin mgl32.Vec2) {
	rectangle.origin = origin
}

func (rectangle Rectangle) GetOrigin() mgl32.Vec2 {
	return rectangle.origin
}

func (rectangle *Rectangle) SetColor(color mgl32.Vec4) {
	rectangle.color = color
}

func (rectangle Rectangle) GetColor() mgl32.Vec4 {
	return rectangle.color
}

func (rectangle *Rectangle) SetStroke(stroke float32) {
	rectangle.stroke = stroke
}

func (rectangle Rectangle) GetStroke() float32 {
	return rectangle.stroke
}

func (rectangle *Rectangle) SetShader(shader shader.Shader) {
	rectangle.shader = shader
}

func (rectangle Rectangle) GetShader() *shader.Shader {
	return &rectangle.shader
}

type Circle struct {
	rectangle Rectangle
	radius    float32
	center    mgl32.Vec2
}

func (circle *Circle) Init(position mgl32.Vec2, radius float32) {
	circle.rectangle.Init(position, mgl32.Vec2{radius, radius})
	circle.rectangle.shader.Init(circleVertexShaderSource, circleFragmentShaderSource)
	circle.center = position.Add(circle.rectangle.size.Mul(1 / 2))
	circle.radius = radius
}

func (circle Circle) Push() {
	circle.rectangle.Push()
	radiusUniform := gl.GetUniformLocation(circle.rectangle.shader.GetProgram(), gl.Str("radius\x00"))
	gl.Uniform1f(radiusUniform, circle.radius)
	centerUniform := gl.GetUniformLocation(circle.rectangle.shader.GetProgram(), gl.Str("center\x00"))
	gl.Uniform2f(centerUniform, circle.center.X(), circle.center.Y())
}

func (circle Circle) Pop() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
	circle.rectangle.Pop()
}

func (circle *Circle) Move(x, y float32) {
	circle.SetPosition(mgl32.Vec2{circle.rectangle.position.X() + x, circle.rectangle.position.Y() + y})
}

func (circle *Circle) SetScale(scale float32) {
	circle.rectangle.scale = scale
}

func (circle Circle) GetScale() float32 {
	return circle.rectangle.scale
}

func (circle *Circle) SetAngle(angle float32) {
	circle.rectangle.angle = angle
}

func (circle Circle) GetAngle() float32 {
	return circle.rectangle.angle
}

func (circle *Circle) SetPosition(position mgl32.Vec2) {
	circle.rectangle.position = position
	circle.center = position.Add(circle.rectangle.size.Mul(1 / 2))
}

func (circle Circle) GetPosition() mgl32.Vec2 {
	return circle.rectangle.position
}

func (circle *Circle) SetSize(size mgl32.Vec2) {
	circle.rectangle.size = size
	circle.center = circle.rectangle.position.Add(size.Mul(1 / 2))
}

func (circle Circle) GetSize() mgl32.Vec2 {
	return circle.rectangle.size
}

func (circle *Circle) SetOrigin(origin mgl32.Vec2) {
	circle.rectangle.origin = origin
}

func (circle Circle) GetOrigin() mgl32.Vec2 {
	return circle.rectangle.origin
}

func (circle *Circle) SetCircle(radius float32) {
	circle.radius = radius
}

func (circle Circle) GetCircle() float32 {
	return circle.radius
}

func (circle Circle) GetCenter() mgl32.Vec2 {
	return circle.center
}

func (circle *Circle) SetShader(shader shader.Shader) {
	circle.rectangle.shader = shader
}

func (circle Circle) GetShader() *shader.Shader {
	return &circle.rectangle.shader
}

type Sprite struct {
	rectangle Rectangle
	texture   uint32
}

func (sprite *Sprite) Init(position, size mgl32.Vec2, texturePath string) {
	sprite.rectangle.Init(position, size)
	sprite.rectangle.shader.Init(spriteVertexShaderSource, spriteFragmentShaderSource)
	//load texture
	sprite.texture, sprite.rectangle.err = texture.NewTexture(constants.Param.TexturePath, texturePath)
	if sprite.rectangle.err != nil {
		log.Fatalln(sprite.rectangle.err)
	}

}

func (sprite Sprite) Push() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, sprite.texture)
	sprite.rectangle.Push()
}

func (sprite Sprite) Pop() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
	sprite.rectangle.Pop()
}

func (sprite *Sprite) Move(x, y float32) {
	sprite.SetPosition(mgl32.Vec2{sprite.rectangle.position.X() + x, sprite.rectangle.position.Y() + y})
}

func (sprite *Sprite) SetScale(scale float32) {
	sprite.rectangle.scale = scale
}

func (sprite Sprite) GetScale() float32 {
	return sprite.rectangle.scale
}

func (sprite *Sprite) SetAngle(angle float32) {
	sprite.rectangle.angle = angle
}

func (sprite Sprite) GetAngle() float32 {
	return sprite.rectangle.angle
}

func (sprite *Sprite) SetPosition(position mgl32.Vec2) {
	sprite.rectangle.position = position
}

func (sprite Sprite) GetPosition() mgl32.Vec2 {
	return sprite.rectangle.position
}

func (sprite *Sprite) SetSize(size mgl32.Vec2) {
	sprite.rectangle.size = size
}

func (sprite Sprite) GetSize() mgl32.Vec2 {
	return sprite.rectangle.size
}

func (sprite *Sprite) SetOrigin(origin mgl32.Vec2) {
	sprite.rectangle.origin = origin
}

func (sprite Sprite) GetOrigin() mgl32.Vec2 {
	return sprite.rectangle.origin
}

func (sprite *Sprite) SetShader(shader shader.Shader) {
	sprite.rectangle.shader = shader
}

func (sprite Sprite) GetShader() *shader.Shader {
	return &sprite.rectangle.shader
}

func (sprite *Sprite) SetTexture(texture uint32) {
	sprite.texture = texture
}

func (sprite Sprite) GetTexture() uint32 {
	return sprite.texture
}
