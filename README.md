# prc_hub backend

## Get started

```console
$ go run .
  --jwt-issuer <jwt_issuer>
  --jwt-secret <jwt_secret>
  --admin-email <admin_user_email>
  --admin-password <admin_user_password>
  --allow-origin <allow_origin>
```

# 学内 ISUCON に参加した

@tiTakung が学内で実施した ISUCON に参加した。その時の write-up

alp を使ってレスポンスまでの時間を計測するベースに改善していった

### nginx を使うようにポートを 1323 -> 1322 に変更

[alp 用 config を追加 · K-shir0/eisucon@df6c562 (github.com)](https://github.com/K-shir0/eisucon/commit/df6c562f763463b98c27176a3948570ae37a9dbe)

### gzip を使うようにした

返却されている json は gzip を使って圧縮して送信していなかった。

echo の実行前に１行追加するだけで対応完了

```go
e.Use(middleware.Gzip())
```

### users GET の n+1 問題を解決

1600 -> 1800

200点くらい伸びた

[update: users n+1 解消 · K-shir0/eisucon@f2393c5 (github.com)](https://github.com/K-shir0/eisucon/commit/f2393c599dee29bec563af6646bec58d0d5bfd81)

ただ自分が sql 忘れてたのもあり HAVING してしまったのちに WHERE に直す

### events GET の n+1 問題を解決 

[update: event n+1 問題を解消 · K-shir0/eisucon@cd68c23 (github.com)](https://github.com/K-shir0/eisucon/commit/cd68c2309253dc4bdf648c2de37a633db4e11a4b)
1800 -> 4000 くらい

N * N * N + 1 問題が発生していてここを改善することで一番速くすることができた。ライブラリは sqlc を使った。（go に依存しないライブラリなので神）

ちなみに書いた SQL は以下
`sqlc.arg(set_location)` は引数に対応するために使っている CASE 文があるのは sqlc が動的な WHERE に対応していないためである、なので CASE で分岐している

ただのちに分岐させることや HAVING を使っているので WHERE に直したりしている

```sql
-- name: ListEvents :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end,
       documents.id                     AS document_id,
       documents.event_id               AS document_event_id,
       documents.name                   AS document_name,
       documents.url,
       users.id,
       users.name,
       users.email,
       users.password,
       users.post_event_availabled,
       users.manage,
       users.admin,
       users.twitter_id,
       users.github_username,
       COUNT(user_stars.target_user_id) as star_count
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
         JOIN documents ON events.id = documents.event_id
         JOIN users ON events.user_id = users.id
         LEFT JOIN user_stars ON users.id = user_stars.target_user_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end, documents.id,
         documents.event_id, documents.name, documents.url, users.id, users.name, users.email, users.password,
         users.post_event_availabled, users.manage, users.admin, users.twitter_id, users.github_username
HAVING events.name LIKE CASE
                            WHEN sqlc.arg(set_event_name) != '%'
                                THEN sqlc.arg(set_event_name)
                            ELSE events.name
    END
   AND events.location LIKE CASE
                                WHEN sqlc.arg(set_location) != '%'
                                    THEN sqlc.arg(set_location)
                                ELSE events.location
    END
   AND events.published = CASE
                                WHEN sqlc.arg(not_set_published) = false
                                    THEN sqlc.arg(set_published)
                                ELSE events.published
    END;
```

### 毎回 db.Open しないように修正

特に点数にはあまり影響がなかった

[fix: 毎回 db に接続しに行かないように修正 · K-shir0/eisucon@9b8da9a (github.com)](https://github.com/K-shir0/eisucon/commit/9b8da9a7661f63906224fda51c25a6ef3fadee91)
[fix: user 関連を毎回 db に接続しに行かないように修正 · K-shir0/eisucon@c6d50c8 (github.com)](https://github.com/K-shir0/eisucon/commit/c6d50c8c85b044cd0d905e8d629051a6bd3cbbe6)

sql.open を echo まえで実行し、それを参照させる形にした


### sql を分岐させ高速化

4000 -> 6000 

documents がある場合と users がある場合で分岐させた

[update: sql を分岐させ高速化 · K-shir0/eisucon@aea4244 (github.com)](https://github.com/K-shir0/eisucon/commit/aea42448336e0f080afd8e9c39797220e6fe094a)
書いた sql 4つは以下 sqlc のライブラリでなんかしてくれたらいいのに（無理）

```
-- name: ListEventsWithUserAndDocuments :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end,
       documents.id                     AS document_id,
       documents.event_id               AS document_event_id,
       documents.name                   AS document_name,
       documents.url,
       users.id,
       users.name,
       users.email,
       users.password,
       users.post_event_availabled,
       users.manage,
       users.admin,
       users.twitter_id,
       users.github_username,
       COUNT(user_stars.target_user_id) as star_count
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
         JOIN documents ON events.id = documents.event_id
         JOIN users ON events.user_id = users.id
         LEFT JOIN user_stars ON users.id = user_stars.target_user_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end, documents.id,
         documents.event_id, documents.name, documents.url, users.id, users.name, users.email, users.password,
         users.post_event_availabled, users.manage, users.admin, users.twitter_id, users.github_username
HAVING events.name LIKE CASE
                            WHEN sqlc.arg(set_event_name) != '%'
                                THEN sqlc.arg(set_event_name)
                            ELSE events.name
    END
   AND events.location LIKE CASE
                                WHEN sqlc.arg(set_location) != '%'
                                    THEN sqlc.arg(set_location)
                                ELSE events.location
    END
   AND events.published = CASE
                              WHEN sqlc.arg(not_set_published) = false
                                  THEN sqlc.arg(set_published)
                              ELSE events.published
    END;

-- name: ListEventsWithUser :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end,
       users.id,
       users.name,
       users.email,
       users.password,
       users.post_event_availabled,
       users.manage,
       users.admin,
       users.twitter_id,
       users.github_username,
       COUNT(user_stars.target_user_id) as star_count
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
         JOIN users ON events.user_id = users.id
         LEFT JOIN user_stars ON users.id = user_stars.target_user_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end, users.id, users.name,
         users.email, users.password, users.post_event_availabled, users.manage, users.admin, users.twitter_id,
         users.github_username
HAVING events.name LIKE CASE
                            WHEN sqlc.arg(set_event_name) != '%'
                                THEN sqlc.arg(set_event_name)
                            ELSE events.name
    END
   AND events.location LIKE CASE
                                WHEN sqlc.arg(set_location) != '%'
                                    THEN sqlc.arg(set_location)
                                ELSE events.location
    END
   AND events.published = CASE
                              WHEN sqlc.arg(not_set_published) = false
                                  THEN sqlc.arg(set_published)
                              ELSE events.published
    END;

-- name: ListEventsWithDocuments :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end,
       documents.id       AS document_id,
       documents.event_id AS document_event_id,
       documents.name     AS document_name,
       documents.url
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
         JOIN documents ON events.id = documents.event_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end, documents.id,
         documents.event_id, documents.name, documents.url
HAVING events.name LIKE CASE
                            WHEN sqlc.arg(set_event_name) != '%'
                                THEN sqlc.arg(set_event_name)
                            ELSE events.name
    END
   AND events.location LIKE CASE
                                WHEN sqlc.arg(set_location) != '%'
                                    THEN sqlc.arg(set_location)
                                ELSE events.location
    END
   AND events.published = CASE
                              WHEN sqlc.arg(not_set_published) = false
                                  THEN sqlc.arg(set_published)
                              ELSE events.published
    END;

-- name: ListEvents :many
SELECT events.id,
       events.name,
       events.description,
       events.location,
       events.published,
       events.completed,
       events.user_id,
       event_datetimes.event_id,
       event_datetimes.start,
       event_datetimes.end
FROM events
         JOIN event_datetimes ON events.id = event_datetimes.event_id
GROUP BY events.id, events.name, events.description, events.location, events.published, events.completed,
         events.user_id, event_datetimes.event_id, event_datetimes.start, event_datetimes.end
HAVING events.name LIKE CASE
                            WHEN sqlc.arg(set_event_name) != '%'
                                THEN sqlc.arg(set_event_name)
                            ELSE events.name
    END
   AND events.location LIKE CASE
                                WHEN sqlc.arg(set_location) != '%'
                                    THEN sqlc.arg(set_location)
                                ELSE events.location
    END
   AND events.published = CASE
                              WHEN sqlc.arg(not_set_published) = false
                                  THEN sqlc.arg(set_published)
                              ELSE events.published
    END;

```

### sqlite3 に移行（失敗）

onmemory の sqlite3 に移行、さすがに同時に書き込まれるリクエストをテストされていたので無理だった結局 revert

[update: 全てのデータを memory に載せた · K-shir0/eisucon@3a2aed0 (github.com)](https://github.com/K-shir0/eisucon/commit/3a2aed048755bed147ca80359990f97cb0f3474c)

### mysql のパフォーマンスをチューニング

さっきの events の sql を叩くと `Usage Tempolary` が多く発生していたので以下のように設定、ベンチマークのブレで効果があったかは微妙

```
innodb_buffer_pool_size= 1G
tmp_table_size= 1024M
max_heap_table_size= 128M
```

### mysql のデータをすべて onmemory に乗せた

あんまりベンチマークに差は出なかった悲しい

`CREATE TABLE `の後に `ENGINE=MEMORY` を付けるとオンメモリに乗せることができる（再起動するとデータが消えるので注意）

## 参加してみて 

後半は実は Aurora を使ってみたりとかしてみたけどローカルのdbのスピード（ネットワークの速度も含む）には勝てなかった。sqlite3 が動いてたら相当速くなってたと思う（無理だったけど）
