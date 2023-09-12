package split

import (
	"fmt"
	"io"
	"os"
)

func ByByte(filename string, sourceFile *os.File, bytesPerFile int64) error {
	suffixLetters := GenerateSuffixLetters()
	fileCounter := 0

	fileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	numFiles := fileSize / bytesPerFile + 1
	bytesWritten := int64(0)
	buffer := make([]byte, 8)
	
	for i := int64(1); i <= numFiles; i++ {
		firstFilename := "x" + suffixLetters[fileCounter]
		newFile, err := os.Create(firstFilename)
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