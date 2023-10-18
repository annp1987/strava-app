package api

import "time"

func getBeginDay(t time.Time) int64 {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.FixedZone("UTC+7", 0)).Unix()
}

func getEndDay(t time.Time) int64 {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), time.FixedZone("UTC+7", 0)).Unix()
}

func GetDateTime(queryTime string) (begin, end int64) {
	layouts := []string{
		"2006/01/02",
		"01/02/2006",
		"Jan 02, 2006",
		"2006-01-02",
		time.RFC3339,
		time.RFC3339Nano,
	}

	var parsedTime time.Time
	var err error

	// Try parsing the input string with each layout until successful
	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, queryTime)
		if err == nil {
			break
		}
	}
	if err != nil {
		parsedTime = time.Now()
	}
	y, m, d := parsedTime.Date()
	begin = time.Date(y, m, d, 0, 0, 0, 0, time.UTC).Unix()
	end = time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC).Unix()
	return begin, end
}
