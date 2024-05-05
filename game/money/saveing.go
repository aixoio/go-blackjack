package money

import (
	"os"
	"strconv"
)

const FILE_NAME string = "./blackjack-save-file-0.dat"

func SaveBalenceToFile(debug bool) error {
	base := 2
	if debug {
		base = 10
	}

	data := strconv.FormatInt(Balence, base)

	return os.WriteFile(FILE_NAME, []byte(data), 0644)
}

func LoadBalenceFromFile(debug bool) error {
	data, err := os.ReadFile(FILE_NAME)
	if err != nil {
		return err
	}

	base := 2
	if debug {
		base = 10
	}

	bal, err := strconv.ParseInt(string(data), base, 64)
	if err != nil {
		return err
	}

	Balence = bal

	return nil
}
