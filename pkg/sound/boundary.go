package sound

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func FindSounds(saveFolder string) []string {
	var soundList []string
	files, err := ioutil.ReadDir(saveFolder)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		soundList = append(soundList, filepath.Base(file.Name()))
	}
	return soundList
}

func RemoveAllSound(saveFolder string) {
	dir, _ := ioutil.ReadDir(saveFolder)
	for _, d := range dir {
		_ = os.RemoveAll(path.Join([]string{"saveFolder", d.Name()}...))
	}
}