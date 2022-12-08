/*
 * @Author: tj
 * @Date: 2022-12-08 13:48:15
 * @LastEditors: tj
 * @LastEditTime: 2022-12-08 17:45:39
 * @FilePath: \demo\main.go
 */
//go:generate go install github.com/hajimehoshi/file2byteslice
//go:generate file2byteslice -input ship.png -output resources/ship.go -package resources -var ShipPng
//go:generate file2byteslice -input alien.png -output resources/alien.go -package resources -var AlienPng
//go:generate file2byteslice -input config.json -output resources/config.go -package resources -var ConfigJson
package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	err = ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
