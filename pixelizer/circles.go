package svgr

func (pxd pixelData) Circles(dest string, index int) error {

  err := pxd.pixelLooper(func(pxAddr chan pxAddress) {

    for pxa := range pxAddr {

      row  := pxa.row
      col  := pxa.column
      mult := pxd.blockSize

      pxd.wands.dw.SetFillColor(pxa.pixelWand)

      ox := float64(col*mult)
      oy := float64(row*mult)
      px := ox-float64(mult/3)
      py := oy-float64(mult/3)

      pxd.wands.dw.Circle(ox,oy,px,py)
    }
  }, dest)

  return err
}