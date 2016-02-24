package media_converter

import (
  "fmt"
  // "reflect"
  // "strconv"
  "os"
  "os/exec"
  "path/filepath"
  "github.com/satori/go.uuid"
  // "io/ioutil"
)

const (
  srcDir string = "/tmp/dest/" // TODO: change to "/tmp/""
)

type imageSequence struct {
  id    uuid.UUID
  Files chan string
}

func NewImageSequence(videoPath string) imageSequence {

  _id := uuid.NewV4()

  return imageSequence {
    id: _id,
    Files: videoToImages(videoPath, _id),
  }
}

func (sq imageSequence) GetLength() int {
  return len(sq.Files)
}

func removeFile(file string) error {
  err := os.Remove(file)
  if (err != nil) {
    panic(err.Error())
  }
  return err
}

func (sq imageSequence) Clean() {

  for file := range sq.Files {
    err := os.Remove(file)
    if (err != nil) {
      panic(err.Error())
    }
  }
}

func (sq imageSequence) Add(file string) {
  sq.Files <- file
  return
}

func (sq imageSequence) ToMp4(dest string) error {
  // ffmpeg -framerate 1 -pattern_type glob -i '*.jpg' -c:v libx264 out.mp4
  cmd := exec.Command(
    "ffmpeg",
    "-i",
    fmt.Sprintf("%s%s-%%06d.jpg", srcDir, sq.id),
    "-c:v",
    "libx264",
    "-vf",
    "fps=25",
    "-pix_fmt",
    "yuv420p",
    dest,
  )
  err := cmd.Run()

  return err
}

// Run ffmpeg to transform the video into individual jpgs
// Follows pattern /tmp/[uuid]-00001.jpg
func videoToImages(video string, id uuid.UUID) chan string {

  fileChan := make(chan string, 500)

  cmd := exec.Command(
    "ffmpeg",
    "-r",
    "25",
    "-i",
    video,
    "-f",
    "image2",
    fmt.Sprintf("%s%s-%%06d.jpg", srcDir, id),
  )
  if err := cmd.Run(); err != nil {
    panic(err.Error())
  }

  // files, err := filepath.Glob(srcDir + strconv(id))
  files, err := filepath.Glob(fmt.Sprintf("%s%s-*.jpg", srcDir, id))
  if (err != nil) {
    panic(err.Error())
  }

  for _, file := range files {
    fileChan <- file
  }
  close(fileChan)

  return fileChan
}