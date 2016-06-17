package artifact

import (
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"os"
)

// Artifact is created for each piece generated in a summary
type Artifact struct {
	gorm.Model
	FileContents string // File path
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
		contents := []byte(art.FileContents)
		err := ioutil.WriteFile(filepath, contents, 0644)

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
