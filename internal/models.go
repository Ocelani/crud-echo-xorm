package internal

type (
	// Macapa defines the User at the provided database.
	Macapa struct {
		ID        int    `xorm:"pk autoincr"`
		Name      string `xorm:"varchar(200) notnull"`
		Cellphone string `xorm:"varchar(20) notnull"`
	}

	// Varejao defines the User at the provided database.
	Varejao struct {
		ID        int    `xorm:"pk autoincr"`
		Name      string `xorm:"varchar(100) notnull"`
		Cellphone string `xorm:"varchar(13) notnull"`
	}
)
