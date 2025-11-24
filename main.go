package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	//TODO: Change this into a config parameter.
	scripts_dir := "#"

	if scripts_dir[len(scripts_dir)-1:] != "/" {
		scripts_dir = scripts_dir + "/"
	}

	var input_scripts, input_string string
	fmt.Print("Choose scripts, separated by a comma (,): ")
	fmt.Scan(&input_scripts)

	scripts_arr := strings.Split(input_scripts, ",")

	fmt.Print("Input string: ")
	fmt.Scan(&input_string)

	for _, v := range scripts_arr {
		out, err := exec.Command(scripts_dir + v, input_string).Output()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}
}
