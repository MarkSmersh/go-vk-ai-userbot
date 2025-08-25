package utils

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func LoadDotEnv(file string) {
	envs := map[string]string{}

	fileData, err := os.ReadFile(file)

	if err != nil {
		slog.Warn(".env file is missing")
		return
	}

	data := strings.Split(string(fileData), "\n") // ranging through SplitSeq is more effective blah-blah-blah

	for _, l := range data {
		line := strings.Split(l, "=")

		if line != nil && len(line) >= 2 {
			k := line[0]
			v := line[1]

			if k != "" && v != "" {
				if os.Getenv(k) == "" {
					os.Setenv(k, v)
				}
				envs[k] = v
			}
		}
	}
}

func GetEnvInt(key string) int {
	str := os.Getenv(key)

	if str == "" {
		return 0
	}

	v, err := strconv.Atoi(str)

	if err != nil {
		slog.Warn(
			fmt.Sprintf("Env variable %s has value %s, but expected int type", key, str),
		)
	}

	return v
}

func GetEnvFloat(key string) float64 {
	str := os.Getenv(key)

	if str == "" {
		return 0
	}

	v, err := strconv.ParseFloat(str, 64)

	if err != nil {
		slog.Warn(
			fmt.Sprintf("Env variable %s has value %s, but expected float type", key, str),
		)
	}

	return v
}
