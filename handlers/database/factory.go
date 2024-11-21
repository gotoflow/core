package database

func GetDatabase(driver string, host string, port int, username string, password string, database string) Database {
	switch driver {
		case "postgres":
			return &Postgres{
				Host:     host,
				Port:     string(port),
				Username: username,
				Password: password,
				Database: database,
			}
		default:
			return nil
	}
}