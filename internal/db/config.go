package db

type pgsettings struct {
	PgHost string
	PgPort string
	PgUser string
	PgPass string
	PgBase string
}

func init() {
	Cfg := pgsettings{
		"127.0.0.1",
		"5432",
		"postgres",
		"1234",
		"tg_bot",
	}
}
