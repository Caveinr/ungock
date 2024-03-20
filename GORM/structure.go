package GORM

import "time"

type DB struct {
	DSN string `toml:"DSN"`
}

type Auth_uid struct {
	UID  string `gorm:"primaryKey"`
	SAK  uint8  `gorm:"primaryKey"`
	Note string
}

type AuthLog struct {
	Id        uint      `gorm:"column:ID;type:int(11) unsigned;AUTO_INCREMENT;primary_key" json:"ID"`
	REQUESTER string    `gorm:"column:REQUESTER;type:varchar(16);NOT NULL" json:"REQUESTER"`
	METHOD    string    `gorm:"column:METHOD;type:enum('MFRC522','GO_CQHTTP');NOT NULL" json:"METHOD"`
	STATUS    string    `gorm:"column:STATUS;type:enum('PASSED','FAILED');NOT NULL" json:"STATUS"`
	TIME      time.Time `gorm:"column:TIME;type:datetime;default:CURRENT_TIMESTAMP" json:"TIME"`
}

func (m *AuthLog) TableName() string {
	return "auth_log"
}

//CREATE TABLE auth_log(
//ID int unsigned auto_increment primary key,
//REQUESTER varchar(16) NOT NULL,
//METHOD ENUM('MFRC522','GO_CQHTTP') NOT NULL,
//STATUS ENUM('PASSED','FAILED') NOT NULL,
//TIME datetime DEFAULT CURRENT_TIMESTAMP
//);
