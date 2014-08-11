package cellphone

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"
)

type Call struct {
	CallTime    time.Time
	CountryCode string
	Prefix      string
	Number      string
	Duration    uint
}

func LoadCalls(fileName string) ([]Call, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	calls := make([]Call, len(lines))

	for i, line := range lines {
		parts := strings.Split(line[1], " ")
		time, err := time.Parse("02/01/2006 15:04:05", line[0])
		if err != nil {
			return nil, err
		}
		duration, err := strconv.Atoi(line[2])
		if err != nil {
			return nil, err
		}

		calls[i] = Call{
			time,
			parts[0],
			parts[1],
			parts[2],
			uint(duration),
		}
	}

	return calls, nil
}
