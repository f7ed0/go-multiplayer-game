package gamemap

import (
	"encoding/json"
	"io"
	"os"

	"github.com/f7ed0/golog/lg"
)

func LoadMap(path string) (gm HeadLessMap, err error) {
	gm.Path = path
	// LOADING HITBOXES
	f, err := os.Open(path + "/hitboxes.json")
	if err != nil {
		return
	}
	res, err := io.ReadAll(f)
	if err != nil {
		return
	}
	f.Close()
	err = json.Unmarshal(res, &gm)
	if err != nil {
		return
	}
	lg.Debug.Println(gm)
	return
}
