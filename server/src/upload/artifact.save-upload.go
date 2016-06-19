package upload

import (
	"github.com/ZombieHippie/test-gin/server/src/artifact"
	"github.com/ZombieHippie/test-gin/server/src/shared"
	"io"
	"log"
	"os"
	"path"
)

// SaveUpload saves the ArtifactUpload to a location on the harddisk and
// creates an Artifact in the database
func (upload *ArtifactUpload) SaveUpload(destPath string, data io.Reader, zipped bool) (artifact.Artifact, error) {
	art := artifact.Artifact{
		Path:          destPath,
		LocalPath:     upload.Path,
		PostProcessor: upload.PostProcessor,
		Label:         upload.Label,
	}

	// create its directory, unless it exists, those errors are OK
	mkerr := os.MkdirAll(path.Dir(art.Path), 0777)
	if mkerr != nil && !os.IsExist(mkerr) {
		log.Fatalln("Creating directory error!")
	}

	file, err := os.Create(art.Path)
	defer file.Close()

	if err != nil {
		log.Fatalln("error creating Artifact file.", err)
		return art, err
	}

	_, err = io.Copy(file, data)

	if err != nil {
		log.Fatalln("error copying Artifact into file.", err)
		return art, err
	}

	if zipped {
		zippedFileName := art.Path + ".zip"
		if err := os.Rename(art.Path, zippedFileName); err != nil {
			log.Fatalln("error renaming Artifact to *.zip file.", err)
			return art, err
		}

		if err := shared.Unzip(zippedFileName, art.Path); err != nil {
			log.Fatalln("error unzipping Artifact file.", err)
			return art, err
		}
	}

	return art, nil
}
