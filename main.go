package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var r *regexp.Regexp

const BREAK_COMMENT_TEMPLATE = `
{{range .}}
break {{.FilePath}}:{{.LineNumber}}{{formatCondition .Condition}}
commands
{{range .Commands}}{{.}}
{{end}}end
{{end}}
`

func init() {
	var err error
	r, err = regexp.Compile(`//break\s*(?:if\s*([^\:]*))?(?:\:(.*))?`)
	if err != nil {
		log.Fatal(err)
	}
}

type BreakPoint struct {
	LineNumber          int64
	FilePath, Condition string
	Commands            []string
}

func ParseLine(line string) *BreakPoint {
	match := r.FindStringSubmatch(line)
	cmds := make([]string, 0, 1)
	if len(match[2]) > 0 {
		for _, cmd := range strings.Split(match[2], ";") {
			cmd := strings.Trim(cmd, " ")
			cmds = append(cmds, cmd)
		}
	}
	cond := strings.Trim(match[1], " ")
	bp := &BreakPoint{
		Condition: cond,
		Commands:  cmds,
	}
	return bp
}

func ParseFile(fileName string) ([]*BreakPoint, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bps := make([]*BreakPoint, 0, 1)
	scanner := bufio.NewScanner(f)
	for i := int64(0); scanner.Scan(); i++ {
		line := scanner.Text()
		if r.MatchString(line) {
			bp := ParseLine(line)
			bp.LineNumber = i
			bp.FilePath = fileName
			bps = append(bps, bp)
		}
	}
	return bps, nil
}

func ParseFiles(fileNames []string) ([]*BreakPoint, error) {
	res := make([]*BreakPoint, 0, 1)
	for _, fileName := range fileNames {
		bps, err := ParseFile(fileName)
		if err != nil {
			return nil, err
		}
		for _, bp := range bps {
			res = append(res, bp)
		}
	}
	return res, nil
}

func FormatBreakPoints(bps []*BreakPoint, w io.Writer) error {

	funcMap := template.FuncMap{
		"formatCondition": func(s string) string {
			s = strings.Trim(s, " ")
			if len(s) == 0 {
				return ""
			}
			return " if " + s
		},
	}

	tmpl, err := template.New("breakpoints").Funcs(funcMap).Parse(BREAK_COMMENT_TEMPLATE)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(w, bps); err != nil {
		return err
	}
	return nil
}

func main() {

	bps, err := ParseFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	if err := FormatBreakPoints(bps, os.Stdout); err != nil {
		log.Fatal(err)
	}

}
