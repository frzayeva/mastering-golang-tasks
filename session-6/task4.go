package main

import "fmt"

func get_capital(country string, countries map[string]string) string {

	if capital, ok := countries[country]; ok {
		return capital
	} else {
		return fmt.Sprintf("Capital of %s is not found", country)
	}
}

func main() {
	countries := map[string]string{
		"Azerbaijan": "Baku",
		"USA":        "Washington",
		"Germany":    "Berlin",
		"Japan":      "Tokyo",
	}
	countries["Russia"] = "Moscow"
	countries["Ireland"] = "Dublin"

	for country, capital := range countries {
		fmt.Printf("%s : %s\n", country, capital)
	}
	fmt.Println(get_capital("France", countries))
}
