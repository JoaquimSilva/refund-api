package database

import (
	"refund-api/models"
	"refund-api/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "host=localhost user=root password=root dbname=root port=5430 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	utils.CatchErrorDB(err)
}

func FindByAgencyId(agency string, refunds []models.Refund) *gorm.DB {
	return DB.Where("agency_id = ? ", agency).Find(&refunds)
}

func FindByAgencyIdAndPerPeriod(agency string, start string, end string, refunds []models.Refund) *gorm.DB {
	return DB.Where("agency_id = ? AND date_requested BETWEEN ? AND ? ", agency, start, end).Find(&refunds)
}

func FindByTicket(number string, refund models.Refund) *gorm.DB {
	return DB.Where("ticket_number = ? ", number).Find(&refund)
}

func FindByTicket64(number int64, refund models.Refund) *gorm.DB {
	return DB.Where("ticket_number = ? ", number).Find(&refund)
}
