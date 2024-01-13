package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/fs"
	"path/filepath"
)

func LoadEnv() error {
	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}

	return nil
}

func FindFiles(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}
