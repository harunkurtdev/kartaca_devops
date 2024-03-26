package model

type Ulkeler struct {
	Ulke    string `bson:"ulke" json:"ulke"`
	Nufus   int    `bson:"nufus" json:"nufus"`
	Baskent string `bson:"baskent" json:"baskent"`
}
