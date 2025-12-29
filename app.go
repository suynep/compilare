package main

import (
	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/manager"
	"github.com/suynep/compilare/tests"
)

func main() {
	database.MustInitDB()             // should be the first call. Always! :)
	manager.CheckAndSaveLastRunTime() // runs everything (for the time being)
	tests.TestWebApiServer()
	database.Close()
}
