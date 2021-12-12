package shellUtils

import (
	"bytes"
	"os"
	"os/exec"
)

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var cmd *exec.Cmd
	if os.PathSeparator == '/' {
		cmd = exec.Command(ShellToUseUnix, "-c", command)
	} else {
		cmd = exec.Command(ShellToUseWin, "/C", command)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// GetGitStats function will return the git stats in the following format:
// "d8e6e45 \n d8e6e45d52f7bf164a995e22abb81ffc6e3eeae1 \n 3 0"
func GetGitStats() string {
	stdout, _, err := Shellout(gitCmd)
	if err != nil || len(stdout) == 0 {
		return ""
	}

	return stdout
}

func RestartBot(isWindows bool) {
	if !isWindows {
		startProcess(ShellToUseUnix, "-c", "./run.sh &")
	} else {
		startProcess(ShellToUseWin, "/C", "./Upgrade.bat")
	}
}

func startProcess(args ...string) (p *os.Process, err error) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}
