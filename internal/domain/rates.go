package domain

import (
	"encoding/xml"
	"time"
)

type Rates struct {
	XMLName     xml.Name `xml:"rates"`
	Text        string   `xml:",chardata"`
	Generator   string   `xml:"generator"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Copyright   string   `xml:"copyright"`
	Date        string   `xml:"date"`
	Item        []Item   `xml:"item"`
}

type Item struct {
	Text        string `xml:",chardata"`
	Fullname    string `xml:"fullname"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Quant       string `xml:"quant"`
	Index       string `xml:"index"`
	Change      string `xml:"change"`
}

type ItemDTO struct {
	Title string    `db:"title" json:"title"`
	Code  string    `db:"code" json:"code"`
	Value float64   `db:"value" json:"value"`
	Date  time.Time `db:"a_date" json:"date"`
}
