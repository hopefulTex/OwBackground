package main

import (
	"fmt"
	"os"
	"strings"
)

func findGame() (string, error) {
	var blzPath string = "/Program Files (x86)/Overwatch/_retail_/Overwatch.exe"
	var stmPath string = "/ProgramFiles (x86)/Steam/steamapps/common/Overwatch/retail/overwatch.exe"
	var err error

	_, err = os.Stat(blzPath)
	if err != nil {
		fmt.Printf("getGame err: %s\n\n", err.Error())
		_, err = os.Stat(stmPath)
		if err != nil {
			return "", err
		} else {
			fmt.Printf("\nfound game file: %s\n", stmPath)
			return stmPath, nil
		}
	} else {
		fmt.Printf("\nfound game file: %s\n", blzPath)
		return blzPath, nil
	}
}

func getBg() (string, error) {
	var name string = "Heroes"
	var path string = "/.config/OwBackground/config.txt"
	var err error

	home, err := os.UserHomeDir()
	if err == nil {
		path = home + path
		fmt.Printf("\nConfig Path (?): %s\n", path)
		str, err := read(path)
		if err != nil {
			f, err := os.Create(path)
			if err == nil {
				f.WriteString(fmt.Sprintf("background: %s\n", name))
			}
			f.Close()
			return name, err
		}
		fmt.Printf("\nconfig text:\n%s\n\n", str)
		lines := strings.Split(str, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "background:") {
				name = strings.TrimPrefix(line, "background:")
				name = strings.TrimSpace(name)
				break
			}
		}
	}

	return name, err
}

func read(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	fmt.Printf("\nopened file: %s\n", path)

	return string(bytes), nil
}
