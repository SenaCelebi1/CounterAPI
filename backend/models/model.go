package models

type CounterValue struct {
	Num1      int  `json:"num1"`
	Increased bool `json:"inc"`
	Decreased bool `json:"dec"`
}

type ValueResponseData struct {
	Add int `json:"add"`
	Sub int `json:"sub"`
}
