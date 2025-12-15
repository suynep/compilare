package main

import (
	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/tests"
)

func main() {
	database.MustInitDB() // should be the first call. Always! :)
	tests.TestDatabaseSaves()
}
