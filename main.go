package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// TODO: Change this into a config parameter.
	scripts_dir := "/workspaces/amici/"

	var chosen_script string
	fmt.Print("Choose script: ")
	fmt.Scan(&chosen_script)
	curr_script := scripts_dir + chosen_script

	var input_string string
	fmt.Print("Input string: ")
	fmt.Scan(&input_string)

	out, err := exec.Command(curr_script, input_string).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
