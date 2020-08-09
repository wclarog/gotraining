package event

import (
	"context"
	"errors"
	goCsv "github.com/JensRantil/go-csv"
	"io"
	"os"
	"time"
)

type Repository interface {
	NextEvents(ctx context.Context) ([]DTOEvent, error)
}

type repository struct {
	fileName string
}

func NewRepository(fileName string) Repository {
	return &repository{
		fileName: fileName,
	}
}

func (r repository) NextEvents(_ context.Context) ([]DTOEvent, error) {
	var results []DTOEvent

	csvFile, errOpen := os.Open(r.fileName)
	defer csvFile.Close()
	if errOpen != nil {
		return []DTOEvent{}, errOpen
	}

	dialect := goCsv.Dialect {
		Delimiter: ',',
		Quoting: goCsv.QuoteNonNumeric,
		DoubleQuote: goCsv.DoDoubleQuote,
		QuoteChar: '"',
	}

	csvReader := goCsv.NewDialectReader(csvFile, dialect)
	/*
	csvReader.Comma = ','
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = 3
	csvReader.LazyQuotes = false
	*/

	// First line is header, data discarded.
	record, err := csvReader.Read()
	if err == io.EOF {
		return []DTOEvent{}, errors.New("file does not exist")
	}

	for {
		// Read each record from csv
		record, err = csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return []DTOEvent{}, err
		}

		newDate, err := time.Parse("2019-12-31", record[2])
		if err != nil {
			return []DTOEvent{}, err
		}

		var newEvent = DTOEvent{
			Title:       record[0],
			Description: record[1],
			Date:        newDate,
		}

		results = append(results, newEvent)
	}

	return results, nil
/*
	return []DTOEvent{
		DTOEvent{
			Title:       "Title1",
			Description: "Description1",
			Date:        time.Date(2020, 9, 2, 0, 0, 0, 0, time.UTC),
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
*/
}
