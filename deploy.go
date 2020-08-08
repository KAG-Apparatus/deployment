package main

import (
	"errors"
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

type deployParameters struct {
	sourceDir          string
	destinationDir     string
	autoconfigPath     string
	gamemode           string
	rconPassword       string
	randomRconPassword bool
	serverName         string
	serverInfo         string
	kagPath            string
}

func deploy(params deployParameters) error {
	if params.sourceDir == "" || params.destinationDir == "" {
		return fmt.Errorf("you must specify source and destination")
	}

	err := copy.Copy(
		params.sourceDir,
		params.destinationDir,
		copy.Options{
			Skip: func(src string) (bool, error) {
				return strings.HasPrefix(src, ".") || strings.HasSuffix(src, ".git"), nil
			},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to copy: %v", err)
	}

	if params.autoconfigPath == "" {
		return nil
	}

	if params.gamemode == "" {
		return errors.New("gamemode not specified")
	}

	if params.kagPath == "" {
		return errors.New("kag executable not specified")
	}

	kagDir := filepath.Dir(params.kagPath)

	autoconfigFile, err := ioutil.ReadFile(params.autoconfigPath)
	if err != nil {
		log.Fatalf("failed to read autoconfig file: %v", err)
	}

	lines := strings.Split(string(autoconfigFile), "\n")
	for i, line := range lines {
		if strings.Contains(line, "sv_gamemode") {
			lines[i] = fmt.Sprintf("sv_gamemode = %s", params.gamemode)
			continue
		}
		if params.rconPassword != "" && strings.Contains(line, "sv_rconpassword") {
			lines[i] = fmt.Sprintf("sv_rconpassword = %s", params.rconPassword)
			continue
		}
		if params.randomRconPassword && strings.Contains(line, "sv_rconpassword") {
			rand.Seed(time.Now().UnixNano())
			password := generatePassword(6, 0, 2, 2)
			log.Printf("sv_rconpassword = %s", password)
			lines[i] = fmt.Sprintf("sv_rconpassword = %s", password)
			continue
		}
		if params.serverName != "" && strings.Contains(line, "sv_name") {
			lines[i] = fmt.Sprintf("sv_name = %s", params.serverName)
			continue
		}
		if params.serverInfo != "" && strings.Contains(line, "sv_info") {
			lines[i] = fmt.Sprintf("sv_info = %s", params.serverInfo)
			continue
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(params.autoconfigPath, []byte(output), 0644)
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
	lines = append(lines, params.gamemode)
	output = strings.Join(lines, "\n")
	err = ioutil.WriteFile(modsFilePath, []byte(output), 0644)
	if err != nil {
		log.Fatalf("failed to write mods file: %v", err)
	}

	cmd := exec.Cmd{
		Dir:    kagDir,
		Path:   params.kagPath,
		Stdout: os.Stdout,
	}
	err = cmd.Run()
	if err != nil {
		log.Printf("error running KAG executable: %s", err)
	}
	return nil
}
