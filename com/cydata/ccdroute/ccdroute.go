package ccdroute

type CcdRoute struct {
	id          int    `json:"id"`
	address     string `json:"address"`
	mask        string `json:"mask"`
	description string `json:"description"`
}
