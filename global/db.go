package global

import (
	"database/sql"

	"github.com/impact-eintr/eorm"
)

var DB *sql.DB
var DBClient *eorm.Client
