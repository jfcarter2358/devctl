package main

import (
	"devctl/container"
	"fmt"
	"os"
)

var Version = "1.0.0"

func main() {

	showHelp := false
	command := ""

	idx := 1
	for idx < len(os.Args) {
		if os.Args[idx] == "-h" || os.Args[idx] == "--help" {
			showHelp = true
			idx += 1
			continue
		}
		command = os.Args[idx]
		idx += 1
	}

	if showHelp {
		printHelp(command)
		os.Exit(0)
	}

	switch command {
	case "init":
		doInit()
	case "up":
		doUp()
	case "shell":
		doShell()
	case "version":
		doVersion()
	}

}

func printError(err error) {
	fmt.Println(err.Error())
}

func printHelp(command string) {
	switch command {
	case "":
		printMainHelp()
	case "init":
		printInitHelp()
	case "up":
		printUpHelp()
	case "shell":
		printShellHelp()
	case "version":
		printVersionHelp()
	}
}

func printMainHelp() {
	help := `usage: %v [options] command
    options:
        -h, --help      Show this help message and exit
    arguments:
        init            Write out a template .devctl.json file
        up              Re-open this directory in the container defined in the .devctl.json file
        shell           Spawn a shell in the running container associated with this directory
        version         Show the devctl version and exit
`
	fmt.Printf(help, os.Args[0])
}

func printInitHelp() {
	help := `usage: %v init [options]
    options:
        -h, --help      Show this help message and exit
`
	fmt.Printf(help, os.Args[0])
}

func printUpHelp() {
	help := `usage: %v up [options]
    options:
        -h, --help      Show this help message and exit
`
	fmt.Printf(help, os.Args[0])
}

func printShellHelp() {
	help := `usage: %v shell [options]
    options:
        -h, --help      Show this help message and exit
`
	fmt.Printf(help, os.Args[0])
}

func printVersionHelp() {
	help := `usage: %v shell [options]
    options:
        -h, --help      Show this help message and exit
`
	fmt.Printf(help, os.Args[0])
}

func doInit() {
	container.DumpContainerJSON()
	os.Exit(0)
}

func doUp() {
	c, err := container.LoadContainer()
	if err != nil {
		fmt.Println(err)
	}
	c.LaunchContainer()
	c.RemoveContainer()
}

func doShell() {
	c, err := container.LoadContainer()
	if err != nil {
		fmt.Println(err)
	}
	c.ShellContainer()
}

func doVersion() {
	output := fmt.Sprintf("devctl version %v", Version)
	fmt.Println(output)
	os.Exit(0)
}
