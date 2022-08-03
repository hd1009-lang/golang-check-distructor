package types

type DistributorNewsList struct {
	PageSize uint16 `json:"page_size"`
	Page     uint16 `json:"page"`
}
type DistributorNewsItem struct {
	Supplier_id         uint64 `json:"supplier_id"`
	Distributor_news_id uint64 `json:"distributor_news_id"`
	Categories_id       uint64 `json:"categories_id"`
	Update_at           string `json:"update_at"`
	Expiry_date         string `json:"expiry_date"`
	Enable              bool   `json:"enable"`
	Content             string `json:"content"`
	Title               string `json:"title"`
}
type DistributorNewsForm struct {
	Supplier_id         string `json:"supplier_id"`
	Distributor_news_id string `json:"distributor_news_id"`
	Categories_id       string `json:"categories_id"`
	Expiry_date         string `json:"expiry_date"`
	Enable              bool   `json:"enable"`
	Content             string `json:"content"`
	Title               string `json:"title"`
}
