package filters

type User struct {
	Email string
	ID    string
}

func (f *User) WhereClause() (string, string) {
	if f.Email != "" {
		return "email=?", f.Email
	}

	if f.ID != "" {
		return "id=?", f.ID
	}

	return "", ""
}
