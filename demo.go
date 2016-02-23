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
  maxFrames float64 = 3
)

var vid string = "./src/richter-cut.mp4"


func worker(imgFiles []string, jobs <-chan int, results chan<- string) {

  for job := range jobs {
    _dest := fmt.Sprintf("%s.png",imgFiles[job])
    px := pixelizr.NewPixelizr(imgFiles[job], 90)
    px.BlocksPng(_dest)
    fmt.Println(fmt.Sprintf("succcess! %s", _dest))
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

  for x := 0; x < frames; x++ {
    imgSequence = imgSequence.Add(<-pngs)
  }

  go imgSequence.Clean()
}