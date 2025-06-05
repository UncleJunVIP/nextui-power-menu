package main

import (
	gaba "github.com/UncleJunVIP/gabagool/pkg/gabagool"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	gaba.InitSDL(gaba.GabagoolOptions{
		WindowTitle:    "Save Sync",
		ShowBackground: true,
	})

	defer gaba.CloseSDL()

	mainMenuItems := []gaba.MenuItem{
		{
			Text:               "Reboot",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           "Reboot",
		},
		{
			Text:               "Shutdown",
			Selected:           false,
			Focused:            false,
			NotMultiSelectable: false,
			Metadata:           "Shutdown",
		},
	}

	options := gaba.DefaultListOptions("Power Menu", mainMenuItems)
	options.EnableAction = true
	options.FooterHelpItems = []gaba.FooterHelpItem{
		{ButtonName: "B", HelpText: "Quit"},
		{ButtonName: "A", HelpText: "Select"},
	}

	for {

		sel, err := gaba.List(options)
		if err != nil {
			log.Fatalf("Error displaying menu: %v", err)
		}

		if sel.IsNone() || sel.Unwrap().SelectedIndex == -1 {
			os.Exit(0)
		}

		switch sel.Unwrap().SelectedItem.Metadata.(string) {
		case "Reboot":
			gaba.ProcessMessage("Rebooting...", gaba.ProcessMessageOptions{ShowThemeBackground: true}, func() (interface{}, error) {
				cmd := exec.Command("reboot")
				cmd.Start()
				time.Sleep(5 * time.Second)
				return nil, nil
			})

		case "Shutdown":
			gaba.ProcessMessage("Shutting down...", gaba.ProcessMessageOptions{ShowThemeBackground: true}, func() (interface{}, error) {
				cmd := exec.Command("poweroff")
				cmd.Start()
				time.Sleep(5 * time.Second)
				return nil, nil
			})
		default:
			os.Exit(0)
		}
	}
}
