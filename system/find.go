package system

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/warrengalyen/fontivate/utils"
)

func FindFiles(root string) []string {
	var found []string
	err := filepath.WalkDir(root, func(f string, item fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Unsupported mime type: %w", err)
		}
		if _, ok := utils.SystemFontTypes[filepath.Ext(item.Name())]; ok {
			found = append(found, f)
			fmt.Println(found)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("WalkDir process error")
	}
	return found
}