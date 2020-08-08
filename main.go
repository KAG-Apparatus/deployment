package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/otiai10/copy"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
	sourceDir := flag.String("src", "", "-src [Source Directory]")
	destinationDir := flag.String("dst", "", "-dst [Destination Directory]")
	autoconfigPath := flag.String("autoconfig", "", "-autoconfig [Autoconfig File]")
	gamemode := flag.String("gamemode", "", "-gamemode [Selected Gamemode]")
	rconPassword := flag.String("rconpassword", "", "-rconpassword [rcon Administrative Password]")
	randomRconPassword := flag.Bool("randomrcon", false, "-randomrcon")
	serverName := flag.String("name", "", "-name [Server Name]")
	serverInfo := flag.String("info", "", "-info [Server Info]")
	kagPath := flag.String("kag", "", "-kag [Kag Path]")
	flag.Parse()

	err := copy.Copy(
		*sourceDir,
		*destinationDir,
		copy.Options{
			Skip: func(src string) (bool, error) {
				return strings.HasPrefix(src, ".") || strings.HasSuffix(src, ".git"), nil
			},
		},
	)
	if err != nil {
		log.Fatalf("failed to copy: %v", err)
	}

	if *autoconfigPath == "" {
		return
	}

	if *gamemode == "" {
		log.Fatalln(errors.New("gamemode not specified"))
	}

	if *kagPath == "" {
		log.Fatalln(errors.New("kag executable not specified"))
	}

	kagDir := filepath.Dir(*kagPath)

	autoconfigFile, err := ioutil.ReadFile(*autoconfigPath)
	if err != nil {
		log.Fatalf("failed to read autoconfig file: %v", err)
	}

	lines := strings.Split(string(autoconfigFile), "\n")
	for i, line := range lines {
		if strings.Contains(line, "sv_gamemode") {
			lines[i] = fmt.Sprintf("sv_gamemode = %s", *gamemode)
			continue
		}
		if *rconPassword != "" && strings.Contains(line, "sv_rconpassword") {
			lines[i] = fmt.Sprintf("sv_rconpassword = %s", *rconPassword)
			continue
		}
		if *randomRconPassword && strings.Contains(line, "sv_rconpassword") {
			rand.Seed(time.Now().UnixNano())
			password := generatePassword(6, 0, 2, 2)
			log.Printf("sv_rconpassword = %s", password)
			lines[i] = fmt.Sprintf("sv_rconpassword = %s", password)
			continue
		}
		if *serverName != "" && strings.Contains(line, "sv_name") {
			lines[i] = fmt.Sprintf("sv_name = %s", *serverName)
			continue
		}
		if *serverInfo != "" && strings.Contains(line, "sv_info") {
			lines[i] = fmt.Sprintf("sv_info = %s", *serverInfo)
			continue
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(*autoconfigPath, []byte(output), 0644)
	if err != nil {
		log.Fatalf("failed to write autoconfig file: %v", err)
	}

	modsFilePath := fmt.Sprintf("%s%c%s", kagDir, os.PathSeparator, "mods.cfg")
	modsFile, err := ioutil.ReadFile(modsFilePath)
	if err != nil {
		log.Fatalf("failed to read mods file: %v", err)
	}

	lines = strings.Split(string(modsFile), "\n")
	for i, line := range lines {
		if !strings.HasPrefix(line, "#") {
			lines[i] = fmt.Sprintf("# %s", line)
			continue
		}
	}
	lines = append(lines, *gamemode)
	output = strings.Join(lines, "\n")
	err = ioutil.WriteFile(modsFilePath, []byte(output), 0644)
	if err != nil {
		log.Fatalf("failed to write mods.cfg file: %v", err)
	}

	cmd := exec.Cmd{
		Dir:    kagDir,
		Path:   *kagPath,
		Stdout: os.Stdout,
	}
	err = cmd.Run()
	if err != nil {
		log.Printf("error running KAG executable: %s", err)
	}
}

func errorCheck(err error) {
	log.Fatalln(err)
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
