package migrations

import (
	"github.com/lillez7/HRMS/internals/database"
	"github.com/lillez7/HRMS/internals/types"
)

func AutoMigrate() {
	database.DB.AutoMigrate(&types.Employee{})
}
