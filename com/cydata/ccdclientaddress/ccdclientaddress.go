package ccdclientaddress

type CcdClientAddress struct {
	id            int    `json:"id"`
	accountId     int    `json:"accountId"`
	clientAddress string `json:"clientAddress"`
	mask          string `json:"mask"`
}
