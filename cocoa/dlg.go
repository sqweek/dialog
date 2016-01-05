package cocoa

// #cgo darwin LDFLAGS: -framework Cocoa
// #include <stdlib.h>
// #include "dlg.h"
import "C"

import (
	"bytes"
	"errors"
	"unsafe"
)

type AlertParams struct {
	p C.AlertDlgParams
}

func mkAlertParams(msg, title string, style C.AlertStyle) *AlertParams {
	a := AlertParams{C.AlertDlgParams{msg: C.CString(msg), style: style}}
	if title != "" {
		a.p.title = C.CString(title)
	}
	return &a
}

func (a *AlertParams) run() C.DlgResult {
	return C.alertDlg(&a.p)
}

func (a *AlertParams) free() {
	C.free(unsafe.Pointer(a.p.msg))
	if a.p.title != nil {
		C.free(unsafe.Pointer(a.p.title))
	}
}

func nsStr(s string) unsafe.Pointer {
	return C.NSStr(unsafe.Pointer(&[]byte(s)[0]), C.int(len(s)))
}

func YesNoDlg(msg, title string) bool {
	a := mkAlertParams(msg, title, C.MSG_YESNO)
	defer a.free()
	return a.run() == C.DLG_OK
}

func ErrorDlg(msg, title string) {
	a := mkAlertParams(msg, title, C.MSG_ERROR)
	defer a.free()
	a.run()
}

func FileDlg(save int, title string, exts []string, relaxExt bool) (string, error) {
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
	if len(exts) > 0 {
		cext := make([]unsafe.Pointer, len(exts))
		for i, ext := range exts {
			i := i
			cext[i] = nsStr(ext)
			defer C.NSRelease(cext[i])
		}
		p.exts = (*unsafe.Pointer)(unsafe.Pointer(&cext[0]))
		p.numext = C.int(len(cext))
		if relaxExt {
			p.relaxext = 1;
		}
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
