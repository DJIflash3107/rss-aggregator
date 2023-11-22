-- +goose Up

alter table users add column api_key varchar(69) unique not null default(LOWER(
    CONCAT(
      SUBSTRING(CONV(FLOOR(RAND() * 4294967295), 10, 16), 1, 16),
      SUBSTRING(CONV(FLOOR(RAND() * 4294967295), 10, 16), 1, 16),
      SUBSTRING(CONV(FLOOR(RAND() * 4294967295), 10, 16), 1, 16),
      SUBSTRING(CONV(FLOOR(RAND() * 4294967295), 10, 16), 1, 5)
    )
  ));

-- +goose Down

alter table users drop column api_key;