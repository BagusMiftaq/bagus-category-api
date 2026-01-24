package model

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"desc"`
}

var ListCategory = []Category{
	{ID: 1, Name: "Ensiklopedia Bahasa Semut", Description: "Mystical Science"},
	{ID: 2, Name: "99 Dongeng Sebelum Kerja", Description: "Myth"},
	{ID: 3, Name: "Kisi-Kisi Pertanyaan Alam Kubur", Description: "Spiritual"},
	{ID: 4, Name: "Merantau ke Framework Seberang", Description: "Adventure"},
}