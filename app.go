package main

import (
	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/manager"
	"github.com/suynep/compilare/tests"
	"github.com/suynep/compilare/ui"
)

const IS_UI = true

func main() {
	if IS_UI {
		database.MustInitDB() // should be the first call. Always! :)
		ui.BuildUi()
		database.Close()
	} else {
		database.MustInitDB()             // should be the first call. Always! :)
		manager.CheckAndSaveLastRunTime() // runs everything (for the time being)
		tests.TestWebApiServer()
		database.Close()
	}
}
