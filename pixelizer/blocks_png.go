package svgr

import (
  "fmt"
  "github.com/gographics/imagick/imagick"
)

func (pxd pixelData) BlocksPng(dest string) {


  pxd.writePng(func(row, col, idx int) {

      mult := pxd.blockSize

      pxd.wands.pw.SetColor(fmt.Sprintf("#%x",[]uint8{
        pxd.data[idx],
        pxd.data[idx+1],
        pxd.data[idx+2],
      }))
      pxd.wands.dw.SetFillColor(pxd.wands.pw)

      ox := float64(col*mult)
      oy := float64(row*mult)
      px := ox-float64(mult/3)
      py := oy-float64(mult/3)

      pxd.wands.dw.Circle(ox,oy,px,py)

    }, dest)

  /*
  mult := 1080 / pxd.rows
  dw   := imagick.NewDrawingWand()
  bg   := imagick.NewPixelWand()
  pw   := imagick.NewPixelWand()

  bg.SetColor("#000000")

  i := 0

  // Iterate Over Rows
  
  for row := 0; row < pxd.rows; row++ {

    for col := 0; col < pxd.columns; col++ {

      pw.SetColor(fmt.Sprintf("#%x",[]uint8{
        pxd.data[i],
        pxd.data[i+1],
        pxd.data[i+2],
      }))
      dw.SetFillColor(pw)

      ox := float64(col*mult)
      oy := float64(row*mult)
      px := ox-float64(mult/3)
      py := oy-float64(mult/3)

      dw.Circle(ox,oy,px,py)

      i = i+3
    }
  }

  mw := imagick.NewMagickWand()
  mw.NewImage(1920,1080,bg)
  mw.SetImageFormat("png")
  mw.DrawImage(dw)
  mw.SetAntialias(false)
  mw.WriteImage(dest)
  */
}


func (pxd pixelData) writePng(renderMethod func(int, int, int), dest string) {

  bg := imagick.NewPixelWand()

  idx := 0
  
  for row := 0; row < pxd.rows; row++ {

    for col := 0; col < pxd.columns; col++ {

      renderMethod(row, col, idx)
      /*
      pw.SetColor(fmt.Sprintf("#%x",[]uint8{
        pxd.data[i],
        pxd.data[i+1],
        pxd.data[i+2],
      }))
      dw.SetFillColor(pw)

      ox := float64(col*mult)
      oy := float64(row*mult)
      px := ox-float64(mult/3)
      py := oy-float64(mult/3)

      dw.Circle(ox,oy,px,py)
      */

      idx += 3
    }
  }

  mw := imagick.NewMagickWand()
  mw.NewImage(1920,1080,bg)
  mw.SetImageFormat("png")
  mw.DrawImage(pxd.wands.dw)
  mw.SetAntialias(false)
  mw.WriteImage(dest)
}

/*
func writeGroup(pxa *pixelArray, frameIndex int, renderMethod func([]uint8, int, int) string) {

  fmt.Printf("writing <g> %d...", frameIndex + 1)
  
  pxa.svgContent.g += fmt.Sprintf("<g id=\"f%d\">", frameIndex)

  i := 0

  // Iterate over rows
  for row := 0; row < pxa.h; row++ {

    // Iterate over columns
    for col := 0; col < pxa.w; col++ {

      // Iterate through each pixel of the image (input) and call the renderMethod on it.
      pxa.svgContent.g += renderMethod(
        []uint8{
          pxa.pixelData[frameIndex][i], 
          pxa.pixelData[frameIndex][i+1], 
          pxa.pixelData[frameIndex][i+2],
        },
        col,
        row,
      )

      i = i+3
    }
  }

  pxa.svgContent.g += "</g>"

  fmt.Println("success!")

  return
}
*/