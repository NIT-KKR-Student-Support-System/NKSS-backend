// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: announcement.sql

package query

import (
	"context"
	"time"
)

const getAcademicAnnouncements = `-- name: GetAcademicAnnouncements :exec
SELECT date_of_creation, title, title_link, kind
FROM academic_announcement
ORDER BY date_of_creation DESC
`

type GetAcademicAnnouncementsRow struct {
	DateOfCreation time.Time `json:"date_of_creation"`
	Title          string    `json:"title"`
	TitleLink      string    `json:"title_link"`
	Kind           string    `json:"kind"`
}

func (q *Queries) GetAcademicAnnouncements(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, getAcademicAnnouncements)
	return err
}

const insertAcademicAnnouncement = `-- name: InsertAcademicAnnouncement :exec
INSERT INTO academic_announcement (
    date_of_creation, title, title_link, kind
)
VALUES (
    $1,
    $2,
    $3,
    'academic'
) ON CONFLICT (date_of_creation, title)
DO NOTHING
`

type InsertAcademicAnnouncementParams struct {
	DateOfCreation time.Time `json:"date_of_creation"`
	Title          string    `json:"title"`
	TitleLink      string    `json:"title_link"`
}

func (q *Queries) InsertAcademicAnnouncement(ctx context.Context, arg InsertAcademicAnnouncementParams) error {
	_, err := q.db.ExecContext(ctx, insertAcademicAnnouncement, arg.DateOfCreation, arg.Title, arg.TitleLink)
	return err
}
