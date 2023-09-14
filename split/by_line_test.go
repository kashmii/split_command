package split

import (
	"bufio"
	"log"
	"os"
	"testing"
)

// 指定した行数ごとにファイルが分割されるかをテスト

// 1.規定行数ごとにファイルが分割されるか

func TestByLine(t *testing.T) {
	// ・分割するファイル自体を作成
	filename := "dummy.txt"
	lineCount := 0
	sourceFile, _ := os.CreateTemp("", filename)
	defer sourceFile.Close()
	sourceFile.WriteString("Line 1\nLine 2\nLine 3\nLine 4\nLine 5\n")

	ExpectContent := []string {"Line 1\n", "Line 2\n", "Line 3\n", "Line 4\n", "Line 5\n"}

	// ・該当するメソッドを使って分割
	ByLine(filename, sourceFile, 2)
	// ・分割されたファイルの数を確認
	expected := "Line 1\nLine 2\n"
	// テスト用のファイルを読み込み
	// data, err := os.ReadFile("xaa")
	file, err := os.Open("xaa")
	if err != nil {
		log.Fatalf("ファイルを開けませんでした: %v", err)
	}
	defer file.Close()

	// ファイルからテキストを行単位で読み込む
	scanner := bufio.NewScanner(file)
	// ファイルの内容と期待値を比較
	for scanner.Scan() {
		line := scanner.Text()
		if line != ExpectContent[lineCount] {
			t.Errorf("File content does not match. Got: %s, Expected: %s", line, expected)
		}
		lineCount++
	}

	// ・分割されたファイルの行数を確認

	os.Remove(sourceFile.Name())
}
