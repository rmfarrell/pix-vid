package svgr

import (
  "fmt"
  "github.com/gographics/imagick/imagick"
)

func (pxd pixelData) Hexagons(dest string, index int) error {

  err := pxd.pixelLooper(func(pxAddr chan pxAddress) {

    for pxa := range pxAddr {

      row  := float64(pxa.row)
      col  := float64(pxa.column)
      mult := float64(pxd.blockSize)

      // Stagger to create interlocking pixels
      var z float64 = staggerHexagons(pxa.column, mult)
    

      pxd.wands.pw.SetColor(fmt.Sprintf("#%x", pxa.rgb))
      pxd.wands.dw.SetFillColor(pxd.wands.pw)

      coords := []imagick.PointInfo {
        {
          X: col * mult - mult * .2,
          Y: row * mult + (mult/2) + z,
        },
        {
          X: col * mult + mult * .2,
          Y: row * mult + z,
        },
        {
          X: col * mult + mult * .8,
          Y: row * mult + z,
        },
        {
          X: col * mult + mult + mult * .2,
          Y: row * mult + (mult/2) + z,
        },
        {
          X: col * mult + mult * .8,
          Y: row * mult + mult + z,
        },
        {
          X: col * mult + mult * .2,
          Y: row * mult + mult + z,
        },
        {
          X: col * mult - mult * .2,
          Y: row * mult + (mult/2) + z,
        },
      }

      pxd.wands.dw.Polygon(coords)
    }
  }, dest)
  return err
}

func staggerHexagons(col int, multiplier float64) (stagger float64) {

  if col % 2 == 0 {
    stagger = multiplier/2
  }
  return -stagger
}