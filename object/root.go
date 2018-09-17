package object

import (
	"database/sql"
)

type (
	RootObject struct {
		ID   uint
		Path string
	}
	RootObjects []*RootObject
)

func FetchRootObjects(db *sql.DB) (RootObjects, error) {
	rows, err := db.Query("select id, path from root_object")
	if err != nil {
		return nil, err
	}

	dest := make(RootObjects, 0)
	for rows.Next() {
		var obj RootObject
		if err := rows.Scan(&obj.ID, &obj.Path); err != nil {
			return nil, err
		}
		dest = append(dest, &obj)
	}
	return dest, nil
}
