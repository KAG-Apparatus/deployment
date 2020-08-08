package main

import "flag"

func main() {
	sourceDir := flag.String("src", ".", "-src [Source Directory]")
	destinationDir := flag.String("dst", "./Mods", "-dst [Destination Directory]")
	autoconfigFile := flag.String("autoconfig", "./autoconfig.cfg", "-autoconfig [Autoconfig File]")
	gamemode := flag.String("gamemode", "", "-gamemode [Selected Gamemode]")
	rconPassword := flag.String("rconpassword", "", "-rconpassword [rcon Administrative Password]")
	serverName := flag.String("name", "", "-name [Server Name]")
	serverInfo := flag.String("info", "", "-info [Server Info]")
	flag.Parse()
}
