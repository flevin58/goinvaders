package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

func FilterSlice[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func GetConfigPath(filename string) (string, error) {
	cfgDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find the home directory")
	}
	cfgDir = filepath.Join(cfgDir, ".config", "goinvaders")
	err = os.MkdirAll(cfgDir, 0775)
	if err != nil {
		return "", fmt.Errorf("could not create the config folder %s", cfgDir)
	}
	return filepath.Join(cfgDir, filename), nil
}
