package svgr

import (
  "fmt"
  "math"
)

const (
  baseSize       float64 = 2
  sequenceLength int     = 50
  rangeModifier  float64 = 12.5
)

// TODO: DRY this out; consolidate w/Circles
func (pxd pixelData) FreakyCircles(dest string, index int) error {

  var circleSize float64

  // Circles get slightly larger and smaller by index according to this formula
  circleSize = math.Abs(float64(index % sequenceLength) - float64(sequenceLength/2)) / rangeModifier + baseSize

  err := pxd.pixelLooper(func(pxAddr chan pxAddress) {

    for pxa := range pxAddr {

      row  := pxa.row
      col  := pxa.column
      mult := pxd.blockSize

      pxd.wands.pw.SetColor(fmt.Sprintf("#%x", pxa.rgb))
      pxd.wands.dw.SetFillColor(pxd.wands.pw)

      ox := float64(col*mult)
      oy := float64(row*mult)
      px := ox-float64(mult) / circleSize
      py := oy-float64(mult) / circleSize

      pxd.wands.dw.Circle(ox,oy,px,py)
    }
  }, dest)

  return err
}