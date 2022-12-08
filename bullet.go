/*
 * @Author: tj
 * @Date: 2022-12-08 14:57:51
 * @LastEditors: tj
 * @LastEditTime: 2022-12-08 17:27:10
 * @FilePath: \demo\bullet.go
 */
package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	Entity
	image       *ebiten.Image
	width       int
	height      int
	x           float64
	y           float64
	speedFactor float64
}

func NewBullet(cfg *Config, ship *Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)

	return &Bullet{
		image:       img,
		width:       cfg.BulletWidth,
		height:      cfg.BulletHeight,
		x:           ship.x + float64(ship.width-cfg.BulletWidth)/2,
		y:           float64(cfg.ScreenHeight - ship.height - cfg.BulletHeight),
		speedFactor: cfg.BulletSpeedFactor,
	}
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.x, bullet.y)
	screen.DrawImage(bullet.image, op)
}

func (bullet *Bullet) Update() {
	bullet.y -= bullet.speedFactor
}

func (bullet *Bullet) OutOfScreen() bool {
	return bullet.y < -float64(bullet.height)
}

func (bullet *Bullet) Width() int {
	return bullet.width
}

func (bullet *Bullet) Height() int {
	return bullet.height
}

func (bullet *Bullet) X() float64 {
	return bullet.x
}

func (bullet *Bullet) Y() float64 {
	return bullet.y
}
