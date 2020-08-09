package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appID = "com.github.KAG-Apparatus.deployment"

type selection int

const (
	source selection = iota
	destination
	autoconfig
	kag

	open
	save
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet

	application *gtk.Application
	builder     *gtk.Builder
	mainWindow  *gtk.ApplicationWindow
	params      deployParameters
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
		cmd, err := deploy(deployParameters{
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
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			log.Printf("error running KAG executable: %s", err)
		}
		return
	}

	// Create a new application.
	var err error
	application, err = gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		log.Println("application activate")

		// Get the GtkBuilder UI definition in the glade file.
		builder, err = gtk.BuilderNewFromFile("resources/gui/main.glade")
		errorCheck(err)

		// Get the object with the id of "main_window".
		obj, err := builder.GetObject("main_window")
		errorCheck(err)
		// Verify that the object is a pointer to a gtk.ApplicationWindow.
		mainWindow, err = isWindow(obj)
		errorCheck(err)

		// Map the handlers to callback functions, and connect the signals
		// to the Builder.
		signals := map[string]interface{}{
			"on_main_window_destroy":            onMainWindowDestroy,
			"on_menu_item_open":                 onMenuItemOpen,
			"on_button_source":                  onButtonSource,
			"on_button_destination":             onButtonDestination,
			"on_button_autoconfig":              onButtonAutoconfig,
			"on_button_kag":                     onButtonKAG,
			"on_switch_rcon":                    onSwitchRCON,
			"on_entry_name":                     onEntryName,
			"on_entry_info":                     onEntryInfo,
			"on_button_deploy":                  onButtonDeploy,
			"on_file_chooser_dialog_delete":     onFileChooserDialogDelete,
			"on_text_view_output_size_allocate": onTextViewOutputSizeAllocate,
		}
		builder.ConnectSignals(signals)

		// Show the Window and all of its components.
		//mainWindow.Maximize()
		mainWindow.Show()
		application.AddWindow(mainWindow)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("Application finished.")
	})

	// Launch the application
	os.Exit(application.Run(os.Args))
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	rand.Seed(time.Now().UnixNano())

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
