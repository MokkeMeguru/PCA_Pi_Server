package sound

import (
	"fmt"
	"os"
	"os/exec"
)

var command = "wget"

func init() {
	if !isCommandAvailable(command) {
		_ = fmt.Errorf("%s is not available", command)
	}
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command(name, "--help")
	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd is not found: %s", cmd)
		return false
	}
	return true
}

func AssetDownload(url string, saveTo string) bool {
	opts := []string{"-o", saveTo, url}
	cmd := exec.Command(command, opts...)
	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd cannot be completed\n")
		os.RemoveAll(saveTo)
		return false
	}
	return true
}

func SaveSound(s *Sound) bool {
	saveTo := saveFolder + s.SoundName
	return AssetDownload(s.SoundSource, saveTo)
}
