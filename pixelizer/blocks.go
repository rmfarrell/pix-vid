package svgr

import (
  "fmt"
)

func (pxd pixelData) Blocks(dest string) {


  pxd.pixelLooper(func(pxData chan pxAddress) {

    select {
    case data := <-pxData:

      row  := data.row
      col  := data.column
      mult := pxd.blockSize

      fmt.Println(row, col, mult, data.rgb)

      pxd.wands.pw.SetColor("#FFFFFF")
      pxd.wands.dw.SetFillColor(pxd.wands.pw)

      ox := float64(col*mult)
      oy := float64(row*mult)
      px := ox-float64(mult/3)
      py := oy-float64(mult/3)

      pxd.wands.dw.Circle(ox,oy,px,py)

      pxd.wg.Done()

    default:
      return
    }
  })

  return
}