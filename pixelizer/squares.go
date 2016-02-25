package svgr

import (
  "fmt"
  "github.com/gographics/imagick/imagick"
)

func (pxd pixelData) Squares(dest string, index int) error {

  err := pxd.pixelLooper(func(pxAddr chan pxAddress) {

    for pxa := range pxAddr {

      row  := float64(pxa.row)
      col  := float64(pxa.column)
      mult := float64(pxd.blockSize)

      pxd.wands.pw.SetColor(fmt.Sprintf("#%x", pxa.rgb))
      pxd.wands.dw.SetFillColor(pxd.wands.pw)

      coords := []imagick.PointInfo {
        {
          X: col * mult,
          Y: row * mult,
        },
        {
          X: col * mult + mult,
          Y: row * mult,
        },
        {
          X: col * mult + mult,
          Y: row * mult + mult,
        },
        {
          X: col * mult,
          Y: row * mult + mult,
        },
      }

      pxd.wands.dw.Polygon(coords)
    }
  }, dest)
  return err
}