/*
 * @Author: tj
 * @Date: 2022-12-08 17:17:53
 * @LastEditors: tj
 * @LastEditTime: 2022-12-08 17:27:40
 * @FilePath: \demo\object.go
 */
package main

type Entity interface {
	Width() int
	Height() int
	X() float64
	Y() float64
}
