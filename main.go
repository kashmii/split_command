package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"split/split"
)

type Options struct {
	linesPerFile int
}

func main() {
	var options Options
	options.linesPerFile = 1000

	bFlag := flag.Int64("b", -1, "An integer value for option -b")
	lFlag := flag.Int("l", -1, "An integer value for option -l")
	nFlag := flag.Int("n", -1, "An integer value for option -n")
	flag.Parse()

	args := flag.Args()
	filename := args[0]
	fileInfo, err := os.Stat(filename)

	// コマンドライン引数が足りない場合
	if len(args) < 1 || filename == "-" {
		fmt.Println("本来の挙動: 標準入力から読み込んだものを新ファイルに書き込む")
		return
	}

	// ファイルが存在しない場合
	if os.IsNotExist(err) {
		log.Fatalf("File '%s' does not exist.\n", filename)
	} else if err != nil {
		log.Fatal("Error:", err)
	}

	// ディレクトリ名が入力された場合
	if fileInfo.IsDir() {
		log.Fatalf("'%s' is a directory.\n", filename)
	} 

	// 読み書き可能かチェック
	if fileInfo.Mode().Perm()&0400 == 0 {
		log.Fatalf("File '%s' is not readable.\n", filename)
	} 
	if fileInfo.Mode().Perm()&0200 == 0 {
		log.Fatalf("File '%s' is not writable.\n", filename)
	}

	// ==================================================
	sourceFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sourceFile.Close()

	// オプションに応じて分割
	if *bFlag != -1 {
		split.ByByte(filename, sourceFile, *bFlag)
	} else if *lFlag != -1 {
		options.linesPerFile = *lFlag
		split.ByLine(filename, sourceFile, options.linesPerFile)
	} else if *nFlag != -1 {
		split.ByNumber(filename, sourceFile, *nFlag)
	} else {
		split.ByLine(filename, sourceFile, options.linesPerFile)
	}
}