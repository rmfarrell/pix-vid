package svgr

import (
  // "fmt"
)

type RGBChannel struct {
  offset       float64
  size         float64
  opacity      float64
  fill         string
  stroke       uint8
  strokeColor  string
}

type ChannelSettings struct {
  red, green, blue RGBChannel
}


// TODO: pass mult as arg
func (pxd pixelData) MultiChannelCircles(dest string, index int) error {

  cs := ChannelSettings {

    red: RGBChannel {
      offset:      0,
      opacity:     .2,
      fill:        "#e68c9b",
      stroke:      4,
      strokeColor: "#e68c9b",
    },
    green: RGBChannel {
      offset:  -3,
      opacity: .7,
      fill:    "#9AE68C",
    },
    blue: RGBChannel {
      offset:  10,
      opacity: .5,
      fill:    "#8CC7E6",
    },
  }

  rgbChannels := []RGBChannel {cs.red, cs.blue, cs.green}

  err := pxd.pixelLooper(func(pxAddr chan pxAddress) {

    for pxa := range pxAddr {

      row  := float64(pxa.row)
      col  := float64(pxa.column)
      mult := float64(pxd.blockSize)

      for idx, channel := range rgbChannels {

        // Calculate how large each cirlce should be
        circleSize := mult / (float64(pxa.rgb[idx]) / 255 * 15 + 3)

        pxd.wands.pw.SetColor(channel.fill)
        // TODO
        // pxd.wands.dw.SetStrokeWidth(0)
        // pxd.wands.dw.SetStrokeColor(channel.stroke)
        pxd.wands.pw.SetOpacity(channel.opacity)
        pxd.wands.dw.SetFillColor(pxd.wands.pw)

        ox := float64(col*mult + channel.offset)
        oy := float64(row*mult + channel.offset)
        px := ox-float64(circleSize)
        py := oy-float64(circleSize)

        pxd.wands.dw.Circle(ox,oy,px,py)
      }
    }
  }, dest)

  return err
}