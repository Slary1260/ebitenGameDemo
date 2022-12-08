/*
 * @Author: tj
 * @Date: 2022-12-08 14:07:59
 * @LastEditors: tj
 * @LastEditTime: 2022-12-08 17:50:08
 * @FilePath: \demo\config.go
 */
package main

import (
	"bytes"
	"encoding/json"
	"image/color"

	"demo/resources"
)

type Config struct {
	ScreenWidth       int        `json:"screenWidth"`
	ScreenHeight      int        `json:"screenHeight"`
	Title             string     `json:"title"`
	BgColor           color.RGBA `json:"bgColor"`
	ShipSpeedFactor   float64    `json:"shipSpeedFactor"`
	BulletWidth       int        `json:"bulletWidth"`
	BulletHeight      int        `json:"bulletHeight"`
	BulletSpeedFactor float64    `json:"bulletSpeedFactor"`
	BulletColor       color.RGBA `json:"bulletColor"`
	MaxBulletNum      int        `json:"maxBulletNum"`
	BulletInterval    int        `json:"bulletInterval"`
	AlienSpeedFactor  float64    `json:"alienSpeedFactor"`
	TitleFontSize     int        `json:"titleFontSize"`
	FontSize          int        `json:"fontSize"`
	SmallFontSize     int        `json:"smallFontSize"`
}

func loadConfig() *Config {
	// f, err := os.Open("config.json")
	// if err != nil {
	// 	log.Fatalf("os.Open failed: %v\n", err)
	// }
	// err = json.NewDecoder(f).Decode(&cfg)

	var cfg Config
	err := json.NewDecoder(bytes.NewReader(resources.ConfigJson)).Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
