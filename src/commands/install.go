package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"compose-generator/utils"

	"github.com/fatih/color"
)

// ---------------------------------------------------------------- Public functions ---------------------------------------------------------------

// Install Docker and Docker Compose with a single command
func Install(flagOnlyCompose bool, flagOnlyDocker bool) {
	const WindowsInstallerURL = "https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe"

	utils.ClearScreen()

	if runtime.GOOS == "windows" { // Running on windows
		// Download Docker installer
		fmt.Print("Downloading Docker installer ...")
		filePath := os.TempDir() + "/DockerInstaller.exe"
		err := utils.DownloadFile(WindowsInstallerURL, filePath)
		if err != nil {
			panic(err)
		}
		utils.PrintDone()
		// Run Docker installer
		utils.Pl("Starting installation ...")
		utils.Pel()
		cmd := exec.Command(filePath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		utils.Pel()
	} else { // Running on linux
		if utils.IsPrivileged() {
			// Install lsb_release if not installed already
			utils.ExecuteAndWait("apt-get", "update")

			// Install Docker
			if !flagOnlyCompose {
				utils.ExecuteAndWait("curl", "-fsSL", "https://get.docker.com", "|", "sh")
			}

			// Install Docker Compose
			if !flagOnlyDocker {
				cmd := exec.Command("uname", "-s")
				output, _ := cmd.CombinedOutput()
				unameS := strings.TrimRight(string(output), "\r\n")

				cmd = exec.Command("uname", "-m")
				output, _ = cmd.CombinedOutput()
				unameM := strings.TrimRight(string(output), "\r\n")

				utils.ExecuteAndWaitWithOutput("curl", "-L", "https://github.com/docker/compose/releases/download/1.28.2/docker-compose-"+unameS+"-"+unameM, "-o", "/usr/local/bin/docker-compose")
				utils.ExecuteAndWaitWithOutput("chmod", "+x", "/usr/local/bin/docker-compose")
			}
		} else {
			color.Red("Please execute this command with root privileges. The cli is not able to install Docker and Docker Compose without those privileges.")
		}
	}

	// Check if installation was successful
	if utils.CommandExists("docker") {
		cmd := exec.Command("docker", "-v")
		dockerVersion, _ := cmd.CombinedOutput()
		cmd = exec.Command("docker-compose", "-v")
		composeVersion, _ := cmd.CombinedOutput()
		if flagOnlyCompose {
			color.Yellow("Congrats! You have installed " + strings.TrimRight(string(composeVersion), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
		} else if flagOnlyDocker {
			color.Yellow("Congrats! You have installed " + strings.TrimRight(string(dockerVersion), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
		} else {
			color.Yellow("Congrats! You have installed " + strings.TrimRight(string(dockerVersion), "\r\n") + " and " + strings.TrimRight(string(composeVersion), "\r\n") + ". You now can start by executing 'compose-generator generate' to generate your compose file.")
		}
	} else {
		color.Red("An error occurred while installing Docker.")
	}
}

// --------------------------------------------------------------- Private functions ---------------------------------------------------------------
