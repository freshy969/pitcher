package pitcher

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const trackQuery = `SELECT track.gid, rec.gid as recording_id, track.name,
       track.length, track.position, medium.position AS medium_position,
			 album.gid "album.gid", album.name "album.name",
			 artist.gid "artist.gid", artist.name "artist.name", album.id "album.id",
			 release_date.date_year "album.releasedate.date_year",
			 release_date.date_month "album.releasedate.date_month",
			 release_date.date_day "album.releasedate.date_day"
       FROM track JOIN recording AS rec ON (rec.id = track.recording)
			 JOIN artist AS artist ON artist.id = track.artist_credit
       LEFT JOIN medium ON medium.id = track.medium
       LEFT JOIN release as album ON album.id = medium.release
			 LEFT JOIN LATERAL (SELECT date_year, date_month, date_day FROM release_country WHERE release=album.id) release_date ON true
       WHERE track.gid = :gid`

type trackQueryParams struct {
	GID string `db:"gid"`
}

// CreateDB returns database connection
func CreateDB() (db *sqlx.DB, err error) {
	return sqlx.Open("postgres", "dbname=musicbrainz user=postgres port=5440 sslmode=disable search_path=musicbrainz")
}

// GetTrackData returns Track matching MusicBrainz ID
func GetTrackData(db *sqlx.DB, trackID string) (*[]Track, error) {
	params := trackQueryParams{
		GID: trackID,
	}

	tracks := []Track{}

	query, err := db.PrepareNamed(trackQuery)
	if err != nil {
		return nil, err
	}

	err = query.Select(&tracks, params)
	if err != nil {
		return nil, err
	}

	return &tracks, nil
}