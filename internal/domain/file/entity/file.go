package entity

// File struct defines the file entity.
type TxFile struct {
	// ID is the ID of the file.
	Name string
	// Path is the path of the file.
	Path string
	// Hash is the hash of the file.
	Hash string
	// Lines is the number of lines in the file.
	Lines int64
	//
}

// NewTxFile returns a new TxFile instance.
func NewTxFile(name, path, hash string, lines int64) (file *TxFile) {
	file = &TxFile{
		Name:  name,
		Path:  path,
		Hash:  hash,
		Lines: lines,
	}
	return
}
