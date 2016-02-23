package main

import (
  _ "fmt"
  pixelizr "./pixelizer"
)

const (
  dest  string  = "./dest/"
  src   string  = "./src/"
)

func main() {

  px := pixelizr.NewPixelizr("src/harvey.jpg", 90)
  px.BlocksPng("dest/harvey.png")
}