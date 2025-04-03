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

	InsertNewCustomUrl = `
	INSERT INTO custom_urls (name, url_id)
	VALUES (?, ?)
	`
)
