Tags: #review #golang #programming 

The package `image` defines the `Image` interface

```go
package image

type Image interface {
    ColorModel() color.Model // Return the image's color model
    Bounds() Rectangle // Return the domain for which At() can return non-zero color
    At(x, y int) color.Color // Return the color of the pixel at (x, y)
}

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds()) // return the non-zero color domain
	fmt.Println(m.At(0, 0).RGBA()) // Return the color at (0,0) as RGBA format
}
```


## Exercise

```go
package main

import (
	"golang.org/x/tour/pic"
	"image"
  "image/color"
)

type Image struct{
	width, height int
	color uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0,i.width, i.height)
}

func (i Image) At(x,y int) color.Color {
	return color.RGBA{i.color + uint8(x), i.color + uint8(y), 255, 255}
}
	
func main() {
	m := Image{100,100,100}
	pic.ShowImage(m)
}


```
