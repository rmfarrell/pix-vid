package svgr

import (
  _ "image"
  "github.com/gographics/imagick/imagick"
  "io/ioutil"
)

const (
  //Apply adaptive sharpening to shrunk images
  AdaptiveSharpenVal  float64   = 16 
  //Amount of randomness to apply to "Funky" pixelation methods
  Funkiness           int       = 6 
)

// The actual body of the svg output
type svgContent struct {
  start, g, end string
}

type pixelArray struct {
  svgContent
  // An array of array of exported pixels.
  // Each parent array is a frame in an animation (a single frame is) pixelData[0]
  // The child arrays represent the rgb value of every pixel of the image
  pixelData   [][]uint8
  // Height and width of the image(s)
  w,h         int
  name        string
}

type wands struct {
  pw *imagick.PixelWand       
  mw *imagick.MagickWand
  dw *imagick.DrawingWand
}

type pixelData struct{
  data      []uint8
  rows      int
  columns   int
  blockSize int 
  wands
}

func NewPixelizr(img string, targetRes int) pixelData {

  reader, err := ioutil.ReadFile(img)
  if err != nil {
    panic(err.Error())
  }

  wand := imagick.NewMagickWand()

  if err := wand.ReadImageBlob(reader); err != nil {
    panic(err.Error())
  }

  width, height := shrinkImage(wand, targetRes)

  px, err := wand.ExportImagePixels(0,0,width,height,"RGB", imagick.PIXEL_CHAR)
  if err != nil {
    panic(err.Error())
  }

  return pixelData {
    data:      px.([]uint8),
    rows:      int(height),
    columns:   int(width),
    blockSize: int(1080/height),
    wands:     intitializeWands(),
  }
}

func intitializeWands() wands {

  return wands {
    pw: imagick.NewPixelWand(),
    mw: imagick.NewMagickWand(),
    dw: imagick.NewDrawingWand(),
  }
}

/*
* Private Methods
*/

// Shrink an image so that its longest dimension is no longer than maxSize
func shrinkImage(wand *imagick.MagickWand, maxSize int) (w,h uint) {

  w,h = getDimensions(wand)

  shrinkBy := 1

  if w >= h {
    shrinkBy = int(w)/maxSize
  } else {
    shrinkBy = int(h)/maxSize
  }

  wand.AdaptiveResizeImage(
    uint(int(w)/shrinkBy), 
    uint(int(h)/shrinkBy),
  )

  // Sharpen the image to bring back some of the color lost in the shrinking
  wand.AdaptiveSharpenImage(0,AdaptiveSharpenVal)

  w,h = getDimensions(wand)

  return
}

// Returns an the width and height of magick wand
func getDimensions(wand *imagick.MagickWand) (w,h uint) {
  h = wand.GetImageHeight()
  w = wand.GetImageWidth()
  return
}