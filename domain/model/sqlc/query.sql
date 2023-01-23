-- name: GetUser :one
SELECT u.id,
       u.name,
       u.email,
       u.password,
       u.post_event_availabled,
       u.manage,
       u.admin,
       u.twitter_id,
       u.github_username,
       COUNT(s.target_user_id) AS star_count
FROM users u
         LEFT JOIN user_stars s ON u.id = s.target_user_id
GROUP BY u.id
HAVING u.email LIKE CASE
                        WHEN sqlc.arg(set_email) != '%'
                            THEN sqlc.arg(set_email)
                        ELSE u.email
    END;

-- name: GetEventWithUserAndDocuments :many
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
HAVING events.id = sqlc.arg(set_event_id);

-- name: GetEventWithUser :many
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
HAVING events.id = sqlc.arg(set_event_id);

-- name: GetEventWithDocuments :many
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
HAVING events.id = sqlc.arg(set_event_id);

-- name: GetEvent :many
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
HAVING events.id = sqlc.arg(set_event_id);

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
