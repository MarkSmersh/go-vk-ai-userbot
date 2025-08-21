package utils

import (
	"log"
	"os"
	"strings"
)

func GetEnv(file string) map[string]string {
	envs := map[string]string{}

	fileData, err := os.ReadFile(file)

	if err != nil {
		log.Fatalln(".env is missing")
		return envs
	}

	data := strings.Split(string(fileData), "\n") // ranging through SplitSeq is more effective blah-blah-blah

	for _, l := range data {
		line := strings.Split(l, "=")

		if line != nil && len(line) >= 2 {
			k := line[0]
			v := line[1]

			if k != "" && v != "" {
				envs[k] = v
			}
		}
	}

	return envs
}
