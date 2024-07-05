package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func is_hidden(path string) bool {
	parts := strings.Split(path, string(os.PathSeparator))

	for _, part := range parts {
		if strings.HasPrefix(part, ".") {
			return true
		}
	}

	return false
}

func calc_indent_level(line, trimmed_line, indent_type string) int {
	if indent_type != "" {
		return (len(line) - len(trimmed_line)) / len(indent_type)
	}

	return 0
}

func find_indent_type(lines []string) string {
	for _, line := range lines {
		trimmed_line := strings.TrimLeftFunc(line, unicode.IsSpace)

		if len(trimmed_line) < len(line) {
			indent_type := line[0:(len(line) - len(trimmed_line))]

			return indent_type
		}
	}

	return ""
}

func find_max_indent_level(lines []string, indent_type string) int {
	max_indent_level := 0

	for _, line := range lines {
		trimmed_line := strings.TrimLeftFunc(line, unicode.IsSpace)
		indent_level := calc_indent_level(line, trimmed_line, indent_type)

		if indent_level > max_indent_level {
			max_indent_level = indent_level
		}
	}

	return max_indent_level
}

func do_invent(file_path string) error {
	contents, err := os.ReadFile(file_path)
	if err != nil {
		return err
	}

	if len(contents) == 0 {
		return errors.New("file is empty")
	}

	lines := strings.Split(string(contents), "\n")

	indent_type := find_indent_type(lines)
	max_indent_level := find_max_indent_level(lines, indent_type)

	var new_contents string
	for i, line := range lines {
		trimmed_line := strings.TrimLeftFunc(line, unicode.IsSpace)
		indent_level := calc_indent_level(line, trimmed_line, indent_type)

		inverse_indent_level := max_indent_level - indent_level
		inverse_indent := strings.Repeat(indent_type, inverse_indent_level)

		new_contents += inverse_indent + trimmed_line

		if i < len(lines)-1 {
			new_contents += "\n"
		}
	}

	err = os.WriteFile(file_path, []byte(new_contents), os.ModeAppend)
	if err != nil {
		return err
	}

	return nil
}

func print_usage() {
	fmt.Println(`usage:
  single file:
    invent --file <file>

  directory:
    invent --dir <dir>`)
	os.Exit(0)
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 || len(args) > 2 {
		print_usage()
	}

	switch args[0] {
	case "--file":
		file_path := args[1]

		err := do_invent(file_path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("* invented", file_path)

	case "--dir":
		dir_path := args[1]

		err := filepath.Walk(dir_path, func(file_path string, file_info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !is_hidden(file_path) &&
				!file_info.IsDir() &&
				file_info.Mode()&0111 == 0 {

				err := do_invent(file_path)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("* invented", file_path)

				return nil
			}

			return nil
		})
		if err != nil {
			log.Fatal(err)
		}

	default:
		print_usage()
	}
}
