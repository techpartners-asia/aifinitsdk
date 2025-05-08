package aifinitsdk_product

import (
	"encoding/json"
	"io"
)

type Product struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	Weight         int      `json:"weight"`
	WeightVariance int      `json:"weightVariance"`
	ImgUrl         string   `json:"imgUrl"`
	ItemCode       string   `json:"itemCode"`
	CollType       int      `json:"collType"` // collection type: 1- single, 2- multiple
	UpdateTime     string   `json:"updateTime"`
	CreateTime     string   `json:"createTime"`
	Status         int      `json:"status"`
	QrCodes        string   `json:"qrCodes"`
	ItemCodes      []string `json:"itemCodes"`
	ActualImgs     []string `json:"actualImgs"`
	WeightFile     string   `json:"weightFile"`
}

func (p *Product) String() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type NewProductApplication struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Weight  int    `json:"weight"`
	QrCodes string `json:"qrCodes"`

	ImgFiles     []io.Reader `json:"-"` // product image files
	ImgFileNames []string    `json:"-"` // product image file names

	PhysicalImgFiles     []io.Reader `json:"-"` // physical image files IMPORTANT: at least 2 and bar code clearly visible
	PhysicalImgFileNames []string    `json:"-"` // physical image file names IMPORTANT: at least 2 and bar code clearly visible

	WeightFile     io.Reader `json:"-"` // docs: weight of pictures
	WeightFileName string    `json:"-"` // docs: weight of pictures
}

func (n *NewProductApplication) String() string {
	bytes, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type LastInfo struct {
	Count          int   `json:"count"`
	LastUpdateTime int64 `json:"lastUpdateTime"`
}
