package d2d

// # Package d2d
import (
	"log"

	"github.com/DualGo/dualGo/engine/shader"
	"github.com/DualGo/dualGo/engine/texture"
	"github.com/DualGo/dualGo/engine/utils"

	"github.com/DualGo/gl/v4.1-core/gl"
	"github.com/DualGo/mathgl/mgl32"
)
const(
	rectangleVertexShaderSource = `
	#version 330 core
	uniform mat4 projection; 
	uniform mat4 model;
	in vec3 vertexPosition;
	in vec2 vertTexCoord;
	void main(){
		gl_Position = projection*model*vec4(vertexPosition, 1);
	
	}
	` + "\x00"
	rectangleFragmentShaderSource = `
	#version 330 core
	out vec4 color;
	void main(){
  		color = ;
	}
	` + "\x00"
)

//- ## Interface Drawable2D
//	- > Push()
// 
//	- > Pop()
// 
//	- > GetShader() `*shader.Shader`
// 
type Drawable2D interface {
	Push()
	Pop()
	GetShader() *shader.Shader
	IsTextured() bool
}

//- ## Struct Rectangle `implements Drawable2D`

type Rectangle struct{
	position       mgl32.Vec2
	size           mgl32.Vec2
	origin         mgl32.Vec2
	scale          float32
	angle          float32
	color		   mgl32.Vec4		
	scaleMat       mgl32.Mat4
	rotationMat    mgl32.Mat4
	translationMat mgl32.Mat4
	shader         *shader.Shader
	err            error

}

//	- ### Init(position, size, `mgl32.Vec2`, shader.Shader)
//		- > init the rectangle
// 
//		- > return `void`
// 
func (rectangle *Rectangle) Init(position, size mgl32.Vec2, shader *shader.Shader){
	rectangle.shader = shader
	rectangle.shader.Use()
	rectangle.position = position
	rectangle.size = size
	rectangle.origin = mgl32.Vec2{0, 0}
	rectangle.scale = 1
	rectangle.angle = 0
	rectangle.color = constants.Param.DefaultColor
	gl.UseProgram(0)
}

//	- ### Push()
//		- > push elements to be rendered
// 
//		- > retur `void`
// 
func (rectangle Rectangle) Push(){
	model := mgl32.Translate3D(rectangle.position.X(), rectangle.position.Y(), 0)
	model = model.Mul4(mgl32.Translate3D(rectangle.origin.X(), rectangle.origin.Y(), 0))
	model = model.Mul4(mgl32.HomogRotate3D(rectangle.angle, mgl32.Vec3{0, 0, 1}))
	model = model.Mul4(mgl32.Translate3D(-rectangle.origin.X(), -rectangle.origin.Y(), 0))
	model = model.Mul4(mgl32.Scale3D(rectangle.size.X(), rectangle.size.Y(), 1))

	modelUniform := gl.GetUniformLocation(rectangle.shader.GetProgram(), gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
}

//	- ### Pop()
//		- > pop elements
// 
//		- > return `void`
// 
func (rectangle Rectangle) Pop(){

}

//	- ### IsTextured() 
//		- > to know if a primitive is textured
// 
//		- > return `bool`
// 

func (rectangle Rectangle) IsTextured() bool {
	return false
}

//	- ### Move(x, y `float32`)
//		- > move the rectangle
// 
//		- > return `void`
// 
func (rectangle *Rectangle) Move(x, y float32) {
	rectangle.SetPosition(mgl32.Vec2{rectangle.position.X() + x, rectangle.position.Y() + y})
}

//	- ### SetScale(scale `float32`)
//		- > set the scale of the scale 
// 
//		- > return `void`
// 
func (rectangle *Rectangle) SetScale(scale float32) {
	rectangle.scale = scale
}

//	- ### GetScale()
//		- > return the scale of the rectangle
// 
//		- > return `float32`
// 
func (rectangle Rectangle) GetScale() float32 {
	return rectangle.scale
}

//	- ### SetAngle(angle `float32`)
//		- > set rotation angle of the rectangle
// 
//		- > return `void`
// 
func (rectangle *Rectangle) SetAngle(angle float32) {
	rectangle.angle = angle
}

//	- ### GetAngle()
//		- > return rotation angle of the rectangle
// 
//		- > return `float32`
// 
func (rectangle Rectangle) GetAngle() float32 {
	return rectangle.angle
}

//	- ### SetPosition(position `mgl32.Vec2`)
//		- > set the position of the rectangle
// 
//		- > return `void`
// 
func (rectangle *Rectangle) SetPosition(position mgl32.Vec2) {
	rectangle.position = position
}

//	- ### GetPosition()
//		- > return the position of the rectangle
// 
//		- > return `mgl32.Vec2`
// 
func (rectangle Rectangle) GetPosition() mgl32.Vec2 {
	return rectangle.position
}

//	- ### SetSize(size  `mgl32.Vec2`)
//		- > set the size of the rectangle
// 
//		- > return `void`
// 
func (rectangle *Rectangle) SetSize(size mgl32.Vec2) {
	rectangle.size = size
}

//	- ### GetSize()
//		- > return the size of the rectangle
// 
//		- > return `mgl32.Vec2`
// 
func (rectangle Rectangle) GetSize() mgl32.Vec2 {
	return rectangle.size
}

//	- ### SetOrigin(origin `mgl32.Vec2`)
//		- > set oririgin of the rectangle
// 
//		- > return `void`
// 
func (rectangle *Rectangle) SetOrigin(origin mgl32.Vec2) {
	rectangle.origin = origin
}

//	- ### GetOrigin()
//		- > return the origin of the srite
// 
//		- > return `mgl32.Vec2`
// 
func (rectangle Rectangle) GetOrigin() mgl32.Vec2 {
	return rectangle.origin
}

//	- ### GetShader()
//		- > return the sahder of the rectangle
// 
//		- > return `*shader.Shader`
// 
func (rectangle Rectangle) GetShader() *shader.Shader {
	return rectangle.shader
}

//- ## Struct Sprite `implements Drawable2D`

type Sprite struct {
	rectangle Rectangle
	texture   uint32
}

//	- ### Init(position, size `mgl32.Vec2`, texturePath `string`, shader `*shader.Shader`)
//		- > init the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) Init(position, size mgl32.Vec2, texturePath string, shader *shader.Shader) {
	sprite.rectangle.Init(position, size, shader)
	//load texture
	sprite.texture, sprite.rectangle.err = texture.NewTexture(constants.Param.TexturePath, texturePath)
	if sprite.rectangle.err != nil {
		log.Fatalln(sprite.rectangle.err)
	}

}

