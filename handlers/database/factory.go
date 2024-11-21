package database

func GetDatabase(driver string, host string, port string, username string, password string, database string) Database {
	switch driver {
		case "postgres":
			return &Postgres{
				Host:     host,
				Port:     port,
				Username: username,
				Password: password,
				Database: database,
			}
		default:
			return nil
	}
}