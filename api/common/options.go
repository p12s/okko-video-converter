package common

type UserData struct {
	WidthList  []int     `json:"widthList"`
	CoefList   []float64 `json:"coefList"`
	IsAddWebp  bool      `json:"isAddWebp"`
	IsCompress bool      `json:"isCompress"`
}

type UserUrl struct {
	Url string `json:"url"`
}
