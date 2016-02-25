package svgr

import (
  "fmt"
  "github.com/gographics/imagick/imagick"
)

func (pxd pixelData) Triangles(dest string, index int) error {

  err := pxd.pixelLooper(func(pxAddr chan pxAddress) {

    for pxa := range pxAddr {

      row  := float64(pxa.row)
      col  := float64(pxa.column)
      mult := float64(pxd.blockSize)

      pxd.wands.pw.SetColor(fmt.Sprintf("#%x", pxa.rgb))
      pxd.wands.dw.SetFillColor(pxd.wands.pw)

      var coords []imagick.PointInfo

      // Invert triangles each row, so they interlock
      if (pxa.row % 2 == 0) {

        // Down-pointing triangle
        coords = []imagick.PointInfo {
          {
            X: col * mult - mult * .5,
            Y: row * mult,
          },
          {
            X: col * mult + mult * .5,
            Y: row * mult,
          },
          {
            X: col * mult,
            Y: row * mult + mult * 2,
          },
        }

      } else {

        // Up-pointing triangle
        coords = []imagick.PointInfo {
          {
            X: col * mult,
            Y: row * mult + mult,
          },
          {
            X: col * mult + mult,
            Y: row * mult + mult,
          },
          {
            X: col * mult + mult * .5,
            Y: row * mult - mult,
          },
        }
      }

      pxd.wands.dw.Polygon(coords)
    }
  }, dest)
  return err
}