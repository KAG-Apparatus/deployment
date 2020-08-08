package main

import (
	"flag"
	"log"
	"math/rand"
	"strings"
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

	if flag.NFlag() > 0 {
		err := deploy(deployParameters{
			sourceDir:          *sourceDir,
			destinationDir:     *destinationDir,
			autoconfigPath:     *autoconfigPath,
			gamemode:           *gamemode,
			rconPassword:       *rconPassword,
			randomRconPassword: *randomRconPassword,
			serverName:         *serverName,
			serverInfo:         *serverInfo,
			kagPath:            *kagPath,
		})
		if err != nil {
			log.Fatalf("error on deploying: %v", err)
		}
		return
	}
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
