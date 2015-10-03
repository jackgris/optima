package models

type Advertiser struct {
	Age       int      `json:"age"`
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Sex       string   `json:"sex"`
	Nse       string   `json:"nse"`
	Coverage  string   `json:"coverage"`
	Interets  []string `json:"interests"`
	Category  string   `json:"category"`
	Budget    int      `json:"budget"`
	Objetives string   `json:"objetives"`
}
