package main

import (
  "fmt"
  "math"
  // "reflect"
  // "os"
  media_converter "./media_converter"
  pixelizr "./pixelizer"
)

const (
  dest      string  = "./dest/"
  src       string  = "./src/"
  maxFrames float64 = 10
)

var vid string = "./src/richter-cut.mp4"


func worker(jobs <-chan string, results chan<- string) {

  for job := range jobs {
    _dest := fmt.Sprintf("%s.png",job)

    px, err := pixelizr.NewPixelizr(job, 60)
    if(err != nil) {
      panic(err.Error())
    }

    err = px.BlocksPng(_dest)
    if (err != nil) {
      panic(err.Error())
    }

    fmt.Println(fmt.Sprintf("succcess! %s", job))
    results <- _dest
  }
}

func main() {

  // jobs := make(chan string, 500)
  pngs := make(chan string, 500)

  imgSequence := media_converter.NewImageSequence(vid)

  frames := int(math.Min(maxFrames, float64(len(imgSequence.Files))))

  for i := 0; i < 6; i ++ {
    go worker(imgSequence.Files, pngs)
  }

  for a := 0; a < frames; a++ {
    <-pngs
  }

  imgSequence.Clean()
}