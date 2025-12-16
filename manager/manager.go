package manager

import (
	"encoding/json"
	"os"
	"time"

	"github.com/suynep/compilare/types"
)

/*
	Manage run variables: Last Run Time, Last Fetched Posts, etc.
*/

const (
	CONFIG_PATH      = `config.json`
	STORY_SAVE_DELTA = 0.5 // in minutes (for the time being; just for testing)
)

var (
	CONFIG = new(types.Config)
)

func SaveLastRunTime() {
	CONFIG.Run.Time = time.Now()

	data, err := json.Marshal(*CONFIG)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(CONFIG_PATH, data, 0644)

	if err != nil {
		panic(err)
	}
}
