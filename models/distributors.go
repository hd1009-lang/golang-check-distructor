package models

import (
	"database/sql"
	"director/database"
	"director/helpers"
	"director/types"
	"fmt"
	"log"
	"time"
)

func GetDistributorNews(page, limit uint16) []*types.DistributorNewsItem {
	var distributorNewsList []*types.DistributorNewsItem
	db := database.MariaDB()
	defer db.Close()
	query := `select dn.supplier_id, dn.distributor_news_id, dn.categories_id, dn.update_at, dn.expiry_date, dn.status, dn.content, dn.title
			  from distributor_news dn
			  LIMIT ? OFFSET ?
`
	var supplier_id, distributor_news_id, categories_id uint64
	var update_at, expiry_date, content, title, status sql.NullString
	if limit == 0 {
		query := `select dn.supplier_id, dn.distributor_news_id, dn.categories_id, dn.update_at, dn.expiry_date, dn.status, dn.content, dn.title
			  		from distributor_news dn`
		result, err := db.Query(query)
		defer result.Close()
		if err != nil {
			log.Println("GetNews 1", err.Error())
		}
		for result.Next() {
			err := result.Scan(&supplier_id, &distributor_news_id, &categories_id, &update_at, &expiry_date, &status, &content, &title)
			if err != nil {
				log.Println("GetNews 2", err.Error())
			}
			news := types.DistributorNewsItem{Supplier_id: supplier_id, Distributor_news_id: distributor_news_id, Categories_id: categories_id, Update_at: update_at.String, Expiry_date: expiry_date.String, Status: status.String, Content: content.String, Title: title.String}
			distributorNewsList = append(distributorNewsList, &news)

		}
	} else {
		result, err := db.Query(query, limit, (limit * (page - 1)))
		if err != nil {
			fmt.Println("GetProduct 3", err.Error())
		}
		for result.Next() {
			err := result.Scan(&supplier_id, &distributor_news_id, &categories_id, &update_at, &expiry_date, &status, &content, &title)
			if err != nil {
				fmt.Println("GetProduct 4", err.Error())
			}

			news := types.DistributorNewsItem{Supplier_id: supplier_id, Distributor_news_id: distributor_news_id, Categories_id: categories_id, Update_at: update_at.String, Expiry_date: expiry_date.String, Status: status.String, Content: content.String, Title: title.String}
			distributorNewsList = append(distributorNewsList, &news)
		}
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
