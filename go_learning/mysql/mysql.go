package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Article struct {
	Aid			int64	`db:"aid"`
	Title		string	`db:"title"`
	Desc		string	`db:"desc"`
	UpdateTime 	string	`db:"update_time"`
	PublishTime	string	`db:"publish_time"`
	Status		int		`db:"status"`
	UID			int64	`db:"uid"`
	Uname		string	`db:"uname"`
	ExtData		string	`db:"ext_data"`
}

func MySQLQueryDemo() {
	PrepareDemo()
}

func QueryDemo() {
	db, err := initDb()
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Query("select aid, title, `desc`, update_time, publish_time, status, uid, uname, ext_data from tblArticle where 1=1 limit ?", 3)
	if err != nil {
		fmt.Printf("db query failed, err:%s\n", err)
		return
	}
	defer rows.Close()

	articles := []Article{}
	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Aid, &article.Title, &article.Desc, &article.UpdateTime, &article.PublishTime, &article.Status, &article.UID, &article.Uname, &article.ExtData)
		if err != nil {
			fmt.Printf("row scan failed, err:%s\n", err)
			return 
		}
		articles = append(articles, article)
	}
	fmt.Printf("articles: %+v", articles)
}

func QueryRowDemo() {
	db, err := initDb()
	if err != nil {
		return
	}
	defer db.Close()

	var article Article
	err = db.QueryRow("select aid, title, `desc`, update_time, publish_time, status, uid, uname, ext_data from tblArticle where 1=1 limit ?", 3).Scan(&article.Aid, &article.Title, &article.Desc, &article.UpdateTime, &article.PublishTime, &article.Status, &article.UID, &article.Uname, &article.ExtData)
	if err != nil {
		fmt.Printf("db query failed, err:%s\n", err)
		return
	}

	fmt.Printf("article: %+v", article)
}

func ExecDemo() {
	db, err := initDb()
	if err != nil {
		return
	}
	defer db.Close()

	res, err := db.Exec("update tblArticle set aid=1 where aid = ?", 1)
	if err != nil {
		fmt.Printf("db query failed, err:%s\n", err)
		return
	}
	lastInsertID, err := res.LastInsertId()
	rowsAffected, err := res.RowsAffected()

	fmt.Printf("lastInsertId: %d, rowsAffected: %d, err: %v", lastInsertID, rowsAffected, err)
}

func PrepareDemo() {
	db, err := initDb()
	if err != nil {
		return
	}
	defer db.Close()

	stat, err := db.Prepare("select aid, title from tblArticle where 1=1 limit ?")
	if err != nil {
		fmt.Printf("db prepare failed, err:%s\n", err)
		return
	}


	var article Article
	err = stat.QueryRow(3).Scan(&article.Aid, &article.Title)
	if err != nil {
		fmt.Printf("db query failed, err:%s\n", err)
		return
	}

	fmt.Printf("article: %+v", article)
}

func initDb() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:SunJianwei627@tcp(127.0.0.1:16033)/dc_article?charset=utf8&timeout=3000ms&parseTime=true&loc=Local")
	if err != nil {
		fmt.Printf("init db failed, err:%s\n", err)
	}
	return
}


