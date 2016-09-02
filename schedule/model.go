package schedule

type Job struct {
	//CodeName string `json:"code_name"`
	Image   string `json:"image"`
	Payload string `json:"payload"`
}

type ReqData struct {
	Jobs []*Job `json:"jobs"`
}
