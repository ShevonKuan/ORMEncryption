package utils

import (
	"os"

	"github.com/alecthomas/chroma/quick"
)

// PrintEncryptedData 打印加密后的数据
func PrintDecryptedData(decryptedData string) {
	quick.Highlight(os.Stdout, decryptedData, "xml", "terminal", "fruity")
}
