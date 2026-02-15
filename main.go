package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"strings"
	"bufio"
)

func main() {
	//TODO: Change this into a config parameter.
	scripts_dir := "."

	if scripts_dir[len(scripts_dir)-1:] != "/" {
		scripts_dir = scripts_dir + "/"
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose scripts, separated by a comma (,): ")
	input_scripts, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read input: ", err)
	}

	fmt.Print("Input string: ")
	input_string, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read input: ", err)
	}
	input_string = strings.TrimSpace(input_string)

	for v := range strings.SplitSeq(input_scripts, ",") {
		out, err := exec.Command(scripts_dir + strings.TrimSpace(v), input_string).Output()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}
}
