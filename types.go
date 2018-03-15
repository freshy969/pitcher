package pitcher

import (
	"database/sql"
	"encoding/json"
)

// NullString that can be marshalled with null value
type NullString struct {
	sql.NullString
}

// MarshalJSON that returns null-values on NULL sql columns
func (r NullString) MarshalJSON() ([]byte, error) {
	if r.Valid {
		return json.Marshal(r.String)
	}
	return json.Marshal(nil)
}

// Artist structure
type Artist struct {
	GID  string `db:"gid" json:"mbid"`
	Name string `db:"name"json:"name"`
}

// ReleaseDate structure
type ReleaseDate struct {
	Year  NullString `db:"date_year" json:"year"`
	Month NullString `db:"date_month" json:"month"`
	Day   NullString `db:"date_day" json:"day"`
}

// Album structure
type Album struct {
	ID          int          `db:"id" json:"-"`
	GID         string       `db:"gid" json:"mbid"`
	Name        string       `db:"name" json:"name"`
	ReleaseDate *ReleaseDate `json:"release_date"`
}

// Track structure
type Track struct {
	GID            string  `db:"gid" json:"mbid"`
	RecordingID    string  `db:"recording_id" json:"recording_mbid"`
	Name           string  `db:"name" json:"name"`
	MediumPosition int     `db:"medium_position" json:"medium_position"`
	Position       int     `db:"position" json:"position"`
	Length         int     `db:"length" json:"length"`
	Artist         *Artist `json:"artist"`
	Album          *Album  `json:"album"`
}
