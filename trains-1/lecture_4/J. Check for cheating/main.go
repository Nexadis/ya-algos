package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() string {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	var caseSens, startDigit string
	fmt.Fscan(r, &caseSens, &startDigit)

	isCaseSens := isActive(caseSens)
	isStartDigit := isActive(startDigit)

	keyWords := make(map[string]struct{}, n)
	changeCase := caser(isCaseSens)

	for i := 0; i < n; i++ {
		var kw string
		fmt.Fscan(r, &kw)
		keyWords[changeCase(kw)] = struct{}{}
	}

	isVar := varChecker(keyWords, isCaseSens, isStartDigit)
	variables := map[string]int{}
	names := make([]string, 0, 10)

	for {
		var word string
		_, err := fmt.Fscan(r, &word)
		if err != nil {
			break
		}
		word = changeCase(word)
		vars := fetchVars(word)
		for _, v := range vars {
			if isVar(v) {
				if _, ok := variables[v]; !ok {
					names = append(names, v)
				}
				variables[v] += 1
			}
		}

	}

	mk := 0
	ans := ""
	for _, name := range names {
		v := variables[name]
		if v > mk {
			ans = name
			mk = v
		}
	}
	return ans
}

func isActive(value string) bool {
	return value == "yes"
}

func fetchVars(line string) []string {
	vars := make([]string, 0, 10)
	word := strings.Builder{}

	for i := 0; i < len(line); i++ {
		if unicode.IsLetter(rune(line[i])) || unicode.IsDigit(rune(line[i])) || line[i] == '_' {
			word.WriteByte(line[i])
			continue
		}
		if word.Len() > 0 {
			vars = append(vars, word.String())
			word.Reset()
		}
	}
	if word.Len() > 0 {
		vars = append(vars, word.String())
		word.Reset()
	}

	return vars
}

func varChecker(kwDict map[string]struct{}, isCaseSens, isStartDigit bool) func(w string) bool {
	return func(w string) bool {
		if !isStartDigit && unicode.IsDigit(rune(w[0])) {
			return false
		}
		if !isCaseSens {
			w = strings.ToLower(w)
		}
		if countDigits(w) == len(w) {
			return false
		}
		if _, ok := kwDict[w]; !ok {
			return true
		}
		return false
	}
}

func caser(isCaseSens bool) func(string) string {
	if isCaseSens {
		return func(s string) string { return s }
	}
	return func(s string) string {
		return strings.ToLower(s)
	}
}

func countDigits(w string) int {
	digits := 0
	for _, v := range w {
		if unicode.IsDigit(v) {
			digits++
		}
	}
	return digits
}
