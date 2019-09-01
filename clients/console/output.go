package console

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Output implements the io.Writer interface. We use it to redraw the console.
type Output struct{}

// Write redraws the screen with the contents of p.
func (_ *Output) Write(p []byte) (int, error) {
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "linux":
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	case "windows":
		c := exec.Command("cmd", "/c", "cls")
		c.Stdout = os.Stdout
		c.Run()
	default:
		// When the operating system is not supported we simply will
		// not clear the console and just print a continuous stream of
		// Game states instead.
	}

	return fmt.Println(string(p))
}
