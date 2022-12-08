/*
 * @Author: tj
 * @Date: 2022-12-08 14:07:22
 * @LastEditors: tj
 * @LastEditTime: 2022-12-08 17:51:49
 * @FilePath: \demo\input.go
 */
package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	msg            string
	lastBulletTime time.Time
}

func (i *Input) Update(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		i.msg = "left pressed"
		g.ship.x -= g.cfg.ShipSpeedFactor
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		i.msg = "right pressed"
		g.ship.x += g.cfg.ShipSpeedFactor
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && len(g.bullets) < g.cfg.MaxBulletNum && time.Since(i.lastBulletTime).Milliseconds() > int64(g.cfg.BulletInterval) {
		i.msg = "space pressed"
		bullet := NewBullet(g.cfg, g.ship)
		g.addBullet(bullet)
		i.lastBulletTime = time.Now()
	}
}
