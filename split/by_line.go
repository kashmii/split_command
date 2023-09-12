package split

import (
	"bufio"
	"fmt"
	"os"
)

func ByLine(filename string, sourceFile *os.File, linesPerFile int) error {
	suffixLetters := GenerateSuffixLetters()
	fileCounter := 0
	lineCount := 0

	// 修正すべき？： 1つ目のファイルをループの外で作成している
	firstFilename := "x" + suffixLetters[fileCounter]
	newFile, err := os.Create(firstFilename)
	if err != nil {
		fmt.Println("Error:", err)
		return err
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
			return err
		}

		// 指定行数ごとに新しいファイルを作成
		if lineCount%linesPerFile == 0 {
			newFile.Close()
			fileCounter++
			if fileCounter >= len(suffixLetters) {
				fmt.Println("split: too many files")
				return nil
			}
			
			// 修正すべき？： newFileという変数が2箇所で使われている
			newFilename := "x" + suffixLetters[fileCounter]
			newFile, err = os.Create(newFilename)
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}

			newFile, err = os.Create(newFilename)
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
			defer newFile.Close()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
	return nil
}