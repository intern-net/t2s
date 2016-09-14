package say

import "os/exec"

// Text2Speech writes wav binary.
func Text2Speech(filename, text, speaker string) error {
	cmd := exec.Command("say", "--file-format", "m4af", "-o", filename, text)
	return cmd.Run()
}
