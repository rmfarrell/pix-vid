package main

import (
  "fmt"
  // "reflect"
  // "os"
  media_converter "./media_converter"
  pixelizr "./pixelizer"
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

  pxs := make(chan []byte)

  gif := media_converter.VideoToAnimatedGif(file, 360, 180)

  imgs := media_converter.SeparateAnimatedGif(gif)

  fmt.Println(len(imgs))

  for i := 0; i < len(imgs); i ++ {
    go pixelizr.ReadImage(imgs[0], pxs)
  }

  media_converter.Cleanup()

  /*
  lwf := pixelizr.NewSvgr(imgs, 20, "lemmy_guitar")

  lwf.FunkyTriangles()
  */
}