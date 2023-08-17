package split

func GenerateSuffixLetters() []string {
	var suffixLetters []string

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	// 2文字の組み合わせを生成（676通り）
	for _, c1 := range alphabet {
		for _, c2 := range alphabet {
			suffixLetters = append(suffixLetters, string(c1)+string(c2))
		}
	}

	return suffixLetters
}