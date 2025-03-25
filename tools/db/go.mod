module github.com/meesooqa/cheque/tools/db

go 1.24.1

replace github.com/meesooqa/cheque/common => ../../common

replace github.com/meesooqa/cheque/db => ../../db

require github.com/meesooqa/cheque/db v0.0.0-00010101000000-000000000000

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/meesooqa/cheque/common v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/postgres v1.5.11 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
