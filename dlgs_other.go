//go:build !windows && !(linux && !android) && !darwin

package dialog

func (b *MsgBuilder) yesNo() bool {
	return true
}

func (b *MsgBuilder) info() {
}

func (b *MsgBuilder) error() {
}

func (b *FileBuilder) load() (string, error) {
	return b.run(false)
}

func (b *FileBuilder) save() (string, error) {
	return b.run(true)
}

func (b *FileBuilder) run(save bool) (string, error) {
	return "", nil
}

func (b *DirectoryBuilder) browse() (string, error) {
	return "", nil
}
