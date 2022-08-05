package types

type DistributorNewsListParams struct {
	PageSize   uint16 `json:"page_size"`
	Page       uint16 `json:"page"`
	SupplierId uint64 `json:"supplier_id"`
}
type DistributorNewsItem struct {
	Supplier_id         uint64 `json:"supplier_id"`
	Distributor_news_id uint64 `json:"distributor_news_id"`
	Categories_id       uint64 `json:"categories_id"`
	Update_at           string `json:"update_at"`
	Expiry_date         string `json:"expiry_date"`
	Status              string `json:"status"`
	Content             string `json:"content"`
	Title               string `json:"title"`
}
type DistributorNewsForm struct {
	Categories_id uint64 `json:"categories_id" validate:"required"`
	Expiry_date   string `json:"expiry_date" validate:"required"`
	Status        string `json:"status" validate:"required"`
	Content       string `json:"content" validate:"required"`
	Title         string `json:"title" validate:"required"`
}
