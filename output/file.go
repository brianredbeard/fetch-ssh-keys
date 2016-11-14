package output

import (
	"io/ioutil"
	"os"
)

// FileWriter writes the output to file
type FileWriter struct {
	targetFile string
	fileMode   os.FileMode
}

// NewFileWriter creates new FileWriter what writes to targetFile
func NewFileWriter(targetFile string) *FileWriter {
	return &FileWriter{
		targetFile: targetFile,
		fileMode:   0600,
	}
}

func (w *FileWriter) write(output string) error {
	return ioutil.WriteFile(w.targetFile, []byte(output), w.fileMode)
}
