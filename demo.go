package main

import (
  "fmt"
  // "reflect"
  "os"
  "os/exec"
  "github.com/satori/go.uuid"
  "io/ioutil"
  media_converter "./media_converter"
  pixelizr "./pixelizer"
)

const (
  dest  string  = "./dest/"
  src   string  = "./src/"
)

var file string = "./src/lemmy_guitar.gif"

func main() {

  // Open the file
  reader, err := os.Open(file)
  if err != nil {
    panic(err.Error())
  }
  defer reader.Close()

  // Store each file in memory so we get access to each frame of the animated gif
  imgFiles := separateAnimatedGif(reader)

  // Create our 
  lwf := svgr.NewSvgr(imgFiles, 20, "lemmy_guitar")

  /*  

  Example of a 3-channel single frame ouput

  for x:=0; x < len(imgFiles); x++ {
    lwf.SingleChannel("red", "#f03c3c", .6, 50, 0, false, x)
    lwf.SingleChannel("blue", "#3c9cf0", .6, 50, -8, false, x)
    lwf.SingleChannel("green", "#63f03c", .4, 50, 6, false, x)
  }

  */

  lwf.FunkyTriangles()

  lwf.Save(dest + lwf.GetName() + ".svg")
}