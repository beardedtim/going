package data

type Author struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
}

func FindAllAuthors() ([]Author, error) {
	var authors []Author
	var err error

	err = db.Model(&authors).Select()

	return authors, err
}

type AuthorCreate struct {
	Name   string
	Avatar string
}

func CreateNewAuthor(author AuthorCreate) (*Author, error) {
	var err error
	data := &Author{
		Name:   author.Name,
		Avatar: author.Avatar,
	}

	_, err = db.Model(data).Insert()

	return data, err
}
