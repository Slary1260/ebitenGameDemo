/*
 * @Author: tj
 * @Date: 2022-12-08 16:55:29
 * @LastEditors: tj
 * @LastEditTime: 2022-12-08 17:50:59
 * @FilePath: \demo\collision.go
 */
package main

// CheckCollision 一个矩形的4个顶点只要有一个位于另外一个矩形中，就认为它们碰撞
func CheckCollision(entityA, entityB Entity) bool {
	alienTop, alienLeft := entityB.Y(), entityB.X()
	alienBottom, alienRight := entityB.Y()+float64(entityB.Height()), entityB.X()+float64(entityB.Width())
	// 左上角
	x, y := entityA.X(), entityA.Y()
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	// 右上角
	x, y = entityA.X()+float64(entityA.Width()), entityA.Y()
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	// 左下角
	x, y = entityA.X(), entityA.Y()+float64(entityA.Height())
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	// 右下角
	x, y = entityA.X()+float64(entityA.Width()), entityA.Y()+float64(entityA.Height())
	if y > alienTop && y < alienBottom && x > alienLeft && x < alienRight {
		return true
	}

	return false
}
