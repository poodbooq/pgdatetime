package pgdatetime

import (
	"time"

	"github.com/go-pg/pg/v10/types"
)

type PGDate struct {
	Year  uint
	Month uint
	Day   uint
}

var _ types.ValueAppender = (*PGTime)(nil)

func (pgd PGDate) AppendValue(b []byte, flags int) ([]byte, error) {
	if flags == 1 {
		b = append(b, '\'')
	}

	b = pgd.TimeStd().AppendFormat(b, pgDateFormat)

	if flags == 1 {
		b = append(b, '\'')
	}

	return b, nil
}

var _ types.ValueScanner = (*PGTime)(nil)

func (pgd *PGDate) ScanValue(rd types.Reader, n int) error {
	if n <= 0 {
		return nil
	}
	var t time.Time

	tmp, err := rd.ReadFullTemp()
	if err != nil {
		return err
	}

	t, err = time.ParseInLocation(pgDateFormat, string(tmp), time.UTC)
	if err != nil {
		return err
	}

	pgdTmp := ToPgDateUTC(&t)
	pgd.Year = pgdTmp.Year
	pgd.Month = pgdTmp.Month
	pgd.Day = pgdTmp.Day

	return nil
}
