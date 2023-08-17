package split

import (
	"bufio"
	"fmt"
	"os"
)

func ByLine(filename string, sourceFile *os.File, linesPerFile int) {
	suffixLetters := GenerateSuffixLetters()
	fileCounter := 0
	lineCount := 0

	firstFilename := "x" + suffixLetters[fileCounter]
	newFile, err := os.Create(firstFilename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// ファイルから読み込み
	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()

		// 新しいファイルに行を書き込み
		_, err := newFile.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// 指定行数ごとに新しいファイルを作成
		if lineCount%linesPerFile == 0 {
			newFile.Close()
			fileCounter++
			if fileCounter >= len(suffixLetters) {
				fmt.Println("split: too many files")
				return
			}
			
			newFilename := "x" + suffixLetters[fileCounter]
			newFile, err = os.Create(newFilename)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			newFile, err = os.Create(newFilename)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer newFile.Close()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}