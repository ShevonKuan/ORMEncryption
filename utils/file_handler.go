package utils

import "os"

// readFile 读取文件内容并返回字符串
func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// writeFile 将字符串写入文件
func WriteFile(input string, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(input)
	if err != nil {
		return err
	}

	return nil
}
