package media_converter

import (
  "fmt"
  // "reflect"
  "os"
  "os/exec"
  "github.com/satori/go.uuid"
  "io/ioutil"
)


func CreatePix() {
  
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

func Cleanup() {

  // Remove the temporary video
  err := os.Remove("./dest/test.gif")
  if (err != nil) {
    panic(err.Error())
  }
}

// Run ffmpeg to transform the video into an animated gif
func VideoToAnimatedGif(video string, width, height int) {

  // TODO: figure out a real tmp dir
  // dest := "./dest/test.gif"


  // ffmpeg -i "input.mov" -an -f image2 "output_%05d.jpg"
  cmd := exec.Command(
    "ffmpeg",
    "-i",
    video,
    "-an",
    "-f",
    "image2",
    fmt.Sprintf("dest/output%%04d"),
  )
  if err := cmd.Run(); err != nil {
    panic(err.Error())
  }

  // Read th
  /*reader, err := os.Open(dest)
  if err != nil {
    panic(err.Error())
  }
  defer reader.Close()

  return reader*/
}