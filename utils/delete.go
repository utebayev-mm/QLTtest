package utils

func (r Repo) Delete(ID int) error {
	sqlstatement := `DELETE FROM transactions WHERE id=$1`
	_, err := r.db.Exec(sqlstatement, ID)
	if err != nil {
		panic(err)
	}

	return nil
}
