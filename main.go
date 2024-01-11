// list courtesy of SkyBorik on Steam
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(os.Args)
	// read overwatch location
	path, err := findGame()
	if err != nil {
		fmt.Printf("\nerror finding game: %s\n", err.Error())
		return
	}
	// find config file
	bgString, err := getBg()
	if err != nil {
		fmt.Printf("\nerror finding config: %s\n", err.Error())
	}

	fmt.Printf("Path: %s\nBgString: %s\n", path, bgString)

	// // launch overwatch
	launchArg, err := composeLaunchCode(bgString)
	if err != nil {
		fmt.Printf("\nerror composing launch command: %s\n", err.Error())
		return
	}
	fmt.Printf("Launch Code: %s %s\n", path, launchArg)
	program := exec.Command(path, launchArg)
	program.Start()
}

// type background struct {
// 	name string `json:`
// 	hex  string
// 	kind string // "hero", "event", "esports", "holiday"
// }

func composeLaunchCode(bg string) (string, error) {

	bgs := map[string]string{}

	err := json.Unmarshal([]byte(jsonText), &bgs)
	if err != nil {
		return "", err
	}

	if bg == "random" || bg == "Random" {
		index := rand.Intn(len(bgs) - 1)
		i := 0
		for name := range bgs {
			if i == index {
				bg = name
				break
			}
			i++
		}
	}

	key, exists := bgs[bg]
	if !exists {
		return "", fmt.Errorf("cannot find background: %s", bg)
	}

	return "--lobbyMap=" + key, nil
}
