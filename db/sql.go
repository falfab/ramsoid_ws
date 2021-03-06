package db

const CreateStatement = `
CREATE TABLE IF NOT EXISTS Keys(
	ID TEXT PRIMARY KEY,
	Key TEXT NOT NULL
)`

const Insert = `
	INSERT INTO Keys (ID, Key) VALUES (?, ?)
`

const SelectKey  = `
	SELECT ID, Key FROM Keys WHERE ID = ? AND Key = ?
`