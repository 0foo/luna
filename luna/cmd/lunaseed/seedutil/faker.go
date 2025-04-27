package seedutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func CallFakerDynamic(name string) (string, error) {
	name = strings.ToLower(strings.TrimSpace(name))

	parts := strings.Split(name, "|")
	funcName := parts[0]

	lookup := gofakeit.GetFuncLookup(funcName)
	if lookup == nil {
		return "", fmt.Errorf("faker function %s not found", funcName)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var params *gofakeit.MapParams
	if len(parts) > 1 {
		mp := make(gofakeit.MapParams)
		for _, p := range parts[1:] {
			kv := strings.SplitN(p, "=", 2)
			if len(kv) != 2 {
				return "", fmt.Errorf("invalid parameter format: %s", p)
			}
			key := strings.TrimSpace(kv[0])
			val := strings.TrimSpace(kv[1])
	
			mp[key] = gofakeit.MapParamsValue{
				Value: autoConvertValue(val),
			}
		}
		params = &mp
	}
	

	val, err := lookup.Generate(r, params, nil)
	if err != nil {
		return "", fmt.Errorf("error generating faker value for %s: %w", funcName, err)
	}

	return fmt.Sprintf("%v", val), nil
}

func autoConvertValue(val string) string {
	return strings.TrimSpace(val)
}
