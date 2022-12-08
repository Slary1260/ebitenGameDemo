<!--
 * @Author: tj
 * @Date: 2022-12-08 17:52:02
 * @LastEditors: tj
 * @LastEditTime: 2022-12-08 17:55:05
 * @FilePath: \demo\readme.md
-->
# ebiten game demo

## TODO 
1. 没有声音！
2. 外星人没有横向的运动！
3. 分数都没有！

## 打包资源
使用file2byteslice包我们可以将图片和config.json文件打包进二进制程序中，之后编译生成一个二进制程序。然后拷贝这一个文件即可，不用再拷贝图片和其他配置文件
```
github.com/hajimehoshi/file2byteslice
```

```
file2byteslice -input INPUT_FILE -output OUTPUT_FILE -package PACKAGE_NAME -var VARIABLE_NAME
```


## go generate

实际上，我们可以使用go generate让上面的过程更智能一点。在main.go文件中添加如下几行注释：
```
//go:generate go install github.com/hajimehoshi/file2byteslice
//go:generate mkdir resources
//go:generate file2byteslice -input ../images/ship.png -output resources/ship.go -package resources -var ShipPng
//go:generate file2byteslice -input ../images/alien.png -output resources/alien.go -package resources -var AlienPng
//go:generate file2byteslice -input config.json -output resources/config.go -package resources -var ConfigJson
```