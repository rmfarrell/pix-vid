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

var vid string = "./src/betrayed.mp4"


func worker(files <-chan string, results chan<- []uint8) {

  for file := range files {
    results <- pixelizr.ReadImage(file)
  }
}

func main() {

  pxs := make(chan []uint8, 200)
  jobs := make(chan string, 500)

  imgSequence := media_converter.NewImageSequence(vid)

  imgFiles := imgSequence.GetFiles()

  for i := 0; i < 25; i ++ {
    go worker(jobs, pxs)
  }

  for j := 0; j < len(imgFiles); j++ {
    jobs <- imgFiles[j]
  }
  close(jobs)

  imgSequence.Clean()
}