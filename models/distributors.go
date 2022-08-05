package models

import (
	"database/sql"
	"director/database"
	"director/helpers"
	"director/types"
	"fmt"
	"log"
	"strings"
	"time"
)

func GetDistributorNews(params *types.DistributorNewsListParams) []*types.DistributorNewsItem {
	var conditions []string
	var conditionQuery string
	var distributorNewsList []*types.DistributorNewsItem

	var supplier_id, distributor_news_id, categories_id uint64
	var update_at, expiry_date, content, title, status, create_at sql.NullString

	offset := (params.Page - 1) * params.PageSize

	if params.SupplierId != 0 {
		conditions = append(conditions, fmt.Sprintf("supplier_id = %d", params.SupplierId))
	}
	if len(conditions) > 0 {
		conditionQuery = fmt.Sprintf("WHERE %s", strings.Join(conditions, "AND"))
	}
	queryString := fmt.Sprintf("SELECT * FROM distributor_news %s LIMIT %d OFFSET %d", conditionQuery, params.PageSize, offset)
	db := database.MariaDB()
	defer db.Close()
	result, err := db.Query(queryString)
	defer result.Close()
	if err != nil {
		log.Println("GetNews 1", err.Error())
	}
	for result.Next() {
		err := result.Scan(&supplier_id, &update_at, &distributor_news_id, &categories_id, &create_at, &expiry_date, &status, &content, &title)
		if err != nil {
			log.Println("GetNews 2", err.Error())
		}
		news := types.DistributorNewsItem{Supplier_id: supplier_id, Distributor_news_id: distributor_news_id, Categories_id: categories_id, Update_at: update_at.String, Expiry_date: expiry_date.String, Status: status.String, Content: content.String, Title: title.String}
		distributorNewsList = append(distributorNewsList, &news)

	}
	return distributorNewsList
}

func CreateDistributorNews(data *types.DistributorNewsForm, supplier_id uint64) error {
	db := database.MariaDB()
	defer db.Close()
	distributor_news_id := helpers.GenerateId()
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now().In(loc)
	_, err := db.Exec("INSERT INTO distributor_news (supplier_id, create_at, update_at, distributor_news_id, categories_id, status, expiry_date, content,title) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?,? ) ",
		supplier_id, now, now, distributor_news_id, data.Categories_id, data.Status, data.Expiry_date, data.Content, data.Title)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

func DeleteNews(id string) error {
	db := database.MariaDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM distributor_news WHERE distributor_news_id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
func UpdateNews(id string, data *types.DistributorNewsForm) error {
	db := database.MariaDB()
	defer db.Close()
	now := time.Now()
	_, err := db.Exec("update distributor_news set title = ?, content = ?, expiry_date = ?,  update_at=?, categories_id=? where distributor_news_id = ?",
		data.Title, data.Content, data.Expiry_date, now, data.Categories_id, id)
	if err != nil {
		fmt.Println("UpdateNews", err.Error())
		return err
	}
	return nil
}

func UpdateStatusNews(id string, status string) error {
	db := database.MariaDB()
	defer db.Close()
	now := time.Now()
	_, err := db.Exec("update distributor_news set   update_at=?, status=? where distributor_news_id = ?",
		now, status, id)
	if err != nil {
		fmt.Println("UpdateNews", err.Error())
		return err
	}
	return nil
}
