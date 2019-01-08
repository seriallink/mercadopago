package mp

import (
	"time"
)

const IsoFormat = `"2006-01-02T15:04:05.000Z07:00"`

type Iso8601 time.Time

func NewIso8601(t time.Time) Iso8601 {
	return Iso8601(time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location()))
}

func (iso Iso8601) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(iso).Format(IsoFormat)), nil
}

func (iso *Iso8601) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(IsoFormat, string(data))
	if err == nil {
		*iso = Iso8601(t)
	}
	return err
}

func (iso Iso8601) Pointer() *Iso8601 {
	return &iso
}

func (iso Iso8601) Time() time.Time {
	return time.Time(iso)
}

func (iso Iso8601) String() string {
	return iso.Time().String()
}
