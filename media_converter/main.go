package videohandler

// Separate each image in an animated gif and resave in a unique folder
// Create a read of each file in the directory
// Return an array of blobs of each image and the directory
func separateAnimatedGif(animated *os.File) (imageFiles [][]byte) {

  // Generate a UUID and make a directory with corresponding name
  dir := fmt.Sprintf("./_%s", uuid.NewV4())
  if err := os.Mkdir(dir, 0777); err != nil {
    panic(err.Error())
  }

  // Separate and coalesce each frame of the animation into the new folder
  cmd := exec.Command(
    "convert", 
    "-coalesce", 
    animated.Name(), 
    fmt.Sprintf("./%s/image_%%03d.gif", dir),
  )
  cmd.Run()

  // Save a reference to each file in the directory
  files, _ := ioutil.ReadDir(dir)
  for _, f := range files {
    rf, _ := ioutil.ReadFile(dir + "/" + f.Name())
    imageFiles = append(imageFiles, rf)
  }

  // Clean up the temprorary directory once each image is stored in imageFiles blob
  if err := os.RemoveAll(dir); err != nil {
    panic(err.Error())
  }

  return
}

// TODO
func videoToAnimatedGif(video *os.File) *os.File {
  return video
}