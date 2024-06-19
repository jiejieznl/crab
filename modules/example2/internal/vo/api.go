package vo

type ApiSumReq struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type ApiSumRes struct {
	Sum int `json:"sum"`
}
