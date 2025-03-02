package valigen

import (
	"fmt"
	"os"
)

func GenerateValidation(sourcePath, destinationPath, pkgName string) error {
	destFile, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("ファイル生成エラー: %w", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer destFile.Close()

	return nil
}
