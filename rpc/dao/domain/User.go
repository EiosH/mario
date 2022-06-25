package domain

import (
	"time"
)

type User struct {
	ID         int64     `json:"id,omitempty"`
	Code       string    `json:"code,omitempty"`
	Name       string    `json:"name,omitempty"`
	NickName   string    `json:"nick_name,omitempty"`
	IdCard     string    `json:"id_card,omitempty"`
	Mobile     string    `json:"mobile,omitempty"`
	CreateTime time.Time `json:"create_time"`
	ModifyTime time.Time `json:"modify_time"`
}