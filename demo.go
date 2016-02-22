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

var vid string = "./src/betrayed.mp4"


func worker(jobs <-chan string, results chan<- string) {

  for job := range jobs {
    results <- job
  }


  /*for j := range jobs {
    fmt.Println("run")
    results <- pixelizr.ReadImage(j)
  }*/
}

func main() {

  pxs := make(chan string, 200)
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

  // imgs.ToMp4("dest/out.mp4")

  imgSequence.Clean()
}