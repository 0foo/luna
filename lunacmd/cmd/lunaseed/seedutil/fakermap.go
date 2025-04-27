package seedutil

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

// peopleFunctions contains faker functions related to people information.
var FakerMap = map[string]func(params map[string]string) (string, error){
	"firstname": func(params map[string]string) (string, error) {
		return gofakeit.FirstName(), nil
	},
	"lastname": func(params map[string]string) (string, error) {
		return gofakeit.LastName(), nil
	},
	"gender": func(params map[string]string) (string, error) {
		return gofakeit.Gender(), nil
	},
	"email": func(params map[string]string) (string, error) {
		return gofakeit.Email(), nil
	},
	"phone": func(params map[string]string) (string, error) {
		return gofakeit.Phone(), nil
	},
	"username": func(params map[string]string) (string, error) {
		return gofakeit.Username(), nil
	},
	"ssn": func(params map[string]string) (string, error) {
		return gofakeit.SSN(), nil
	},
	"jobtitle": func(params map[string]string) (string, error) {
		return gofakeit.JobTitle(), nil
	},
	"date": func(params map[string]string) (string, error) {
		return gofakeit.Date().Format("2006-01-02"), nil
	},
	"time": func(params map[string]string) (string, error) {
		return gofakeit.Date().Format("15:04:05"), nil
	},
	"weekday": func(params map[string]string) (string, error) {
		return gofakeit.WeekDay(), nil
	},
	"month": func(params map[string]string) (string, error) {
		return gofakeit.MonthString(), nil
	},
	"year": func(params map[string]string) (string, error) {
		return fmt.Sprintf("%f", gofakeit.Year()), nil
	},
	"company": func(params map[string]string) (string, error) {
		return gofakeit.Company(), nil
	},
	"bs": func(params map[string]string) (string, error) {
		return gofakeit.BS(), nil
	},
	"buzzword": func(params map[string]string) (string, error) {
		return gofakeit.BuzzWord(), nil
	},
	"street": func(params map[string]string) (string, error) {
		return gofakeit.Street(), nil
	},
	"city": func(params map[string]string) (string, error) {
		return gofakeit.City(), nil
	},
	"state": func(params map[string]string) (string, error) {
		return gofakeit.State(), nil
	},
	"stateabbr": func(params map[string]string) (string, error) {
		return gofakeit.StateAbr(), nil
	},
	"country": func(params map[string]string) (string, error) {
		return gofakeit.Country(), nil
	},
	"countryabbr": func(params map[string]string) (string, error) {
		return gofakeit.CountryAbr(), nil
	},
	"zip": func(params map[string]string) (string, error) {
		return gofakeit.Zip(), nil
	},
	"latitude": func(params map[string]string) (string, error) {
		return fmt.Sprintf("%f", gofakeit.Latitude()), nil
	},
	"longitude": func(params map[string]string) (string, error) {
		return fmt.Sprintf("%f", gofakeit.Longitude()), nil
	},
	"fruit": func(params map[string]string) (string, error) {
		return gofakeit.Fruit(), nil
	},
	"vegetable": func(params map[string]string) (string, error) {
		return gofakeit.Vegetable(), nil
	},
	"beername": func(params map[string]string) (string, error) {
		return gofakeit.BeerName(), nil
	},
	"beerstyle": func(params map[string]string) (string, error) {
		return gofakeit.BeerStyle(), nil
	},
	"beeralcohol": func(params map[string]string) (string, error) {
		return gofakeit.BeerAlcohol(), nil
	},
	"dessert": func(params map[string]string) (string, error) {
		return gofakeit.Dessert(), nil
	},
	"domainname": func(params map[string]string) (string, error) {
		return gofakeit.DomainName(), nil
	},
	"tld": func(params map[string]string) (string, error) {
		return gofakeit.DomainSuffix(), nil
	},
	"ipv4": func(params map[string]string) (string, error) {
		return gofakeit.IPv4Address(), nil
	},
	"ipv6": func(params map[string]string) (string, error) {
		return gofakeit.IPv6Address(), nil
	},
	"macaddress": func(params map[string]string) (string, error) {
		return gofakeit.MacAddress(), nil
	},
	"url": func(params map[string]string) (string, error) {
		return gofakeit.URL(), nil
	},
	"useragent": func(params map[string]string) (string, error) {
		return gofakeit.UserAgent(), nil
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
	"randomstring": func(params map[string]string) (string, error) {
		length := 8
		if l, ok := params["length"]; ok {
			if n, err := strconv.Atoi(l); err == nil {
				length = n
			}
		}
		return gofakeit.LetterN(uint(length)), nil
	},
	"randomdigit": func(params map[string]string) (string, error) {
		if startStr, okStart := params["start"]; okStart {
			if endStr, okEnd := params["end"]; okEnd {
				start, err1 := strconv.Atoi(startStr)
				end, err2 := strconv.Atoi(endStr)
				if err1 != nil || err2 != nil {
					return "", fmt.Errorf("invalid start or end parameter")
				}
				if start > end {
					return "", fmt.Errorf("start must be less than or equal to end")
				}
				value := gofakeit.Number(start, end)
				return strconv.Itoa(value), nil
			}
		}

		// fallback to length
		length := 8
		if l, ok := params["length"]; ok {
			if n, err := strconv.Atoi(l); err == nil {
				length = n
			}
		}
		return gofakeit.DigitN(uint(length)), nil
	},
	"number": func(params map[string]string) (string, error) {
		min := 0
		max := 1000
		if v, ok := params["min"]; ok {
			if n, err := strconv.Atoi(v); err == nil {
				min = n
			}
		}
		if v, ok := params["max"]; ok {
			if n, err := strconv.Atoi(v); err == nil {
				max = n
			}
		}
		return fmt.Sprintf("%d", gofakeit.Number(min, max)), nil
	},
	"float32": func(params map[string]string) (string, error) {
		min := float32(0)
		max := float32(100)
		if v, ok := params["min"]; ok {
			if n, err := strconv.ParseFloat(v, 32); err == nil {
				min = float32(n)
			}
		}
		if v, ok := params["max"]; ok {
			if n, err := strconv.ParseFloat(v, 32); err == nil {
				max = float32(n)
			}
		}
		return fmt.Sprintf("%f", gofakeit.Float32Range(min, max)), nil
	},
	"float64": func(params map[string]string) (string, error) {
		min := 0.0
		max := 100.0
		if v, ok := params["min"]; ok {
			if n, err := strconv.ParseFloat(v, 64); err == nil {
				min = n
			}
		}
		if v, ok := params["max"]; ok {
			if n, err := strconv.ParseFloat(v, 64); err == nil {
				max = n
			}
		}
		return fmt.Sprintf("%f", gofakeit.Float64Range(min, max)), nil
	},
	"color": func(params map[string]string) (string, error) {
		return gofakeit.Color(), nil
	},
	"hexcolor": func(params map[string]string) (string, error) {
		return gofakeit.HexColor(), nil
	},
	"sentence": func(params map[string]string) (string, error) {
		wordCount := 10
		if val, ok := params["words"]; ok {
			parsed, err := strconv.Atoi(val)
			if err == nil {
				wordCount = parsed
			}
		}
		return gofakeit.Sentence(wordCount), nil
	},

	"paragraph": func(params map[string]string) (string, error) {
		numParagraphs := 1
		sentencesPerParagraph := 3
		wordsPerSentence := 12

		if val, ok := params["paragraphs"]; ok {
			parsed, err := strconv.Atoi(val)
			if err == nil {
				numParagraphs = parsed
			}
		}
		if val, ok := params["sentences"]; ok {
			parsed, err := strconv.Atoi(val)
			if err == nil {
				sentencesPerParagraph = parsed
			}
		}
		if val, ok := params["words"]; ok {
			parsed, err := strconv.Atoi(val)
			if err == nil {
				wordsPerSentence = parsed
			}
		}

		return gofakeit.Paragraph(numParagraphs, sentencesPerParagraph, wordsPerSentence, " "), nil
	},
	"timestamp": func(params map[string]string) (string, error) {
		return gofakeit.Date().Format("2006-01-02 15:04:05"), nil
	},
	"static": func(params map[string]string) (string, error) {
		val, ok := params["value"]
		if !ok {
			return "", fmt.Errorf("missing 'value' parameter for static")
		}
		return val, nil
	},
	"existingid": func(params map[string]string) (string, error) {
		table, ok := params["table"]
		if !ok || table == "" {
			return "", fmt.Errorf("missing 'table' parameter for existingid")
		}

		ids, err := LoadIDsFromTable(table, "id")
		if err != nil {
			return "", err
		}
		if len(ids) == 0 {
			return "", fmt.Errorf("no IDs found in table '%s'", table)
		}

		randomIndex := rand.Intn(len(ids))
		return strconv.Itoa(ids[randomIndex]), nil
	},


}
