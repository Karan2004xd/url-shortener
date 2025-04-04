package utils

import "database/sql"

type RowScanner interface {
	Scan(rows *sql.Rows) error
	New() RowScanner
}

func GetArrayFromRows[T RowScanner](rows *sql.Rows) ([]T, error) {
	var items []T

	defer rows.Close()

	for rows.Next() {
		var temp T 
		item := temp.New()

		if typedItem, ok := item.(T); ok {
			err := typedItem.Scan(rows)
			if err != nil {
				return nil, err
			}
			items = append(items, typedItem)
		}
	}
	return items, nil
}
