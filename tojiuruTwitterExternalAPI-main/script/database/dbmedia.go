package database

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/sony/sonyflake"
)

type MediaFile struct {
	MediaID   string
	Md5       string
	Format    string
	CreatedAt string
}

// media_file table
// Tx
func CreateMediaFileTx(tx *sql.Tx, mediaFile *MediaFile) (mediaID string, err error) {
	statement, err := tx.Prepare("INSERT INTO media_file(media_id, md5, format) values ($1,$2,$3) on conflict(md5) do nothing")
	if err != nil {
		return
	}
	defer statement.Close()
	mediaID = genSonyflake()
	_, err = statement.Exec(mediaID, mediaFile.Md5, mediaFile.Format)
	if err != nil {
		return
	}
	return
}

func FindByMD5(md5 string) (*MediaFile, error) {
	row := Db.QueryRow(
		`SELECT
			media_id, md5, format, created_at
		FROM
			media_file
		WHERE
			md5 = ?`, md5)
	return convertToMediaFile(row)
}

func DeleteMediaFile(mediaid string) error {
	_, err := Db.Exec("DELETE from media_file WHERE media_id = ?", mediaid)
	if err != nil {
		return err
	}
	return err
}

// media table
type Media struct {
	MediaId   string
	MessageId string
	Format    string
}

type MediaList []Media

// Tx
func CreateMediaTx(tx *sql.Tx, mediaFile *Media) error {
	_, err := tx.Exec(
		"INSERT INTO mediaFile(media_id, format) VALUES(?,?)",
		mediaFile.MediaId, mediaFile.Format)
	return err
}

func UpdateMediaByMessageIDTx(tx *sql.Tx, MessageId string) error {
	_, err := tx.Exec(
		`UPDATE
			media
		SET
			tweet_id = ?`, MessageId)
	return err
}

func SelectMediaFindByMessageId(MessageId string) (*MediaList, error) {
	rows, err := Db.Query(
		`SELECT
			media_id, tweet_id, format
		FROM 
			mediaFile
		WHERE
			tweet_id`, MessageId)
	if err != nil {
		return nil, err
	}
	var medialist MediaList
	for rows.Next() {
		media, err := convertToMedia(rows)
		if err != nil {
			return nil, err
		}
		medialist = append(medialist, *media)
	}
	return &medialist, nil
}

// convertfunction
func convertToMediaFile(row *sql.Row) (*MediaFile, error) {
	mediaFile := MediaFile{}
	if err := row.Scan(
		&mediaFile.MediaID,
		&mediaFile.Md5,
		&mediaFile.Format,
		&mediaFile.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &mediaFile, nil
}

func convertToMedia(rows *sql.Rows) (*Media, error) {
	media := Media{}
	if err := rows.Scan(
		&media.MediaId,
		&media.MessageId,
		&media.Format); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &media, nil
}

func genSonyflake() (media_id string) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	media_id = strconv.FormatUint(id, 16)
	return
}
