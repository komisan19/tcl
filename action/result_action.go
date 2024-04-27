package action

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type Target struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status int    `json:"status"`
}

func ResultAction(file *string) ([][]string, error) {
	jsonFile, err := os.Open(*file)
	if err != nil {
		slog.Error("json file isn't open.", err)
		return nil, err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		slog.Error("json file don't read", err)
		return nil, err
	}

	var target map[string][]Target
	json.Unmarshal(jsonData, &target)

	var data [][]string
	for _, s := range target["targets"] {
		result, err := PingAction(s.URL)
		if err != nil {
			slog.Error("%v: this URL isn't open", s.URL, err)
			return nil, err
		}
		data = append(data, []string{s.Name, s.URL, fmt.Sprint(result)})
	}

	return data, nil
}
