package fsys

import (
	"bx/cmd/vars"
	"os"

	"github.com/nultero/tics"
	"github.com/spf13/viper"
)

type cache map[string]interface{}

func ReadCache() (cache, error) {
	dir := viper.GetString(vars.DataDir)
	path := dir + "/cache.txt"

	cache := cache{}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return cache, err
	}

	cache = tics.ToJson(bytes)
	return cache, nil
}
