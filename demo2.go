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

  px, err := pixelizr.NewPixelizr("src/harvey.jpg", 90)
  if (err != nil) {
    panic(err.Error())
  }
  px.Squares("dest/harvey.png", 0)
}