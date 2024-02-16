package DB

type DBWriter interface {
	Write(cookbook CookBook) error
}
