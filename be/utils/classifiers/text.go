package classifiers

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func ExtractTextContext(filename string, query string) map[string]string {
	results := make(map[string]string)

	cmd := exec.Command("grep", "-n", query, filename) // Todo: use some NLP later
	out, err := cmd.CombinedOutput()

	if err != nil {
		results["error"] = fmt.Sprintf("search encoutered some error: %v", err)
		return results
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		var snippet string = scanner.Text()
		match := strings.Split(snippet, ":")

		number := match[0]
		line := strings.Join(match[1:], "")

		results[number] = line
	}

	return results
}
