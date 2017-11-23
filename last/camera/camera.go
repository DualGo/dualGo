package camera

import (
	"github.com/DualGo/dualGo/engine/shader"

	"github.com/DualGo/gl/v4.1-core/gl"
	"github.com/DualGo/mathgl/mgl32"
)

type Camera2D struct {
	projection        mgl32.Mat4
	projectionUniform int32
	size              mgl32.Vec2
	position          mgl32.Vec2
}

func (camera *Camera2D) Init(shader *shader.Shader, width, height int) {

	//matrice de projection
	camera.projection = mgl32.Ortho(0, float32(width), float32(height), 0, -1, 1)
	camera.projectionUniform = gl.GetUniformLocation(shader.GetProgram(), gl.Str("projection\x00"))
	gl.UniformMatrix4fv(camera.projectionUniform, 1, false, &camera.projection[0])
	//position init are 0,0
	camera.size = mgl32.Vec2{float32(width), float32(height)}
	camera.position = mgl32.Vec2{0, 0}

}

func (camera *Camera2D) Update(shader *shader.Shader) {
	gl.Viewport(0, 0, int32(camera.size.X()), int32(camera.size.Y()))
	camera.projection = mgl32.Ortho(camera.position.X(), float32(camera.size.X())+camera.position.X(), float32(camera.size.Y()+camera.position.Y()), 0, -1, 1)
	camera.projectionUniform = gl.GetUniformLocation(shader.GetProgram(), gl.Str("projection\x00"))
	gl.UniformMatrix4fv(camera.projectionUniform, 1, false, &camera.projection[0])

}

func (camera *Camera2D) Move(vector mgl32.Vec2) {
	camera.position = mgl32.Vec2{camera.position.X() + vector.X(), camera.position.Y() + vector.Y()}
}

func (camera *Camera2D) SetPosition(position mgl32.Vec2) {
	camera.position = position
}

func (camera Camera2D) GetPosition() mgl32.Vec2 {
	return camera.position
}

func (camera *Camera2D) SetSize(size mgl32.Vec2) {
	camera.size = size
}

func (camera Camera2D) GetSize() mgl32.Vec2 {
	return camera.size
}
