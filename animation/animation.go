/*
 * @Author: tj
 * @Date: 2022-12-08 18:03:42
 * @LastEditors: tj
 * @LastEditTime: 2022-12-08 18:16:35
 * @FilePath: \demo\animation\animation.go
 */
package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0  // 每个动作的X坐标
	frameOY     = 32 // 每个动作的Y坐标
	frameWidth  = 32 // 每个动作的宽
	frameHeight = 32 // 每个动作的高
	frameCount  = 8  // 人物动作数量
)

var (
	// 加载全部动作的图集
	runnerImage *ebiten.Image
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	fmt.Println("g.count:", g.count)
	i := (g.count / 5) % frameCount
	fmt.Println("i:", i)
	sx, sy := frameOX+i*frameWidth, frameOY
	fmt.Println("sx:", sx, "sy:", sy)
	// 每次只绘画一个动作
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		panic(err)
	}
	runnerImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
