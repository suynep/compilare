package main

import (
	"github.com/suynep/compilare/database"
	_ "github.com/suynep/compilare/manager"
	"github.com/suynep/compilare/tests"
	"github.com/suynep/compilare/ui"
)

const IS_UI = true

func main() {
	database.MustInitDB() // should be the first call. Always! :)
	if IS_UI {
		ui.BuildUi()
	} else {
		// manager.CheckAndSaveLastRunTime() // runs everything (for the time being)
		tests.TestWebApiServer()
		// tests.TestReadForMemoization()
	}
	database.Close()
}
