package svgr

import (
  "github.com/gographics/imagick/imagick"
  "math/rand"
)

func randomizePoint(originalPoint float64, multiplier float64) float64 {

  var (
    highOffset float64 = 4
    lowOffset  float64 = -4
  )

  offset := rand.Float64() * highOffset + lowOffset

  return originalPoint + offset
}

func (pxd pixelData) FunkySquares(dest string, index int) error {

  err := pxd.pixelLooper(func(pxAddr chan pxAddress) {
    for pxa := range pxAddr {

      row  := float64(pxa.row)
      col  := float64(pxa.column)
      mult := float64(pxd.blockSize)
      
      pxd.wands.dw.SetFillColor(pxa.pixelWand)

      coords := []imagick.PointInfo {
        {
          X: randomizePoint(col * mult, mult),
          Y: randomizePoint(row * mult, mult),
        },
        {
          X: randomizePoint(col * mult + mult, mult),
          Y: randomizePoint(row * mult, mult),
        },
        {
          X: randomizePoint(col * mult + mult, mult),
          Y: randomizePoint(row * mult + mult, mult),
        },
        {
          X: randomizePoint(col * mult, mult),
          Y: randomizePoint(row * mult + mult, mult),
        },
      }

      pxd.wands.dw.Polygon(coords)
    }
  }, dest)
  return err
}