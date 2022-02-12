package service

func init() {
	ScheduleTable()
}

func ScheduleTable() {
	go SignService()

}
