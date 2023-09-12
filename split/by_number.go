package split

import (
	"fmt"
	"io"
	"os"
)

func ByNumber(filename string, sourceFile *os.File, numFile int) error {
	suffixLetters := GenerateSuffixLetters()
	fileCounter := 0

	fileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	bytesPerFile := fileSize / int64(numFile)
	bytesWritten := int64(0)
	buffer := make([]byte, 8)

	for i := 1; i <= numFile; i++ {
		newFilename := "x" + suffixLetters[fileCounter]
		newFile, err := os.Create(newFilename)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		defer newFile.Close()

		for bytesWritten < bytesPerFile {
			n, err := sourceFile.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			_, err = newFile.Write(buffer[:n])
			if err != nil {
				return err
			}

			bytesWritten += int64(n)
		}
		fileCounter++
		if fileCounter >= len(suffixLetters) {
			fmt.Println("split: too many files")
			return nil
		}
		bytesWritten = int64(0)
	}
	return nil
}