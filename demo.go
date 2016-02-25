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
  maxFrames float64 = 9000
)

var vid string = "./src/example.mp4"

func worker(frames <-chan media_converter.Frame, results chan<- string) {

  for job := range frames {
    _dest := fmt.Sprintf("%s.png",job.GetPath())

    px, err := pixelizr.NewPixelizr(job.GetPath(), 40)
    if(err != nil) {
      panic(err.Error())
    }

    err = px.Triangles( _dest, job.GetIndex() )
    if (err != nil) {
      panic(err.Error())
    }

    fmt.Println("succcess!" + _dest)
    results <- _dest
  }
}

func main() {


  pngs := make(chan string, 20)

  imgSequence := media_converter.NewImageSequence(vid)

  frameCount := int(math.Min(maxFrames, float64(len(imgSequence.Files))))

  for i := 0; i < 8; i ++ {
    go worker(imgSequence.Files, pngs)
  }

  for a := 0; a < frameCount; a++ {
    <-pngs
  }

  imgSequence.ToMp4(fmt.Sprintf("%s%s.mp4", dest, imgSequence.GetID()))
  imgSequence.Clean()
}