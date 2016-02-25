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

  px, err := pixelizr.NewPixelizr("src/harvey.jpg", 60)
  if (err != nil) {
    panic(err.Error())
  }
  px.MultiChannelCircles("dest/harvey.png", 0)
}