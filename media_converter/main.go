package media_converter

import (
  "fmt"
  // "reflect"
  // "strconv"
  "os"
  "os/exec"
  "path/filepath"
  "github.com/satori/go.uuid"
  "io/ioutil"
)

const (
  srcDir string = "/tmp/dest/" // TODO: change to "/tmp/""
)

type imageSequence struct {
  id    uuid.UUID
  files []string
}

func NewImageSequence(videoPath string) imageSequence {

  _id := uuid.NewV4()

  return imageSequence {
    id: _id,
    files: videoToImages(videoPath, _id),
  }
}

// Separate each image in an animated gif and resave in a unique folder
// Create a read of each file in the directory
// Return an array of blobs of each image and the directory
func SeparateAnimatedGif(animated *os.File) (imageFiles [][]byte) {

  // Generate a UUID and make a directory with corresponding name
  dir := fmt.Sprintf("./%s", uuid.NewV4())
  if err := os.Mkdir(dir, 0777); err != nil {
    panic(err.Error())
  }

  // Separate and coalesce each frame of the animation into the new folder
  cmd := exec.Command(
    "convert", 
    "-coalesce", 
    animated.Name(), 
    fmt.Sprintf("./%s/%%05d.gif", dir),
  )
  cmd.Run()

  // Save a reference to each file in the directory
  files, _ := ioutil.ReadDir(dir)
  for _, f := range files {
    rf, _ := ioutil.ReadFile(dir + "/" + f.Name())
    imageFiles = append(imageFiles, rf)
  }

  /*
  // Clean up the temprorary directory once each image is stored in imageFiles blob
  if err := os.RemoveAll(dir); err != nil {
    panic(err.Error())
  }
  */

  return
}

func (sequence imageSequence) Clean() {

  for _, file := range sequence.files {
    err := os.Remove(file)
    if (err != nil) {
      panic(err.Error())
    }
  }

  return
}

func ImagesToVideo() {
  // ffmpeg -framerate 1 -pattern_type glob -i '*.jpg' -c:v libx264 out.mp4
  cmd := exec.Command(
    "ffmpeg",
    "-i",
    fmt.Sprintf("dest/output%%04d.jpg"),
    "-c:v",
    "libx264",
    "-vf",
    "fps=25",
    "-pix_fmt",
    "yuv420p",
    "out.mp4",
    // "ffmpeg",
    // "-f",
    // "image2",
    // "-s",
    // "1920x1080",
    // "-i",
    // fmt.Sprintf("dest/output%%04d.jpg"),
    // "-vcodec",
    // "libx264",
    // "-crf",
    // "15",
    // "test.mp4",
  )
  if err := cmd.Run(); err != nil {
    panic(err.Error())
  }
}

// Run ffmpeg to transform the video into individual jpgs
// Follows pattern /tmp/[uuid]-00001.jpg
func videoToImages(video string, id uuid.UUID) []string {

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

  return files
}