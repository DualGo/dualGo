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
}
//- ## Struct Sprite `implements Drawable2D`

type Sprite struct {
	position       mgl32.Vec2
	size           mgl32.Vec2
	origin         mgl32.Vec2
	scale          float32
	angle          float32
	scaleMat       mgl32.Mat4
	rotationMat    mgl32.Mat4
	translationMat mgl32.Mat4
	shader         *shader.Shader
	texture        uint32
	err            error
}

//	- ### Init(position, size `mgl32.Vec2`, texturePath `string`, shader `*shader.Shader`)
//		- > init the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) Init(position, size mgl32.Vec2, texturePath string, shader *shader.Shader) {
	sprite.shader = shader
	sprite.shader.Use()
	sprite.position = position
	sprite.size = size
	sprite.origin = mgl32.Vec2{0, 0}
	sprite.scale = 1
	sprite.angle = 0
	//load texture
	sprite.texture, sprite.err = texture.NewTexture(constants.Param.TexturePath, texturePath)
	if sprite.err != nil {
		log.Fatalln(sprite.err)
	}
	gl.UseProgram(0)

}

//	- ### Push()
//		- > push sprite element to be rendered 
// 
//		- > return `void`
// 
func (sprite Sprite) Push() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, sprite.texture)

	model := mgl32.Translate3D(sprite.position.X(), sprite.position.Y(), 0)
	model = model.Mul4(mgl32.Translate3D(sprite.origin.X(), sprite.origin.Y(), 0))
	model = model.Mul4(mgl32.HomogRotate3D(sprite.angle, mgl32.Vec3{0, 0, 1}))
	model = model.Mul4(mgl32.Translate3D(-sprite.origin.X(), -sprite.origin.Y(), 0))
	model = model.Mul4(mgl32.Scale3D(sprite.size.X(), sprite.size.Y(), 1))

	modelUniform := gl.GetUniformLocation(sprite.shader.GetProgram(), gl.Str("model\x00"))

	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

}
//	- ### Pop()
//		- > Pop elements 
// 
//		- > return `void`
// 
func (sprite Sprite) Pop() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

//	- ### Move(x, y `float32`)
//		- > move the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) Move(x, y float32) {
	sprite.SetPosition(mgl32.Vec2{sprite.position.X() + x, sprite.position.Y() + y})
}

//	- ### SetScale(scale `float32`)
//		- > set the scale of the scale 
// 
//		- > return `void`
// 
func (sprite *Sprite) SetScale(scale float32) {
	sprite.scale = scale
}

//	- ### GetScale()
//		- > return the scale of the sprite
// 
//		- > return `float32`
// 
func (sprite Sprite) GetScale() float32 {
	return sprite.scale
}

//	- ### SetAngle(angle `float32`)
//		- > set rotation angle of the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) SetAngle(angle float32) {
	sprite.angle = angle
}

//	- ### GetAngle()
//		- > return rotation angle of the sprite
// 
//		- > return `float32`
// 
func (sprite Sprite) GetAngle() float32 {
	return sprite.angle
}

//	- ### SetPosition(position `mgl32.Vec2`)
//		- > set the position of the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) SetPosition(position mgl32.Vec2) {
	sprite.position = position
}

//	- ### GetPosition()
//		- > return the position of the sprite
// 
//		- > return `mgl32.Vec2`
// 
func (sprite Sprite) GetPosition() mgl32.Vec2 {
	return sprite.position
}

//	- ### SetSize(size  `mgl32.Vec2`)
//		- > set the size of the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) SetSize(size mgl32.Vec2) {
	sprite.size = size
}

//	- ### GetSize()
//		- > return the size of the sprite
// 
//		- > return `mgl32.Vec2`
// 
func (sprite Sprite) GetSize() mgl32.Vec2 {
	return sprite.size
}

//	- ### SetOrigin(origin `mgl32.Vec2`)
//		- > set oririgin of the sprite
// 
//		- > return `void`
// 
func (sprite *Sprite) SetOrigin(origin mgl32.Vec2) {
	sprite.origin = origin
}

//	- ### GetOrigin()
//		- > return the origin of the srite
// 
//		- > return `mgl32.Vec2`
// 
func (sprite Sprite) GetOrigin() mgl32.Vec2 {
	return sprite.origin
}

//	- ### GetShader()
//		- > return the sahder of the sprite
// 
//		- > return `*shader.Shader`
// 
func (sprite Sprite) GetShader() *shader.Shader {
	return sprite.shader
}
