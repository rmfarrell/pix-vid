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


func worker(jobs <-chan []byte, results chan<- []uint8) {
  for j := range jobs {
    fmt.Println("run")
    results <- pixelizr.ReadImage(j)
  }
}

func main() {

  // pxs := make(chan []uint8, 200)
  // jobs := make(chan []byte, 500)
  
  media_converter.VideoToAnimatedGif(file, 360, 180)

  // media_converter.SeparateAnimatedGif(gif)

  // media_converter.Cleanup()

  // pixelizr.ReadImage(imgs[0])

  // for i := 0; i < 25; i ++ {
  //   go worker(jobs, pxs)
  // }

  // for j := 0; j < len(imgs); j++ {
  //   jobs <- imgs[j]
  // }
  // close(jobs)
}