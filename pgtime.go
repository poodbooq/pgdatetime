package pgdatetime

import (
	"time"

	"github.com/go-pg/pg/v10/types"
)

type PGTime struct {
	Hour   uint
	Minute uint
	Second uint
}

var _ types.ValueAppender = (*PGTime)(nil)

func (pgt PGTime) AppendValue(b []byte, flags int) ([]byte, error) {
	if flags == 1 {
		b = append(b, '\'')
	}

	b = pgt.TimeStd().AppendFormat(b, pgTimeFormat)

	if flags == 1 {
		b = append(b, '\'')
	}
	return b, nil
}

var _ types.ValueScanner = (*PGTime)(nil)

func (pgt *PGTime) ScanValue(rd types.Reader, n int) error {
	if n <= 0 {
		return nil
	}
	var t time.Time

	tmp, err := rd.ReadFullTemp()
	if err != nil {
		return err
	}

	t, err = time.ParseInLocation(pgTimeFormat, string(tmp), time.UTC)
	if err != nil {
		return err
	}

	pgtTmp := ToPgTimeUTC(&t)
	pgt.Hour = pgtTmp.Hour
	pgt.Minute = pgtTmp.Minute
	pgt.Second = pgtTmp.Second

	return nil
}