//	- ### Push()
//		- > push sprite element to be rendered 
// 
//		- > return `void`
// 
func (sprite Sprite) Push() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, sprite.texture)
	sprite.rectangle.Push()
}
//	- ### Pop()
//		- > Pop elements 
// 
//		- > return `void`
// 
func (sprite Sprite) Pop() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
	sprite.rectangle.Pop()
}

//	- ### IsTextured() 
//		- > to know if a primitive is textured
// 
//		- > return `bool`
// 

func (sprite Sprite) IsTextured() bool {
	return false
}

//	- ### Move(x, y `float32`)
//		- > move the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) Move(x, y float32) {
	sprite.SetPosition(mgl32.Vec2{sprite.rectangle.position.X() + x, sprite.rectangle.position.Y() + y})
}

//	- ### SetScale(scale `float32`)
//		- > set the scale of the scale 
// 
//		- > return `void`
// 
func (sprite *Sprite) SetScale(scale float32) {
	sprite.rectangle.scale = scale
}

//	- ### GetScale()
//		- > return the scale of the sprite
// 
//		- > return `float32`
// 
func (sprite Sprite) GetScale() float32 {
	return sprite.rectangle.scale
}

//	- ### SetAngle(angle `float32`)
//		- > set rotation angle of the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) SetAngle(angle float32) {
	sprite.rectangle.angle = angle
}

//	- ### GetAngle()
//		- > return rotation angle of the sprite
// 
//		- > return `float32`
// 
func (sprite Sprite) GetAngle() float32 {
	return sprite.rectangle.angle
}

//	- ### SetPosition(position `mgl32.Vec2`)
//		- > set the position of the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) SetPosition(position mgl32.Vec2) {
	sprite.rectangle.position = position
}

//	- ### GetPosition()
//		- > return the position of the sprite
// 
//		- > return `mgl32.Vec2`
// 
func (sprite Sprite) GetPosition() mgl32.Vec2 {
	return sprite.rectangle.position
}

//	- ### SetSize(size  `mgl32.Vec2`)
//		- > set the size of the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) SetSize(size mgl32.Vec2) {
	sprite.rectangle.size = size
}

//	- ### GetSize()
//		- > return the size of the sprite
// 
//		- > return `mgl32.Vec2`
// 
func (sprite Sprite) GetSize() mgl32.Vec2 {
	return sprite.rectangle.size
}

//	- ### SetOrigin(origin `mgl32.Vec2`)
//		- > set oririgin of the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) SetOrigin(origin mgl32.Vec2) {
	sprite.rectangle.origin = origin
}

//	- ### GetOrigin()
//		- > return the origin of the srite
// 
//		- > return `mgl32.Vec2`
// 
func (sprite Sprite) GetOrigin() mgl32.Vec2 {
	return sprite.rectangle.origin
}

//	- ### GetShader()
//		- > return the sahder of the sprite
// 
//		- > return `*shader.Shader`
// 
func (sprite Sprite) GetShader() *shader.Shader {
	return sprite.rectangle.shader
}


