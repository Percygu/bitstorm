package service

import "time"

type ViewPrize struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Img       string    `json:"img"`
	PrizeNum  int       `json:"prize_num"`
	PrizeCode string    `form:"prize_code"`
	PrizeTime uint      `json:"prize_time"`
	LeftNum   int       `json:"left_num"`
	PrizeType uint      `json:"prize_type"`
	PrizePlan string    `json:"prize_plan"`
	BeginTime time.Time `json:"begin_time"`
	EndTime   time.Time `json:"end_time"`
}
