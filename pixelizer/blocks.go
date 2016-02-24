package svgr

import (
  "fmt"
)

func (pxd pixelData) Blocks(dest string) {


  pxd.pixelLooper(func(pxData chan pxAddress) {

    for px := range pxData {
      fmt.Println(px)
    }
  })

  // err := pxd.pixelLooper(func(chan<- pxAddress) {

  //   fmt.Println("test")
  //   return 
  // })

  return
}