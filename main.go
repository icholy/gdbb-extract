package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var r *regexp.Regexp

const BREAK_COMMENT_TEMPLATE = `
break {{.FilePath}}:{{.LineNumber}} {{formatCondition .Condition}}
commands
{{range .Commands}}{{.}}
{{end}}end
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

func main() {

	lines := []string{
		"//break",
		"//break if x > 1",
		"//break : print x",
		"//break if y == 2 : info locals; continue",
	}

	for _, line := range lines {
		bp := ParseLine(line)
		fmt.Println(line)
		fmt.Printf("%#v\n", *bp)
	}

}
