package event

import (
	"github.com/jmoiron/sqlx"
	"strings"
)

type GetDocumentQueryParam struct {
	EventId     *int64  `query:"event_id"`
	Name        *string `query:"name"`
	NameContain *string `query:"name_contain"`
}

func GetDocumentList(db *sqlx.DB, q GetDocumentQueryParam) ([]EventDocument, error) {
	// クエリを作成
	query := "SELECT * FROM documents WHERE"
	queryParams := []interface{}{}
	if q.EventId != nil {
		// イベントIDで絞り込み
		query += " event_id = ? AND"
		queryParams = append(queryParams, *q.EventId)
	}
	if q.Name != nil {
		// ドキュメント名の一致で絞り込み
		query += " name = ? AND"
		queryParams = append(queryParams, *q.Name)
	}
	if q.NameContain != nil {
		// ドキュメント名に文字列が含まれるかで絞り込み
		query += " name LIKE ?"
		queryParams = append(queryParams, "%"+*q.NameContain+"%")
	}
	// 不要な末尾の句を切り取り
	query = strings.TrimSuffix(query, " WHERE")
	query = strings.TrimSuffix(query, " AND")

	// `documents`テーブルからを取得し、変数`documents`に代入する
	var documents []EventDocument
	r, err := db.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	// 1行ずつ読込
	for r.Next() {
		// カラム読み込み用変数
		var ed EventDocument
		err = r.Scan(&ed.Id, &ed.EventId, &ed.Name, &ed.Url)
		if err != nil {
			return nil, err
		}
		documents = append(documents, ed)
	}

	return documents, nil
}
