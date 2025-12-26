package main

import (
	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/manager"
)

func main() {
	database.MustInitDB() // should be the first call. Always! :)
	// tests.TestDatabaseSaves()
	// tests.TestSaveRunTime()
	// tests.TestReadForMemoization()
	manager.CheckAndSaveLastRunTime()

	database.Close()
}
