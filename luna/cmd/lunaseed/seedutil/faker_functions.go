
package seedutil

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

var fakerFunctions = map[string]func(params map[string]string) (string, error){
	"firstname": func(params map[string]string) (string, error) {
		return gofakeit.FirstName(), nil
	},
	"lastname": func(params map[string]string) (string, error) {
		return gofakeit.LastName(), nil
	},
	"email": func(params map[string]string) (string, error) {
		return gofakeit.Email(), nil
	},
	"city": func(params map[string]string) (string, error) {
		return gofakeit.City(), nil
	},
	"sentence": func(params map[string]string) (string, error) {
		words := 10
		if w, ok := params["words"]; ok {
			if n, err := strconv.Atoi(w); err == nil {
				words = n
			}
		}
		return gofakeit.Sentence(words), nil
	},
	"randomstringn": func(params map[string]string) (string, error) {
		length := 8
		if l, ok := params["length"]; ok {
			if n, err := strconv.Atoi(l); err == nil {
				length = n
			}
		}
		return gofakeit.LetterN(uint(length)), nil
	},
	"randomnumbern": func(params map[string]string) (string, error) {
		length := 8
		if l, ok := params["length"]; ok {
			if n, err := strconv.Atoi(l); err == nil {
				length = n
			}
		}
		return gofakeit.DigitN(uint(length)), nil
	},
	"password": func(params map[string]string) (string, error) {
		length := 12
		if l, ok := params["length"]; ok {
			if n, err := strconv.Atoi(l); err == nil {
				length = n
			}
		}
		return gofakeit.Password(true, true, true, true, true, length), nil
	},
	"uuid": func(params map[string]string) (string, error) {
		return gofakeit.UUID(), nil
	},
	"phone": func(params map[string]string) (string, error) {
		return gofakeit.Phone(), nil
	},
	"street": func(params map[string]string) (string, error) {
		return gofakeit.Street(), nil
	},
	"zip": func(params map[string]string) (string, error) {
		return gofakeit.Zip(), nil
	},
	"state": func(params map[string]string) (string, error) {
		return gofakeit.State(), nil
	},
	"country": func(params map[string]string) (string, error) {
		return gofakeit.Country(), nil
	},
	"latitude": func(params map[string]string) (string, error) {
		return fmt.Sprintf("%f", gofakeit.Latitude()), nil
	},
	"longitude": func(params map[string]string) (string, error) {
		return fmt.Sprintf("%f", gofakeit.Longitude()), nil
	},
}

func CallFakerDynamic(input string) (string, error) {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, "|")
	funcName := strings.ToLower(strings.TrimSpace(parts[0]))

	fn, ok := fakerFunctions[funcName]
	if !ok {
		return "", fmt.Errorf("faker function '%s' not found", funcName)
	}

	params := make(map[string]string)
	for _, p := range parts[1:] {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) == 2 {
			params[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	return fn(params)
}
