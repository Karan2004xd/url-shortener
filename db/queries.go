package db

const (
	InsertNewUser = `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`

	SelectUserByEmailAndPassword = `
	SELECT * FROM users
	WHERE email = ?
	`

	SelectAllUsers = `
	SELECT * FROM users
	`

	InsertNewUrl = `
	INSERT INTO urls (id, short_url, long_url, user_id)
	VALUES (?, ?, ?, ?)
	`

	SelectUrlByShortUrl = `
	SELECT * FROM urls WHERE short_url = ?
	`

	UpdateUrl = `
	UPDATE urls
	SET long_url = ?
	WHERE id = ? AND short_url = ?
	`

	DeleteUrlById = `
	DELETE FROM urls WHERE id = ?
	`

	InsertNewCustomUrl = `
	INSERT INTO custom_urls (name, url_id)
	VALUES (?, ?)
	`

	UpdateCustomUrl = `
	UPDATE custom_urls
	SET name = ?
	WHERE url_id = ?
	`

	SelectAllCustomUrlsByUserId = `
	SELECT
		custom_urls.name,
		urls.id,
		urls.short_url,
		urls.long_url,
		urls.created_on,
		urls.user_id
	FROM custom_urls
	LEFT JOIN urls ON custom_urls.url_id = urls.id
	WHERE urls.user_id = ?
	`
)
