package main

import (
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatalln(err)
	}

	rootWin := xproto.Setup(X).DefaultScreen(X).Root
	zero := int16(0)
	if err := xproto.WarpPointerChecked(X, rootWin, rootWin, zero, zero,
		uint16(0), uint16(0), zero, zero).Check(); err != nil {
		log.Fatalln(err)
	}
}
