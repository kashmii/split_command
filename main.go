package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"split/split"
)

func main() {
	bFlag := flag.Int64("b", -1, "An integer value for option -b")
	lFlag := flag.Int("l", -1, "An integer value for option -l")
	nFlag := flag.Int("n", -1, "An integer value for option -n")
	flag.Parse()

	args := flag.Args()
	filename := args[0]

	linesPerFile := 1000

	validation(filename, args)

	sourceFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sourceFile.Close()

	// オプションに応じて分割
	if *bFlag != -1 {
		err := split.ByByte(filename, sourceFile, *bFlag)
		if err != nil {
			fmt.Errorf("An error occurred with ByByte method: %v", err)
		}
	} else if *lFlag != -1 {
		linesPerFile = *lFlag
		err := split.ByLine(filename, sourceFile, linesPerFile)
		if err != nil {
			fmt.Errorf("An error occurred with ByLine method: %v", err)
		}
	} else if *nFlag != -1 {
		err := split.ByNumber(filename, sourceFile, *nFlag)
		if err != nil {
			fmt.Errorf("An error occurred with ByNumber method: %v", err)
		}
	} else {
		err := split.ByLine(filename, sourceFile, linesPerFile)
		if err != nil {
			fmt.Errorf("An error occurred with ByLine method: %v", err)
		}
	}
}

func validation(filename string, args []string) {
	fileInfo, err := os.Stat(filename)

	// コマンドライン引数が足りない場合
	if len(args) < 1 || filename == "-" {
		log.Fatalln("Arrgument is not enough.")
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
}