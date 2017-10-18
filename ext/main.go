package main

import (
	"C"

	"github.com/endotakuya/ires/ext/ires"
)

func init() {}
func main() {}

//export resizeImage
func resizeImage(Uri *C.char, width, height int, Dir, Expire *C.char) *C.char {
	uri		:= C.GoString(Uri)
	dir 	:= C.GoString(Dir)
	expire	:= C.GoString(Expire)

	r := &ires.Ires{
		Uri: uri,
		Size: ires.Size{
			Width: width,
			Height: height,
		},
		Dir: dir,
		Expire: expire,
		IsLocal: false,
	}

	return C.CString(r.Resize())
}

//export cropImage
func cropImage(Uri *C.char, width, height int, Dir, Expire *C.char) *C.char {
	uri		:= C.GoString(Uri)
	dir 	:= C.GoString(Dir)
	expire	:= C.GoString(Expire)

	r := &ires.Ires{
		Uri: uri,
		Size: ires.Size{
			Width: width,
			Height: height,
		},
		Dir: dir,
		Expire: expire,
		IsLocal: false,
	}

	return C.CString(r.Crop())
}

//export resizeToCropImage
func resizeToCropImage(Uri *C.char, width, height int, Dir, Expire *C.char) *C.char {
	uri		:= C.GoString(Uri)
	dir 	:= C.GoString(Dir)
	expire	:= C.GoString(Expire)

	r := &ires.Ires{
		Uri: uri,
		Size: ires.Size{
			Width: width,
			Height: height,
		},
		Dir: dir,
		Expire: expire,
		IsLocal: false,
	}

	return C.CString(r.ResizeToCrop())
}
