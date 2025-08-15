package main

import "strings"

func main() {
	println(DNAStrand("ACTGCTGA"))
}

func DNAStrand(dna string) string {
	var str = []rune{}
	for _, char := range dna {
		switch char {
		case 'A':
			str = append(str, 'T')
		case 'T':
			str = append(str, 'A')
		case 'C':
			str = append(str, 'G')
		case 'G':
			str = append(str, 'C')
		}
	}
	return string(str)
}

func DNAStrandReplacer(dna string) string {
	var replacer *strings.Replacer = strings.NewReplacer(
		"A", "T",
		"T", "A",
		"C", "G",
		"G", "C",
	)
	return replacer.Replace(dna)
}
