package model

type Messages struct {
	GoingOut Action `json:"going_out"`
	GetIn    Action `json:"get_in"`
}

type Action struct {
	Normal []string `json:"normal"`
	Today  []string `json:"today"`
	Week   []string `json:"week"`
}
