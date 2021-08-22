package fswebcam

import (
	"fmt"
	"os"
	"os/exec"
)

func TakePhoto(filename string) {
	cmd := exec.Command(
		"fswebcam",
		// Sets a 2-second delay
		"--delay",  "2",
		// Sets the resolution to 1920x1080
		"--resolution", "1920x1080",
		// Skips first 10 frames
		"--skip", "100",
		filename)

	// configure `Stdout` and `Stderr`
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	// run command
	if err := cmd.Run(); err != nil {
		fmt.Println( "Error:", err )
	}
}