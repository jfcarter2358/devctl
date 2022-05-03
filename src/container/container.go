package container

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type Container struct {
	Name   string   `json:"name" binding:"required"`
	Image  string   `json:"image" binding:"required"`
	Shell  string   `json:"shell" binding:"required"`
	User   string   `json:"user"`
	Mounts []string `json:"mounts"`
	Args   []string `json:"args"`
	Ports  []int    `json:"ports"`
}

func LoadContainer() (*Container, error) {
	if _, err := os.Stat(".devctl.json"); errors.Is(err, os.ErrNotExist) {
		return nil, os.ErrNotExist
	}

	// Open our jsonFile
	jsonFile, err := os.Open(".devctl.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var c *Container
	err = json.Unmarshal([]byte(byteValue), &c)

	if err != nil {
		return nil, err
	}

	if c.Ports == nil {
		c.Ports = make([]int, 0)
	}
	if c.Mounts == nil {
		c.Mounts = make([]string, 0)
	}
	if c.Args == nil {
		c.Args = make([]string, 0)
	}

	return c, nil
}

func (c *Container) LaunchContainer() {
	cwd, _ := os.Getwd()

	command := "docker run -it "
	if c.Name != "" {
		command += fmt.Sprintf("--name %s ", c.Name)
	}
	if c.User != "" {
		command += fmt.Sprintf("-u %s ", c.User)
	} else {
		command += "-u root "
	}
	for _, port := range c.Ports {
		command += fmt.Sprintf("-p %d:%d ", port, port)
	}
	for _, arg := range c.Args {
		command += arg + " "
	}
	command += fmt.Sprintf("-v %v:/tmp/dev -w /tmp/dev ", cwd)
	for _, mount := range c.Mounts {
		command += fmt.Sprintf("-v %s ", mount)
	}
	command += c.Image
	command += " "
	command += c.Shell

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func (c *Container) RemoveContainer() {
	command := fmt.Sprintf("docker rm %s", c.Name)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func (c *Container) ShellContainer() {
	command := "docker exec -it "
	if c.User != "" {
		command += fmt.Sprintf("-u %s ", c.User)
	} else {
		command += "-u root "
	}
	command += c.Name
	command += " "
	command += c.Shell

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func DumpContainerJSON() error {
	c := Container{Name: "foobar", Image: "busybox", User: "root", Shell: "sh", Mounts: make([]string, 0), Ports: make([]int, 0), Args: make([]string, 0)}
	file, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(".devctl.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
