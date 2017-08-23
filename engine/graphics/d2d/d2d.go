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
//- > Push()
//-
//- > Pop()
//-
//- > GetShader() *shader.Shader
//-
type Drawable2D interface {
	Push()
	Pop()
	GetShader() *shader.Shader
}

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

//- ### Init(position, size mgl32.Vec2, texturePath string, shader *shader.Shader)
//- > init the sprite
//-
//- > return void
//-
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

func (sprite Sprite) Pop() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (sprite *Sprite) Move(x, y float32) {
	sprite.SetPosition(mgl32.Vec2{sprite.position.X() + x, sprite.position.Y() + y})
}

func (sprite *Sprite) SetScale(scale float32) {
	sprite.scale = scale
}

func (sprite Sprite) GetScale() float32 {
	return sprite.scale
}

func (sprite *Sprite) SetAngle(angle float32) {
	sprite.angle = angle
}

func (sprite Sprite) GetAngle() float32 {
	return sprite.angle
}

func (sprite *Sprite) SetPosition(position mgl32.Vec2) {
	sprite.position = position
}

func (sprite Sprite) GetPosition() mgl32.Vec2 {
	return sprite.position
}

func (sprite *Sprite) SetSize(size mgl32.Vec2) {
	sprite.size = size
}

func (sprite Sprite) GetSize() mgl32.Vec2 {
	return sprite.size
}

func (sprite *Sprite) SetOrigin(origin mgl32.Vec2) {
	sprite.origin = origin
}

func (sprite Sprite) GetOrigin() mgl32.Vec2 {
	return sprite.origin
}

func (sprite Sprite) GetShader() *shader.Shader {
	return sprite.shader
}
