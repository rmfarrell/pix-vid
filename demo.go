package main

import (
  // "fmt"
  // "reflect"
  // "os"
  media_converter "./media_converter"
  // pixelizr "./pixelizer"
)

const (
  dest  string  = "./dest/"
  src   string  = "./src/"
)

var file string = "./src/daneka.mp4"

func main() {

  // Open the file
  /*
  reader, err := os.Open(file)
  if err != nil {
    panic(err.Error())
  }
  defer reader.Close()
  */

  // Store each file in memory so we get access to each frame of the animated gif

  gif := media_converter.VideoToAnimatedGif(file, 120, 60)

  media_converter.SeparateAnimatedGif(gif)

  media_converter.Cleanup()

  /*
  lwf := pixelizr.NewSvgr(imgs, 20, "lemmy_guitar")

  lwf.FunkyTriangles()
  */
}