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

  px := pixelizr.ReadImage("src/harvey.jpg")
  px.BlocksPng("dest/harvey.png")
}