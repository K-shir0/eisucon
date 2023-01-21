package user

import (
	"strings"
)

type GetUserListQueryParam struct {
	Name                *string `query:"name"`
	NameContain         *string `query:"name_contain"`
	PostEventAvailabled *bool   `json:"post_event_availabled"`
	Manage              *bool   `json:"manage"`
	Admin               *bool   `json:"admin"`
}

func GetList(q GetUserListQueryParam) ([]User, error) {
	// MySQLサーバーに接続
	db, err := OpenMysql()
	if err != nil {
		return nil, err
	}
	// return時にMySQLサーバーとの接続を閉じる
	defer db.Close()

	// クエリを作成
	query := "SELECT u.id, u.name, u.email, u.password, u.post_event_availabled, u.manage, u.admin, u.twitter_id, u.github_username, COUNT(s.target_user_id) AS star_count FROM users u LEFT JOIN user_stars s ON u.id = s.target_user_id  GROUP BY u.id HAVING"
	queryParams := []interface{}{}
	if q.PostEventAvailabled != nil {
		// 権限で絞り込み
		query += " u.post_event_availabled = ? AND"
		queryParams = append(queryParams, *q.PostEventAvailabled)
	}
	if q.Manage != nil {
		// 権限で絞り込み
		query += " u.manage = ? AND"
		queryParams = append(queryParams, *q.Manage)
	}
	if q.Admin != nil {
		// 権限で絞り込み
		query += " u.admin = ? AND"
		queryParams = append(queryParams, *q.Admin)
	}
	if q.Name != nil {
		// ドキュメント名の一致で絞り込み
		query += " u.name = ? AND"
		queryParams = append(queryParams, *q.Name)
	}
	if q.NameContain != nil {
		// ドキュメント名に文字列が含まれるかで絞り込み
		query += " u.name LIKE ?"
		queryParams = append(queryParams, "%"+*q.NameContain+"%")
	}

	//panic(query)

	// 不要な末尾の句を切り取り
	query = strings.TrimSuffix(query, " HAVING")
	query = strings.TrimSuffix(query, " AND")

	// `users`テーブルからを取得
	r, err := db.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	// 取得したテーブルを１行ずつ処理
	// 配列`users`に代入する
	var users []User
	for r.Next() {
		// 一時変数に割り当て
		var (
			id                  int64
			name                string
			email               string
			password            string
			postEventAvailabled bool
			manage              bool
			admin               bool
			twitterId           *string
			githubUsername      *string
			count               uint64
		)
		err = r.Scan(
			&id, &name, &email, &password, &postEventAvailabled,
			&manage, &admin, &twitterId, &githubUsername, &count,
		)
		if err != nil {
			return nil, err
		}

		// 配列に追加
		users = append(
			users,
			User{
				Id:                  id,
				Name:                name,
				Email:               email,
				Password:            password,
				StarCount:           count,
				PostEventAvailabled: postEventAvailabled,
				Manage:              manage,
				Admin:               admin,
				TwitterId:           twitterId,
				GithubUsername:      githubUsername,
			},
		)
	}

	return users, nil
}
