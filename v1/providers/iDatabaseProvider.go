package providers

type DBConn interface {
	Connect() (interface{}, error)
}
