package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func onDialogDelete(dialog *gtk.Dialog) bool {
	dialog.Hide()
	return true
}

func onFileChooserDialogDelete(fileChooserDialog *gtk.FileChooserDialog) bool {
	fileChooserDialog.Hide()
	return true
}

// onMainWindowDestory is the callback that is linked to the
// on_main_window_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.
func onMainWindowDestroy() {
	log.Println("Leaving application. Good bye!")
}

func onFolderChooser(s selection) {
	obj, err := builder.GetObject("folder_chooser")
	errorCheck(err)
	chooser, err := isFileChooserDialog(obj)
	errorCheck(err)

	switch s {
	case source:
		chooser.SetTitle("Select development folder")
		if params.sourceDir != "" {
			parent := filepath.Dir(params.sourceDir)
			chooser.SetCurrentFolder(parent)
		}
		response := chooser.Run()
		if response == gtk.RESPONSE_CANCEL {
			chooser.Hide()
			return
		}
		folder := strings.TrimPrefix(chooser.GetURI(), "file://")
		params.sourceDir = folder
		obj, err = builder.GetObject("entry_src")
		errorCheck(err)
		entrySource, err := isEntry(obj)
		errorCheck(err)
		entrySource.SetText(folder)
		chooser.Hide()

	case destination:
		chooser.SetTitle("Select your KAG server Mods folder")
		if params.destinationDir != "" {
			parent := filepath.Dir(params.destinationDir)
			chooser.SetCurrentFolder(parent)
		}
		response := chooser.Run()
		if response == gtk.RESPONSE_CANCEL {
			chooser.Hide()
			return
		}
		folder := strings.TrimPrefix(chooser.GetURI(), "file://")
		params.destinationDir = folder
		obj, err = builder.GetObject("entry_dst")
		errorCheck(err)
		entryDestination, err := isEntry(obj)
		errorCheck(err)
		entryDestination.SetText(folder)

		gamemode := filepath.Base(params.destinationDir)
		params.gamemode = gamemode
		obj, err = builder.GetObject("entry_gamemode")
		errorCheck(err)
		entryGamemode, err := isEntry(obj)
		errorCheck(err)
		entryGamemode.SetText(gamemode)

		chooser.Hide()

	}
}

func onFileChooser(s selection) {
	obj, err := builder.GetObject("file_chooser")
	errorCheck(err)
	chooser, err := isFileChooserDialog(obj)
	errorCheck(err)

	switch s {
	case autoconfig:
		chooser.SetTitle("Select server's autoconfig.cfg file")
		filter, err := gtk.FileFilterNew()
		errorCheck(err)
		filter.AddPattern("autoconfig.cfg")
		chooser.SetFilter(filter)
		if params.autoconfigPath != "" {
			parent := filepath.Dir(params.autoconfigPath)
			chooser.SetCurrentFolder(parent)
		}
		response := chooser.Run()
		if response == gtk.RESPONSE_CANCEL {
			chooser.Hide()
			return
		}
		file := strings.TrimPrefix(chooser.GetURI(), "file://")
		params.autoconfigPath = file
		obj, err = builder.GetObject("entry_autoconfig")
		errorCheck(err)
		entryAutoconfig, err := isEntry(obj)
		errorCheck(err)
		entryAutoconfig.SetText(file)
		chooser.Hide()

	case kag:
		chooser.SetTitle("Select KAG server executable")
		filter, err := gtk.FileFilterNew()
		errorCheck(err)
		filter.AddPattern("KAG*")
		filter.AddMimeType("application/octet-stream")
		chooser.SetFilter(filter)
		if params.kagPath != "" {
			parent := filepath.Dir(params.kagPath)
			chooser.SetCurrentFolder(parent)
		}
		response := chooser.Run()
		if response == gtk.RESPONSE_CANCEL {
			chooser.Hide()
			return
		}
		file := strings.TrimPrefix(chooser.GetURI(), "file://")
		params.kagPath = file
		obj, err = builder.GetObject("entry_kag")
		errorCheck(err)
		entryKAG, err := isEntry(obj)
		errorCheck(err)
		entryKAG.SetText(file)
		chooser.Hide()

	case open:
		chooser.SetTitle("Select TOML file to save")
		filter, err := gtk.FileFilterNew()
		errorCheck(err)
		filter.AddPattern("*.toml")
		chooser.SetFilter(filter)

		response := chooser.Run()
		if response == gtk.RESPONSE_CANCEL {
			chooser.Hide()
			return
		}

		file := strings.TrimPrefix(chooser.GetURI(), "file://")
		log.Println(file)

		//Do your magic here

		chooser.Hide()
	}
}

func onButtonSource(button *gtk.Button) {
	onFolderChooser(source)
}

func onButtonDestination(button *gtk.Button) {
	onFolderChooser(destination)
}

func onButtonAutoconfig(button *gtk.Button) {
	onFileChooser(autoconfig)
}

func onButtonKAG(button *gtk.Button) {
	onFileChooser(kag)
}

func onSwitchRCON(gtkSwitch *gtk.Switch, isOn bool) {
	obj, err := builder.GetObject("entry_rcon")
	errorCheck(err)
	entryRCON, err := isEntry(obj)
	errorCheck(err)
	if !isOn {
		entryRCON.SetCanFocus(true)
		entryRCON.SetEditable(true)
		return
	}
	password := generatePassword(6, 0, 2, 2)
	params.rconPassword = password
	entryRCON.SetCanFocus(false)
	entryRCON.SetEditable(false)
	entryRCON.SetText(password)
}

func onEntryName(entry *gtk.Entry) {
	text, err := entry.GetText()
	errorCheck(err)
	params.serverName = text
}

func onEntryInfo(entry *gtk.Entry) {
	text, err := entry.GetText()
	errorCheck(err)
	params.serverInfo = text
}

func onButtonDeploy(button *gtk.Button) {
	log.Printf("Deplying server with parameters %v\n", params)

	obj, err := builder.GetObject("textbuffer_output")
	errorCheck(err)
	textBuffer, err := isTextBuffer(obj)
	errorCheck(err)

	cmd, err := deploy(params)
	errorCheck(err)
	out, err := cmd.StdoutPipe()
	errorCheck(err)

	err = cmd.Start()
	errorCheck(err)

	obj, err = builder.GetObject("label_output")
	errorCheck(err)
	labelOutput, err := isLabel(obj)
	errorCheck(err)

	labelOutput.SetText(fmt.Sprintf("Server started with PID %d. Output:", cmd.Process.Pid))

	buffer := make([]byte, 1024)

	go func() {
		defer cmd.Process.Kill()
		iterator := textBuffer.GetEndIter()
		for {
			n, _ := out.Read(buffer)
			glib.IdleAdd(func() {
				textBuffer.Insert(iterator, string(buffer[:n]))
			})
		}
	}()
}

func onTextViewOutputSizeAllocate(textView *gtk.TextView, allocation uintptr) {
	buffer, err := textView.GetBuffer()
	errorCheck(err)
	mark := buffer.CreateMark("end_mark", buffer.GetEndIter(), true)
	textView.ScrollToMark(mark, 0, false, 0, 1)
}

func onMenuItemOpen(menu *gtk.MenuItem) {
	onFileChooser(open)
}
