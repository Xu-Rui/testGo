package main

import (
	"fmt"

	//"encoding/json"
	//"time"
	//"strconv"
	//"helper/common"
	"time"
	//"strconv"
)

type NationTop struct {
	AverageStar  float32 `json:"average_star"`
	AverageDirty  float32 `json:"dity_score"`
	AverageGuide float32 `json:"guide_score"`
	AverageDrive float32 `json:"drive_score"`
	AverageCivil  float32 `json:"civil_score"`
	Timestamp       int64   `json:"timestamp"`
}

func cetStartTimestamp(date string) (time.Time, bool) {
		startDay := time.Now().AddDate(0, 0, -1)
		y, m, d := startDay.Date()
		return time.Date(y, m, d, 0, 0, 0, 0, time.Local), true

}

func main() {
	//var n NationTop
	//topInfo, _ := json.Marshal(n)
	//fmt.Printf("%s",topInfo)
	checkDate, _ := cetStartTimestamp("")
	fmt.Println(checkDate)
}