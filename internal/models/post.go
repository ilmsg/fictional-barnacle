package models

import (
	"fmt"
	"time"
)

type Post struct {
	Id        int
	Title     string
	Content   string
	CreatedAt time.Time
	User      *User
}

func (p *Post) CreatedAtFormat() string {
	return p.CreatedAt.Format(time.RFC822)
}

func Format(t time.Time) string {
	return fmt.Sprintf("%s %02d. %s",
		days[t.Weekday()][:3], t.Day(), months[t.Month()-1][:3],
	)
}

var days = [...]string{
	"วันอาทิตย์", "วันจันทร์", "วันอังคาร", "วันพุธ", "วันพฤหัสบดี", "วันศุกร์", "วันเสาร์"}

var months = [...]string{
	"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
	"กรกฏาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
}
