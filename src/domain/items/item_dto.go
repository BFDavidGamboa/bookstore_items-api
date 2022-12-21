package items

type Item struct {
	Id                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Price             float64     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	TotalQuantity     int         `json:"total_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plaintext"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}
