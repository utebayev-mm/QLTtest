package utils

func (r Repo) DeleteCategory(ID int) error {
	sqlstatement := `DELETE FROM categories WHERE id=$1`
	_, err := r.db.Exec(sqlstatement, ID)
	if err != nil {
		panic(err)
	}

	return nil
}
