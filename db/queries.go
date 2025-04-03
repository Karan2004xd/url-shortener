package db

const InsertNewUser = `
	INSERT INTO users (email, password)
	VALUES (?, ?)
`

const SelectUserByEmailAndPassword = `
	SELECT * FROM users
	WHERE email = ?
`

const SelectAllUsers = `
	SELECT * FROM users
`
