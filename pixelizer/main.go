package svgr

import (
  "github.com/gographics/imagick/imagick"
  "io/ioutil"
  "sync"
  // "fmt"
)

const (
  //Apply adaptive sharpening to shrunk images
  AdaptiveSharpenVal  float64   = 16 
  //Amount of randomness to apply to "Funky" pixelation methods
  Funkiness           int       = 6 
)

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
  wg        sync.WaitGroup
  wands
}

type pxAddress struct {
  row, column int
  rgb         []uint8
}

func NewPixelizr(img string, targetRes int) (pixelData, error) {

  reader, err := ioutil.ReadFile(img)
  if err != nil {
    panic(err.Error())
  }

  var wg sync.WaitGroup

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
    wg:        wg,
    wands:     intitializeWands(),
  }, err
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

func (pxd pixelData) pixelLooper(renderMethod func(chan pxAddress)) {

  idx := 0
  
  for row := 0; row < pxd.rows; row++ {

    for col := 0; col < pxd.columns; col++ {

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

      go renderMethod(pxChan)

      idx += 3
    }
  }

  pxd.wg.Wait()

  return
}

/**
 * Iterate through a exported pixel data, apply a render method to each pixel and save the result
 * @param {func} renderMethod - the method used to render individual pixels
 * @param {string} dest - the intended destination for the saved png.
 * @return error
*/
/*func (pxd pixelData) writePng(renderMethod func(chan<- pxAddress), dest string) error {

  idx := 0
  
  for row := 0; row < pxd.rows; row++ {

    for col := 0; col < pxd.columns; col++ {

      // TODO pass in pixel array instead
      pxChan := buildPixelChannel(pxAddress {
        row:    row,
        column: col,
        idx:    idx,
      })

      render := renderMethod(pxChan)

      <-render

      // pixelChannel <- renderMethod(pxAddress {
      //   row:    row,
      //   column: col,
      //   idx:    idx,
      // })

      idx += 3
    }
  }


  // Save image to dest
  bg := imagick.NewPixelWand()
  mw := imagick.NewMagickWand()
  mw.NewImage(1920,1080,bg)
  mw.SetImageFormat("png")
  mw.DrawImage(pxd.wands.dw)
  mw.SetAntialias(false)
  err := mw.WriteImage(dest)

  return err
}*/

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