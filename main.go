package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Wails Todoapp",
		Width:     860,
		Height:    624,
		MinWidth:  860,
		MinHeight: 624,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		// Menu:             AppMenu,
		OnStartup:     app.startup,
		OnBeforeClose: app.beforeClose,
		Bind: []interface{}{
			app,
		},
		// Linux platform specific options
		Linux: &linux.Options{
			Icon: icon,
			// WindowIsTranslucent: true,
			WebviewGpuPolicy: linux.WebviewGpuPolicyNever,
			// ProgramName:         "wails",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

/*
COMMAND FOR LINUX BUILD:
wails build -clean -o todoapp

COMMAND FOR WINDOWS BUILD:
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc wails build -skipbindings -s -platform windows/amd64 -o todoapp.exe

https://madin.dev/cross-wails/

https://www.google.com/search?q=wails+compile+sqlite+for+windows&sca_esv=ac88267094842ae4&sxsrf=ADLYWIIB60WncbSg7MnPBYV7977EHLjeSQ%3A1729878267256&ei=-9gbZ--gD5qQxc8Pw87I2AE&ved=0ahUKEwiviv_ciqqJAxUaSPEDHUMnEhsQ4dUDCA8&uact=5&oq=wails+compile+sqlite+for+windows&gs_lp=Egxnd3Mtd2l6LXNlcnAiIHdhaWxzIGNvbXBpbGUgc3FsaXRlIGZvciB3aW5kb3dzMggQABiABBiiBDIIEAAYgAQYogQyCBAAGIAEGKIESNZbUL0gWL5DcAJ4AZABAJgBrwGgAc4IqgEDMS44uAEDyAEA-AEBmAIKoAKBCMICChAAGLADGNYEGEfCAgQQIxgnwgIIECEYoAEYwwSYAwCIBgGQBgiSBwMyLjigB7wV&sclient=gws-wiz-serp
*/

/* NATIVE MENU
// Menu settings
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("Delete all data", keys.CmdOrCtrl("d"), func(_ *menu.CallbackData) {
		opt := runtime.MessageDialogOptions{
			Type:          runtime.InfoDialog,
			Title:         "Data deletion",
			Message:       "This action is irreversible, export your data beforehand",
			DefaultButton: "OK",
			CancelButton:  "Cancel",
		}
		res, _ := runtime.MessageDialog(app.ctx, opt)
		if res != "OK" {
			return
		}
		if !app.dropTable() {
			opt := runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Title:   "Data deletion",
				Message: "Something went wrong while deleting the data",
			}
			runtime.MessageDialog(app.ctx, opt)
			return
		}
		runtime.WindowReload(app.ctx)
	})
	FileMenu.AddText("Export data", keys.CmdOrCtrl("e"), func(_ *menu.CallbackData) {
		if !app.exportData() {
			opt := runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Title:   "Exporting data",
				Message: "Something went wrong while exporting the data",
			}
			runtime.MessageDialog(app.ctx, opt)
			return
		}
		opt := runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Exporting data",
			Message: "The data was exported successfully",
		}
		runtime.MessageDialog(app.ctx, opt)
	})
	FileMenu.AddText("Import data", keys.CmdOrCtrl("i"), func(_ *menu.CallbackData) {
		opt := runtime.MessageDialogOptions{
			Type:          runtime.InfoDialog,
			Title:         "Importing data",
			Message:       "This action will overwrite the previous data",
			DefaultButton: "OK",
			CancelButton:  "Cancel",
		}
		res, _ := runtime.MessageDialog(app.ctx, opt)
		if res != "OK" {
			return
		}

		if !app.importData() {
			opt := runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Title:   "Importing data",
				Message: "The .csv file was not found or an error occurred",
			}
			runtime.MessageDialog(app.ctx, opt)
			return
		}

		runtime.WindowReload(app.ctx)
	})
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})
*/
