package model

type Table struct {
	ID       string `json:"id"`
	Number   int    `json:"number"`
	Capacity int    `json:"capacity"`
}

type TableBook struct {
	ID           string `json:"id,omitempty"`
	TableID      string `json:"tableId,omitempty"`
	BookDate     string `json:"bookDate"`
	GuestName    string `json:"guestName"`
	GuestContact string `json:"guestContact"`
	Code         string `json:"code,omitempty"`
}
