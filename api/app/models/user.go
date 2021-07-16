package models

import "time"

type User struct {
	id          int       `xorm:"not null pk autoincr INT(10)"`
	name        string    `xorm:"not null comment('分类名') VARCHAR(255)"`
	password    string    `xorm:"not null"`
	displayName string    `xorm:"not null comment('分类别名') index VARCHAR(255)"`
	createdAt   time.Time `xorm:"created not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	updatedAt   time.Time `xorm:"updated not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
