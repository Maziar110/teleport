package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {

	user_inputs := os.Args[1:]
	your_directory, err := os.Getwd()
	if err != nil {
		fmt.Println("An error appeared while trying to find yor location: ", err)
	}
	// If User enter nothing, it will print the current location
	if len(user_inputs) < 1 {

		fmt.Println("You are in ", your_directory)

	} else {
		// If user pass arguments to app
		handle_request(user_inputs, your_directory)

	}

}

func handle_request(arg []string, path string) {
	file_path := "./addressme.txt"
	if len(arg) > 1 {
		fmt.Println("Input Erro, enter alias as argument withough space")
		return
	} else {
		switch arg[0] {
		case `bulk`:
			fmt.Println("Beta Version, Not tested")
			bulk_commands := bulk_alias_from_file(file_path)
			update_bashrc(bulk_commands)

		case `cat`:
			cat()
		case `help`:
			help()
		default:
			var command []string

			command = append(command, alias_creator(arg[0], path))
			update_bashrc(command)
		}
	}

}

func alias_creator(alias string, path string) string {
	command := fmt.Sprintf("alias %s='cd %s'", alias, path)
	return command
}

// Reads a file and gets the alias from it and returns a Map of "alias": "location"
func bulk_alias_from_file(path string) []string {

	var address_book []string

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("An error while opening %s file", path)
	}
	defer file.Close()
	lines := bufio.NewScanner(file)
	for lines.Scan() {
		// I call this function (lines.text()) just once and use in 2 place to avoid calling it multipe times
		line_string := lines.Text()
		line := strings.Split(line_string, " ")
		if len(line) > 2 {
			fmt.Printf("There are more spaces in line %s which we picked %s as alias and %s as address", line_string, line[0], line[1])
		}
		address_book = append(address_book, alias_creator(line[0], line[1]))
	}
	if err := lines.Err(); err != nil {
		fmt.Println("An error while reading address lines: ", err)
	}
	return address_book
}

func alias_exists(addresses map[string]string, key string) (bool, string) {

	if address, ok := addresses[key]; ok {
		return true, address
	} else {
		return false, "empty"
	}

}

func update_bashrc(commands []string) bool {
	dir, is_directory := where_to_save()
	if is_directory {
		if file, err := os.OpenFile(dir, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644); err != nil {
			fmt.Printf("An error while adding Alias to %s , error is: %s", dir, err)
			return false
		} else {
			defer file.Close()
			for i := 0; i < len(commands); i += 1 {
				_, err := file.WriteString(commands[i] + "\n")
				if err != nil {
					fmt.Println(err)
					return false
				}

			}
			source_bashrc(dir)
			return true

		}

	} else {
		return false
	}
}

/*
Check OS and return a specific directory based on OS
*/
func where_to_save() (string, bool) {
	os_name := runtime.GOOS
	var bashrc_dir string
	switch os_name {
	case "windows":
		fmt.Println("Windows operating system is not supported yet")
		return "", false
	case "darwin":
		fmt.Println("MAC operating system is not supported yet")
		return "", false
	case "linux":
		home_dir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Home directory not found, setting default")
			home_dir = "~/"
		}
		bashrc_dir = home_dir + "/.bash_aliases"
		return bashrc_dir, true
	}
	return "", false
}

func source_bashrc(path string) {
	fmt.Printf("to apply the change, run: `source %s`  \n", path)
}

func cat() {
	path, _ := where_to_save()
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("An error while opening %s file", path)
	}
	defer file.Close()
	lines := bufio.NewScanner(file)
	for lines.Scan() {
		fmt.Println(lines.Text())
	}

}

func help() {
	help := "Supported arguments: \n cat (prints aliases) \n <text> (adds text as an alias for your current location) \n bulk (Beta - adds are locations added in a file with the name of addressme.txt"
	fmt.Println(help)
}
