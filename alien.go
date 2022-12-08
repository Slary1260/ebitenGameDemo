/*
 * @Author: tj
 * @Date: 2022-12-08 16:44:05
 * @LastEditors: tj
 * @LastEditTime: 2022-12-08 17:50:09
 * @FilePath: \demo\alien.go
 */
package main

import (
	"bytes"

	"demo/resources"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Alien struct {
	Entity
	image       *ebiten.Image
	width       int
	height      int
	x           float64
	y           float64
	speedFactor float64
}

func NewAlien(cfg *Config) (*Alien, error) {
	// img, _, err := ebitenutil.NewImageFromFile("alien.png")
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resources.AlienPng))
	if err != nil {
		return nil, err
	}

	width, height := img.Size()
	return &Alien{
		image:       img,
		width:       width,
		height:      height,
		x:           0,
		y:           0,
		speedFactor: cfg.AlienSpeedFactor,
	}, nil
}

func (alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(alien.x, alien.y)
	screen.DrawImage(alien.image, op)
}

func (alien *Alien) Update() {
	alien.y += alien.speedFactor
}

func (alien *Alien) OutOfScreen(cfg *Config) bool {
	return alien.y > float64(cfg.ScreenHeight)
}

func (alien *Alien) Width() int {
	return alien.width
}

func (alien *Alien) Height() int {
	return alien.height
}

func (alien *Alien) X() float64 {
	return alien.x
}

func (alien *Alien) Y() float64 {
	return alien.y
}
