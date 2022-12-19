package main

type Markup struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Value string `json:"value"`
}

type Dto struct {
	Markup Markup `json:"markup"`
}
