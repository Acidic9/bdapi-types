// Scraper is a go script to take the markdown documentation and output it to TypeScript documentation.

package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type Argument struct {
	Name        string
	ArgType     string
	Description string
}

type ReturnType struct {
	ReturnType  string
	Description string
}

type Method struct {
	Name        string
	Args        []Argument
	Description string
	RetType     ReturnType
}

func main() {
	docs, err := ioutil.ReadFile("./docs.md")
	if err != nil {
		log.Fatal(err)
	}

	parts := bytes.Split(docs, []byte("***"))
	methods := make([]Method, 0, len(parts))
	for _, part := range parts {
		part = bytes.TrimSpace(part)

		var method Method
		method.Name = getMethodName(part)
		method.Args = getMethodArgs(part)
		method.Description = getMethodDescription(part)
		method.RetType = getMethodReturnType(part)

		methods = append(methods, method)
	}

	tmpl, err := template.New("documentation").Funcs(template.FuncMap{
		"plus": func(a int, b int) int {
			return a + b
		},
		"def": func(str string, def string) string {
			if str == "" {
				return def
			}
			return str
		},
		"splitLines": func(str string) []string {
			return strings.Split(str, "\n")
		},
		"funcThenVoid": func(t string) string {
			if t == "function" {
				return "() => void"
			}
			return t
		},
		"trimSpace": strings.TrimSpace,
	}).ParseFiles("./documentation.text")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("types.d.ts")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.ExecuteTemplate(f, "documentation", methods)
	if err != nil {
		log.Fatal(err)
	}
}

func getMethodName(part []byte) string {
	exp := regexp.MustCompile("####\\s?\\`(\\w+)")
	matches := exp.FindSubmatch(part)
	if len(matches) < 2 {
		return ""
	}
	return string(matches[1])
}

func getMethodArgs(part []byte) (args []Argument) {
	// exp := regexp.MustCompile("####\\s?\\`\\w+\\(((?:\\w+,?\\s*)*)")
	// matches := exp.FindSubmatch(part)
	// if len(matches) < 2 {
	// 	return
	// }
	// argStr := bytes.Replace(matches[1], []byte{' '}, []byte{0x0}, -1)
	// argKeys := strings.Split(string(argStr), ",")

	exp := regexp.MustCompile("\\|`(\\w+)\\`\\|(\\w+)\\|(.*)\\|")
	matches2 := exp.FindAllSubmatch(part, -1)
	for _, match := range matches2 {
		if len(match) < 2 {
			continue
		}
		var arg Argument
		if len(match) >= 2 {
			arg.Name = string(match[1])
		}
		if len(match) >= 3 {
			arg.ArgType = string(match[2])
		}
		if len(match) >= 4 {
			arg.Description = string(match[3])
		}

		args = append(args, arg)
	}
	return
}

func getMethodDescription(part []byte) string {
	partStr := string(bytes.TrimSpace(part))
	lines := strings.Split(partStr, "\n")
	beginTableExp := regexp.MustCompile("^\\|Parameter")
	var lineStartIndex uint8
	var lineEndIndex uint8 = uint8(len(lines))
	for lineNum, line := range lines {
		if strings.HasPrefix(line, "####") {
			lineStartIndex = uint8(lineNum) + 1
		}
		if beginTableExp.MatchString(line) || strings.HasPrefix(line, "**@Returns*") {
			lineEndIndex = uint8(lineNum) - 1
			break
		}
	}
	return strings.TrimSpace(strings.Join(lines[lineStartIndex:lineEndIndex], "\n"))
}

func getMethodReturnType(part []byte) (retType ReturnType) {
	exp := regexp.MustCompile("\\*{2}@Returns\\*{2}\\s\\`\\{(.*)\\}.*\\-\\s(.*)")
	matches := exp.FindSubmatch(part)
	if len(matches) >= 2 {
		retType.ReturnType = string(matches[1])
	}
	if len(matches) >= 3 {
		retType.Description = string(matches[2])
	}
	return
}
