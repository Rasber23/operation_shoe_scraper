package entity

type article struct {
	category      string
	brand         string
	articleNumber string
	articleName   string
	price         string
	outerMaterial string
	lining        string
	sole          string
}

type articles []article
