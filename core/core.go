package core

import (
	"bufio"
	"os"
	"strings"
)

func ProcessInput(str, banner string) (string, error) {
	filePath := "models/" + banner + ".txt"
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var allResults [][]string
	lines := strings.Split(str, "\\n")
	for _, line := range lines {
		var result []string
		i := 0
		for _, r := range line {
			file.Seek(0, 0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if int(r)-32 == i/9 {
					line := scanner.Text()
					result = append(result, line)
				}
				i++
			}
			if err := scanner.Err(); err != nil {
				return "", err
			}
			i = 0
		}
		allResults = append(allResults, result)
	}

	var output strings.Builder
	empty := true
	for i, result := range allResults {
		if len(allResults)-1 == i && i != 0 && empty && len(result) == 0 {
			break
		}

		if len(result) == 0 {
			output.WriteString("\n")
			continue
		} else {
			empty = false
		}

		for j := 1; j < 9; j++ {
			for k := range result {
				if j+k*9 < len(result) {
					output.WriteString(result[j+k*9])
				}
			}
			output.WriteString("\n")
		}
	}

	return output.String(), nil
}
