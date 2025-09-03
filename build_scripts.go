package main

import (
	"os"
	"os/exec"
)

func buildScript(commands []string) error {
	var scriptExt, runCmd string
	var cmdArgs []string

	tmpFile, err := os.CreateTemp("", "temp_script_*")
	if err != nil {
	}
	defer os.Remove(tmpFile.Name())

	switch SYS_OS {
	case "windows":
		scriptExt = ".ps1"
		runCmd = "powershell"
		cmdArgs = []string{"-ExecutionPolicy", "Bypass", "-File", tmpFile.Name() + scriptExt}
	case "linux", "mac":
		scriptExt = ".sh"
		runCmd = "bash"
		cmdArgs = []string{tmpFile.Name() + scriptExt}
	default:
	}

	scriptPath := tmpFile.Name() + scriptExt
	if err := os.Rename(tmpFile.Name(), scriptPath); err != nil {
	}

	scriptContent := ""
	for _, cmd := range commands {
		scriptContent += cmd + "\n"
	}

	if err := os.WriteFile(scriptPath, []byte(scriptContent), 0700); err != nil {
	}

	// Run scripts
	cmd := exec.Command(runCmd, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
	}

	return nil
}
