package dialog

import (
	"github.com/sqweek/dialog/cocoa"
)

func (b *MsgBuilder) yesNo() bool {
	return cocoa.YesNoDlg(b.Msg, b.Dlg.Title)
}

func (b *FileBuilder) load() (string, error) {
	return b.run(0)
}

func (b *FileBuilder) save() (string, error) {
	return b.run(1)
}

func (b *FileBuilder) run(save int) (string, error) {
	f, err := cocoa.FileDlg(save, b.Dlg.Title)
	if f == "" && err == nil {
		return "", Cancelled
	}
	return f, err
}
