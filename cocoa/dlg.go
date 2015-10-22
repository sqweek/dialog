package cocoa

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <objc/NSObjcRuntime.h>
// #include <stdlib.h>
// #include "dlg.h"
import "C"

import (
	"bytes"
	"errors"
	"unsafe"
)

func YesNoDlg(msg string, title string) bool {
	p := C.AlertDlgParams{
		msg: C.CString(msg),
	}
	defer C.free(unsafe.Pointer(p.msg))
	if title != "" {
		p.title = C.CString(title)
		defer C.free(unsafe.Pointer(p.title))
	}
	return C.alertDlg(&p) == C.DLG_OK
}

func FileDlg(save int, title string) (string, error) {
	buf := make([]byte, 1024)
	p := C.FileDlgParams{
		save: C.int(save),
		buf: (*C.char)(unsafe.Pointer(&buf[0])),
		nbuf: C.int(cap(buf)),
	}
	if title != "" {
		p.title = C.CString(title)
		defer C.free(unsafe.Pointer(p.title))
	}
	switch C.fileDlg(&p) {
	case C.DLG_OK:
		return string(buf[:bytes.Index(buf, []byte{0})]), nil
	case C.DLG_CANCEL:
		return "", nil
	case C.DLG_URLFAIL:
		return "", errors.New("failed to get file-system representation for selected URL")
	}
	panic("unhandled case")
}
