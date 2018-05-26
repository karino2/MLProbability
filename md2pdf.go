package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func main() {

	file, err := os.Open("contents.txt")
	if err != nil {
		panic("can't open content.txt")
	}

	defer file.Close()

	pandocargs := make([]string, 0, 10)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		pandocargs = append(pandocargs, scanner.Text())
	}

	fixed := []string{"-o", "MLProbability.pdf",
//		"-V", "documentclass=ltjarticle",
		"-V", "documentclass=report",
		"-V", "classoption=a4j",
		"-V", "geometry:margin=1in",
		"--pdf-engine=lualatex"}

	pandocargs = append(pandocargs, fixed...)
	// fmt.Println(strings.Join(pandocargs, " "))

	// -o p_space.pdf -V documentclass=ltjarticle -V classoption=a4j -V geometry:margin=1in --pdf-engine=lualatex
	cmd := exec.Command("pandoc", pandocargs...)

	if err = cmd.Run(); err != nil {
		panic("cant execute pandoc " + strings.Join(pandocargs, " "))
	}

}
