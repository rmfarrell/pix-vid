package svgr

import (
  "fmt"
)

func (pxd pixelData) BlocksPng(dest string) error {


  err := pxd.writePng(func(row, col, idx int) {

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

  return err
}