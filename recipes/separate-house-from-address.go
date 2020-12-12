package recipes

import (
	"fmt"
	"strings"
)

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

/*
  SeparateHouseNumber from address line
  credit: https://rosettacode.org/wiki/Separate_the_house_number_from_the_street_name
*/
func SeparateHouseNumber(address string) (street string, house string) {
	addressLine := strings.Replace(address, " - ", "-", 1)
	length := len(addressLine)
	fields := strings.Fields(addressLine)
	size := len(fields)
	last := fields[size-1]
	penult := fields[size-2]

	if isDigit(last[0]) {
		isdig := isDigit(penult[0])
		if size > 2 && isdig && !strings.HasPrefix(penult, "194") {
			house = penult + " " + last
		} else {
			house = last
		}
	} else if size > 2 {
		house = penult + " " + last
	}
	street = strings.TrimRight(addressLine[:length-len(house)], " ")
	return
}

/*
	SeparateHouseNumberTest separate house
	from test scenarios
*/
func TestSeparateHouseNumber() {
	addresses := [...]string{
		"2 Koningskinderenstraat 24 H",
		"Plataanstraat 5",
		"Straat 12",
		"Straat 12 II",
		"Dr. J. Straat   12",
		"Dr. J. Straat 12 a",
		"Dr. J. Straat 12-14",
		"Dr. J. Straat 12 - 14",
		"Laan 1940 - 1945 37",
		"Plein 1940 2",
		"1213-laan 11",
		"16 april 1944 Pad 1",
		"1e Kruisweg 36",
		"Laan 1940-'45 66",
		"Laan '40-'45",
		"Langeloërduinen 3 46",
		"Marienwaerdt 2e Dreef 2",
		"Provincialeweg N205 1",
		"Rivium 2e Straat 59.",
		"Nieuwe gracht 20rd",
		"Nieuwe gracht 20rd 2",
		"Nieuwe gracht 20zw /2",
		"Nieuwe gracht 20zw/3",
		"Nieuwe gracht 20 zw/4",
		"Bahnhofstr. 4",
		"Wertstr. 10",
		"Lindenhof 1",
		"Nordesch 20",
		"Weilstr. 6",
		"Harthauer Weg 2",
		"Mainaustr. 49",
		"August-Horch-Str. 3",
		"Marktplatz 31",
		"Schmidener Weg 3",
		"Karl-Weysser-Str. 6",
	}

	print := "%-30s %-24s %s\n"
	fmt.Printf(print, "Original", "Street", "HouseNumber")
	fmt.Printf(print, "---------------", "------------", "---------------")

	for _, address := range addresses {
		street, house := SeparateHouseNumber(address)
		if house == "" {
			house = "(none)"
		}
		fmt.Printf(print, address, street, house)
	}
}
