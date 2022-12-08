/*
 * @Author: tj
 * @Date: 2022-12-08 14:11:45
 * @LastEditors: tj
 * @LastEditTime: 2022-12-08 17:49:20
 * @FilePath: \demo\ship.go
 */
package main

import (
	"bytes"
	"log"

	_ "image/png"
	// _ "golang.org/x/image/bmp"
	// _ "image/jpeg"
	// _ "image/gif"

	"demo/resources"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct {
	Entity
	image  *ebiten.Image
	width  int
	height int
	x      float64 // x坐标
	y      float64 // y坐标
}

func NewShip(screenWidth, screenHeight int) (*Ship, error) {
	// img, _, err := ebitenutil.NewImageFromFile("ship.png")
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resources.ShipPng))
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
		x:      float64(screenWidth-width) / 2,
		y:      float64(screenHeight - height),
	}

	return ship, nil
}

func (ship *Ship) Draw(screen *ebiten.Image, cfg *Config) {
	op := &ebiten.DrawImageOptions{}

	// 左边半个机身
	if ship.x < -float64(ship.width)/2 {
		ship.x = -float64(ship.width) / 2
	}

	// 右边半个机身
	if ship.x > float64(cfg.ScreenWidth)-float64(ship.width)/2 {
		ship.x = float64(cfg.ScreenWidth) - float64(ship.width)/2
	}

	op.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(ship.image, op)
}

func (ship *Ship) Width() int {
	return ship.width
}

func (ship *Ship) Height() int {
	return ship.height
}

func (ship *Ship) X() float64 {
	return ship.x
}

func (ship *Ship) Y() float64 {
	return ship.y
}
