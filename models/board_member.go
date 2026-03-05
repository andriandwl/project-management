package models

import "time"

type BoardMember struct {
	BoardInternalID int64     `json:"board_internal_id" db:"board_interna_id" gorm:"column:board_internal_id;primaryKey"`
	UserInternalID  int64     `json:"user_internal_id" db:"user_internal_id" gorm:"column:user_internal_id;primaryKey"`
	JoinedAt        time.Time `json:"joined_at" db:"joined_at"`
}
