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
  maxFrames float64 = 100
)

var vid string = "./src/betrayed.mp4"


func worker(imgFiles []string, jobs <-chan int, results chan<- string) {

  for job := range jobs {
    _dest := fmt.Sprintf("%sp-%d.png",dest,job)
    px := pixelizr.NewPixelizr(imgFiles[job], 60)
    px.BlocksPng(_dest)
    fmt.Println(fmt.Sprintf("succcess! %d", job))
    results <- _dest
  }
}

func main() {

  jobs := make(chan int, 500)
  pngs := make(chan string, 500)

  imgSequence := media_converter.NewImageSequence(vid)

  imgFiles := imgSequence.GetFiles()

  frames := int(math.Min(maxFrames, float64(len(imgFiles))))

  for i := 0; i < 20; i ++ {
    go worker(imgFiles, jobs, pngs)
  }

  for j := 0; j < frames; j++ {
    jobs <- j
  }
  close(jobs)

  for a := 0; a < frames; a++ {
    <-pngs
  }

  imgSequence.Clean()
}