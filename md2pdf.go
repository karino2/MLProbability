package main

import (
	"bufio"
	"fmt"
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
		"-V", "documentclass=ltjreport",
		"-V", "classoption=a4j",
		"-V", "geometry:margin=1in",
		"--pdf-engine=lualatex"}

	pandocargs = append(pandocargs, fixed...)
	// fmt.Println(strings.Join(pandocargs, " "))

	// -o p_space.pdf -V documentclass=ltjarticle -V classoption=a4j -V geometry:margin=1in --pdf-engine=lualatex
	cmd := exec.Command("pandoc", pandocargs...)

	if outs, err := cmd.CombinedOutput(); err != nil {
		msg := fmt.Sprintf("%s: %s", outs, err)
		panic("cant execute pandoc " + strings.Join(pandocargs, " ") + " : " + msg)
	}

}
