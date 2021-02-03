package pgdatetime

import (
	"fmt"
	"time"
)

const (
	pgTimeFormat = "15:04:05"
	pgTimeTmpl   = "%d:%d:%d"
	pgDateFormat = "2006-01-02"
	pgDateTmpl   = "%d-%d-%d"
)

func ToPgDateUTC(t *time.Time) PGDate {
	return PGDate{
		Year:  uint(t.Year()),
		Month: uint(t.Month()),
		Day:   uint(t.Day()),
	}
}

func ToPgTimeUTC(t *time.Time) PGTime {
	return PGTime{
		Hour:   uint(t.Hour()),
		Minute: uint(t.Minute()),
		Second: uint(t.Second()),
	}
}

func PgTimeDateToStd(pgd *PGDate, pgt *PGTime) time.Time {
	return time.Date(
		int(pgd.Year),
		time.Month(pgd.Month),
		int(pgd.Day),
		int(pgt.Hour),
		int(pgt.Minute),
		int(pgt.Second),
		0,
		time.UTC,
	)
}

func PgTimeDateToStdInLocation(pgd *PGDate, pgt *PGTime, loc *time.Location) time.Time {
	return time.Date(
		int(pgd.Year),
		time.Month(pgd.Month),
		int(pgd.Day),
		int(pgt.Hour),
		int(pgt.Minute),
		int(pgt.Second),
		0,
		loc,
	)
}

func (pgt PGTime) String() string {
	return fmt.Sprintf(pgTimeTmpl, pgt.Hour, pgt.Minute, pgt.Second)
}

func (pgd PGDate) String() string {
	return fmt.Sprintf(pgDateTmpl, pgd.Year, pgd.Month, pgd.Day)
}

func (pgt PGTime) TimeStd() time.Time {
	var nullTime time.Time
	return time.Date(
		nullTime.Year(),
		nullTime.Month(),
		nullTime.Day(),
		int(pgt.Hour),
		int(pgt.Minute),
		int(pgt.Second),
		nullTime.Nanosecond(),
		time.UTC,
	)
}

func (pgd PGDate) TimeStd() time.Time {
	var nullTime time.Time
	return time.Date(
		int(pgd.Year),
		time.Month(pgd.Month),
		int(pgd.Day),
		nullTime.Hour(),
		nullTime.Minute(),
		nullTime.Second(),
		nullTime.Nanosecond(),
		time.UTC,
	)
}
