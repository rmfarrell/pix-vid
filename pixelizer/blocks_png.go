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
}


func (pxd pixelData) writePng(renderMethod func(int, int, int), dest string) {

  bg := imagick.NewPixelWand()

  idx := 0
  
  for row := 0; row < pxd.rows; row++ {

    for col := 0; col < pxd.columns; col++ {

      renderMethod(row, col, idx)

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