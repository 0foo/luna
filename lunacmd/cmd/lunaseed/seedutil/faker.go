package seedutil

import (
	"fmt"
	"strings"
)


// ParseFakerInput splits input like "sentence|words=5" into function name and parameters.
func ParseFakerInput(input string) (string, map[string]string) {
	parts := strings.Split(input, "|")
	funcName := strings.ToLower(strings.TrimSpace(parts[0]))

	params := make(map[string]string)
	for _, p := range parts[1:] {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) == 2 {
			params[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	return funcName, params
}


func GetFakeData(fields map[string]string, count int) ([]map[string]string, error) {
	rows := make([]map[string]string, 0, count)

	for i := 0; i < count; i++ {
		populated := make(map[string]string)

		for column, fakerInput := range fields {
			funcName, params := ParseFakerInput(fakerInput)

			fn, ok := FakerMap[funcName]
			if !ok {
				return nil, fmt.Errorf("faker function not found for '%s'", funcName)
			}

			fakeValue, err := fn(params)
			if err != nil {
				return nil, fmt.Errorf("error generating fake data for '%s': %w", funcName, err)
			}

			populated[column] = fakeValue
		}

		rows = append(rows, populated)
	}

	return rows, nil
}
