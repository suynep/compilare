package manager

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/suynep/compilare/api"
	"github.com/suynep/compilare/types"
)

/*
	Manage run variables: Last Run Time, Last Fetched Posts, etc.
*/

const (
	CONFIG_PATH              = `config.json`
	STORY_SAVE_DELTA float64 = 24 * 60 // in minutes (for the time being; just for testing)
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

func GetConfig(path string) types.Config {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Error while reading Config JSON: %v", err)
	}

	config := new(types.Config)
	if err = json.Unmarshal(data, config); err != nil {
		log.Fatalf("Error while Unmarshalling JSON: %v", err)
	}

	return *config
}

func CheckAndSaveLastRunTime() {
	currentTime := time.Now()
	config := GetConfig(CONFIG_PATH)

	// if the difference is less than the minimum required delta for saving (e.g., say, 0.5 minutes),
	// update the config file with the new run time.
	if currentTime.Sub(config.Run.Time).Minutes() <= STORY_SAVE_DELTA {
		fmt.Printf("Current Delta %.1f does NOT exceed %.1f minutes cap\nWill only save the current run time...\n", currentTime.Sub(config.Run.Time).Minutes(), STORY_SAVE_DELTA)
	} else {
		fmt.Printf("Current Delta %.1f DOES exceed %.1f minutes cap\ninitiating database refresh...\n", currentTime.Sub(config.Run.Time).Minutes(), STORY_SAVE_DELTA)
		api.SaveTopStoriesDatabase()
		api.SaveBestStoriesDatabase()
		api.SaveNewStoriesDatabase()
		api.FullFlowRSS()
	}
	SaveLastRunTime()
}
