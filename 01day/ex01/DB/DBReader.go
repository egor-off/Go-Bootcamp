package DB

type DBReader interface {
	Read() (*CookBook, error)
}

