package event

import (
	"bufio"
	"context"
	"errors"
	"os"
	"regexp"
	"strings"
	"time"
)

type Repository interface {
	NextEvents(ctx context.Context) ([]DTOEvent, error)
}

type repository struct {
	fileName string
	events   []DTOEvent
}

func NewRepository(fileName string) Repository {
	events, err := readData(fileName)
	if err != nil {
		return nil
	}

	return &repository{
		fileName: fileName,
		events:   events,
	}
}

func (r repository) NextEvents(_ context.Context) ([]DTOEvent, error) {
	results := make([]DTOEvent, 0, len(r.events))
	today := time.Now()

	for _, evt := range r.events {
		if evt.Date.After(today) {
			results = append(results, evt)
		}
	}

	return results, nil
}

func readData(fileName string) ([]DTOEvent, error) {
	var results []DTOEvent
	var newEvent DTOEvent
	reField := regexp.MustCompile(`".*?"`)

	f, err := os.Open(fileName)

	if err != nil {
		return []DTOEvent{}, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		newEvent, err = parseLineData(scanner.Text(), reField)

		if err != nil {
			// Invalid lines are discarded.
			continue
		}

		results = append(results, newEvent)
	}

	return results, nil
}

func parseLineData(lineText string, reField *regexp.Regexp) (DTOEvent, error) {
	fields := reField.FindAllString(lineText, -1)

	if len(fields) != 3 {
		return DTOEvent{}, errors.New("invalid input line")
	}

	trimFields := Map(fields, cleanField)

	dateField, err := time.Parse("2006-1-2", trimFields[2])
	if err != nil {
		return DTOEvent{}, errors.New("invalid date field")
	}

	var newEvent = DTOEvent{
		Title:       trimFields[0],
		Description: trimFields[1],
		Date:        dateField,
	}

	return newEvent, nil
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func cleanField(field string) string {
	return strings.Trim(field, `"`)
}

/*
func dummyData(fileName string) ([]DTOEvent, error) {
	return []DTOEvent{
		DTOEvent{
			Title:       "Title0",
			Description: "Description0",
			Date:        time.Date(2020, 9, 2, 0, 0, 0, 0, time.UTC),
		},
		DTOEvent{
			Title:       "Title1",
			Description: "Description1",
			Date:        time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		DTOEvent{
			Title:       "Title2",
			Description: "Description2",
			Date:        time.Date(2020, 10, 15, 0, 0, 0, 0, time.UTC),
		},
		DTOEvent{
			Title:       "Title3",
			Description: "Description3",
			Date:        time.Date(2020, 12, 25, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}
*/
