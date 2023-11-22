-- name: CreateUser :execresult
INSERT INTO users (id, created_at, updated_at, name, api_key) VALUES (UUID(), NOW(), ?, ?, LOWER(
    CONCAT(
      SUBSTRING(CONV(FLOOR(RAND() * 4294967295), 10, 16), 1, 16),
      SUBSTRING(CONV(FLOOR(RAND() * 4294967295), 10, 16), 1, 16),
      SUBSTRING(CONV(FLOOR(RAND() * 4294967295), 10, 16), 1, 16),
      SUBSTRING(CONV(FLOOR(RAND() * 4294967295), 10, 16), 1, 5)
    )
  ));

-- name: GetUserByAPIKey :one
select * from users where api_key = ?;