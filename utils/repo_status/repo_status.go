package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"
)

func wrapStringColorTag(input string, style int) string {
	return ""
}

// Вывод только список директорий (не файлов) из спика содержимого
// директории, полученного вызовом os.ReadDir(dir)
func printDirs(dirs []fs.DirEntry) {
	for _, dir := range dirs {
		if dir.IsDir() {
			println(dir.Name())
		}
	}
}

func filter(vs []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

var gitPorcelainCodes = map[string]string{
	" M": "updated in index",
	" T": "type changed in index",
	"A ": "added to index",
	" D": "deleted from index",
	" R": "renamed in index",
	" C": "copied in index",
	"??": "untracked",
	"!!": "ignored"}

func checkStatusPorcelain(input []string) string {
	var (
		prefix string
		rt     []string
	)

	// Оставляем только строки, содержащие более 3х символов.
	// Может ли git status --porcelain вернуть значущую строку
	// длинной менее 3х символов?
	filteredInput := filter(input, func(line string) bool {
		// Возвращает именно количество символов в строке, не её размер в байтах!
		return len([]rune(line)) >= 3
	})

	// Если после фильтра слайс пустой, считаем что нечего коммитить
	if len(filteredInput) == 0 {
		return "\033[;32mTherse nothing to commit!\033[0m"
	}

	for _, line := range filteredInput {
		// Берём префикс строки, содержащий код.
		prefix = line[0:2]
		if porcelainCode, ok := gitPorcelainCodes[prefix]; ok {
			rt = append(rt, porcelainCode+" - "+line[3:])
		} else {
			rt = append(rt, "Unknow code with file"+" "+line[3:])
		}
	}

	return strings.Join(rt[:], "\n")
}

func main() {
	workDir, err := os.Getwd()
	if err != nil {
		println(err)
		return
	}

	// Без контроля ошибки, потому что workDir всегда существует
	dirs, _ := os.ReadDir(workDir)

	for _, dir := range dirs {
		if dir.IsDir() {
			os.Chdir(workDir + "/" + dir.Name())

			gitStatus := exec.Command("git", "status", "--porcelain")
			out, err := gitStatus.Output()
			if err != nil {
				fmt.Printf("---=== %v ===---\n", dir.Name())
				fmt.Printf("Directory - %v has not contain a git repo!\n", dir.Name())
				println(err.Error())
				continue
			}

			parsed := checkStatusPorcelain(strings.Split(string(out), "\n"))

			fmt.Printf("---=== %v ===---\n", dir.Name())
			fmt.Printf("%v\n\n", parsed)
		}
	}
}
