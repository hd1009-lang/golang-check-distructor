package models

import (
	"database/sql"
	"director/database"
	"director/types"
	"fmt"
	"log"
)

func GetDistributorNews(page, limit uint16) []*types.DistributorNewsItem {
	var distributorNewsList []*types.DistributorNewsItem
	db := database.MariaDB()
	defer db.Close()
	query := `select dn.supplier_id, dn.distributor_news_id, dn.categories_id, dn.update_at, dn.expiry_date, dn.enable, dn.content, dn.title
			  from distributor_news dn
			  LIMIT ? OFFSET ?
`
	var supplier_id, distributor_news_id, categories_id uint64
	var update_at, expiry_date, content, title sql.NullString
	var enable uint8
	if limit == 0 {
		query := `select dn.supplier_id, dn.distributor_news_id, dn.categories_id, dn.update_at, dn.expiry_date, dn.enable, dn.content, dn.title
			  		from distributor_news dn`
		result, err := db.Query(query)
		defer result.Close()
		if err != nil {
			log.Println("GetNews 1", err.Error())
		}
		for result.Next() {
			err := result.Scan(&supplier_id, &distributor_news_id, &categories_id, &update_at, &expiry_date, &enable, &content, &title)
			if err != nil {
				log.Println("GetNews 2", err.Error())
			}

			active := false
			if enable == 1 {
				active = true
			}
			news := types.DistributorNewsItem{Supplier_id: supplier_id, Distributor_news_id: distributor_news_id, Categories_id: categories_id, Update_at: update_at.String, Expiry_date: expiry_date.String, Enable: active, Content: content.String, Title: title.String}
			distributorNewsList = append(distributorNewsList, &news)

		}
	} else {
		result, err := db.Query(query, limit, (limit * (page - 1)))
		if err != nil {
			fmt.Println("GetProduct 3", err.Error())
		}
		for result.Next() {
			err := result.Scan(&supplier_id, &distributor_news_id, &categories_id, &update_at, &expiry_date, &enable, &content, &title)
			if err != nil {
				fmt.Println("GetProduct 4", err.Error())
			}
			active := false
			if enable == 1 {
				active = true
			}
			news := types.DistributorNewsItem{Supplier_id: supplier_id, Distributor_news_id: distributor_news_id, Categories_id: categories_id, Update_at: update_at.String, Expiry_date: expiry_date.String, Enable: active, Content: content.String, Title: title.String}
			distributorNewsList = append(distributorNewsList, &news)
		}
	}
	return distributorNewsList
}
