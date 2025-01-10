package data

type ClientData struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     *struct {
		ID         int    `json:"id"`
		InboundID  int    `json:"inboundId"`
		Enable     bool   `json:"enable"`
		Email      string `json:"email"`
		Upload     int64  `json:"up"`
		Download   int64  `json:"down"`
		ExpiryTime int    `json:"expiryTime"`
		Total      int    `json:"total"`
		Reset      int    `json:"reset"`
	} `json:"obj"`
}