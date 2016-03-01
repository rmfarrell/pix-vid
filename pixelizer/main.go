package svgr

import (
  "github.com/gographics/imagick/imagick"
  // "io/ioutil"
  // "time"
  // "fmt"
)

const (
  //Apply adaptive sharpening to shrunk images
  AdaptiveSharpenVal  float64   = 16 
  //Amount of randomness to apply to "Funky" pixelation methods
  Funkiness           int       = 6 
)

type wands struct {
  src *imagick.MagickWand
  pw  *imagick.PixelWand       
  mw  *imagick.MagickWand
  dw  *imagick.DrawingWand
}

type pixelData struct{
  data      []uint8
  rows      int
  columns   int
  blockSize int
  wands
}

type pxAddress struct {
  row, column int
  pixelWand   *imagick.PixelWand 
}

func NewPixelizr(img string, targetRes int) (pixelData, error) {
  
  srcWand := imagick.NewMagickWand()

  err := srcWand.ReadImage(img)

  width, height := shrinkImage(srcWand, targetRes)

  return pixelData {
    // data:      px.([]uint8),
    rows:      int(height),
    columns:   int(width),
    blockSize: int(1080/height),
    wands:     wands {
                 src: srcWand,
                 pw:  imagick.NewPixelWand(),
                 mw:  imagick.NewMagickWand(),
                 dw:  imagick.NewDrawingWand(),
               },
  }, err
}

func (pxd pixelData) Clean() () {
  pxd.src.Destroy()
  pxd.pw.Destroy()
  pxd.mw.Destroy()
  pxd.dw.Destroy()
}

// Create a channel with relevant pixel data for pixelLooper
func buildPixelChannel(addresses ...pxAddress) chan pxAddress {
  out := make(chan pxAddress)
  go func(){

    for _, pixel := range addresses {
      out <- pixel
    }
    close(out)
  }()
  return out
}

/**
 * Iterate through a exported pixel data, apply a render method to each pixel and save the result
 * @param {func} renderMethod - the method used to render individual pixels
 * @param {string} dest - the intended destination for the saved png.
 * @return error
*/
func (pxd pixelData) pixelLooper(renderMethod func(chan pxAddress), dest string) (err error) {

  pi := pxd.wands.src.NewPixelIterator()

  pi.SetFirstIteratorRow()

  for row := 0; row < pxd.rows; row++ {

    pi.GetNextIteratorRow()

    for col := 0; col < pxd.columns; col++ {

      pxChan := buildPixelChannel(pxAddress {
          row:       row,
          column:    col,
          pixelWand: pi.GetCurrentIteratorRow()[col],
        })

      renderMethod(pxChan)
    }

    err = pi.SyncIterator()
  }

  err = pxd.save(dest)

  /*
  idx := 0
  
  for row := 0; row < pxd.rows; row++ {

    for col := 0; col < pxd.columns; col++ {

      // Throttle to prevent cgo error
      // time.Sleep(60000)

      // TODO pass in pixel array instead
      pxChan := buildPixelChannel(pxAddress {
        row:    row,
        column: col,
        rgb:    []uint8 {
                  pxd.data[idx],
                  pxd.data[idx+1],
                  pxd.data[idx+2],
                },
      })

      renderMethod(pxChan)

      idx += 3
    }
  }

  err := pxd.save(dest)

  return err
  */

  return err
}

func (pxd pixelData) save(dest string) error {
  
  bg := imagick.NewPixelWand()
  bg.SetColor("#E82A33")
  mw := imagick.NewMagickWand()
  mw.NewImage( uint(pxd.blockSize*pxd.columns) ,1080,bg)
  mw.SetImageFormat("png")
  mw.DrawImage(pxd.wands.dw)
  mw.SetAntialias(false)
  err := mw.WriteImage(dest)

  return err
}

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