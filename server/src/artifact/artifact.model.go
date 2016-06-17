package artifact

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Artifact is created for each piece generated in a summary
type Artifact struct {
	gorm.Model
	FileContents string // File path
	FileName     string
	IsBinary     bool
	Data         string // Some JSON formatted data?
	Label        string
	File         bool
	Passed       int64
	Failed       int64
}

// SaveIntoFile saves the Artifact to a location on the harddisk
func (art *Artifact) SaveIntoFile(filepath string) {
	if filepath == "" {
		log.Fatalln("filepath not provided for save into file. Article not saved.")
		return
	}

	// create it, unless it exists, those errors are OK
	mkerr := os.MkdirAll(path.Dir(filepath), 0777)
	if mkerr != nil && !os.IsExist(mkerr) {
		log.Fatalln("Creating directory error!")
	}

	if art.File { // is already in file form
		// move if in different location
		if filepath != art.FileContents {
			currentlocation := art.FileContents

			err := os.Rename(currentlocation, filepath)
			if err != nil {
				log.Fatalln(err)
				return
			}

			art.FileContents = filepath
		}
	} else {
		var contents []byte
		var err error
		if art.IsBinary {
			contents, err = base64.StdEncoding.DecodeString(art.FileContents)
		} else {
			contents = []byte(art.FileContents)
		}

		if err != nil {
			log.Fatalln(err)
			return
		}

		err = ioutil.WriteFile(filepath, contents, 0644)

		if err != nil {
			log.Fatalln(err)
			return
		}

		art.FileContents = filepath
		art.File = true

	}
}

// ReadFile returns the file contents of this Artifact
func (art Artifact) ReadFile() ([]byte, error) {
	if art.File {
		filepath := art.FileContents
		return ioutil.ReadFile(filepath)
	}

	return []byte(art.FileContents), nil
}
