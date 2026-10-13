package model

type Input struct {
	Input    float32    `json:"input"`
	Targets  []*Target  `json:"targets"`
	Streets  []*Street  `json:"streets"`
	Dividers []*Divider `json:"dividers"`
}

type Target struct {
	For  int     `json:"for"`
	Exit float32 `json:"exit"`
}

type Street struct {
	Name string `json:"name"`
	Ends []int  `json:"ends"`
}

type Divider struct {
	Name  string  `json:"name"`
	Left  float32 `json:"left"`
	Right float32 `json:"right"`
}

type Result struct {
	StreetResults []*StreetResult
}

type StreetResult struct {
	StreetName string
	Input      float32
	Divider    *Divider
	Nodes      []*Node
}

type Node struct {
	Input   float32
	Target  float32
	Divider *Divider
}
