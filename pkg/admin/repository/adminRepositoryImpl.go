package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakjack/isekai-shop-api-tutorial/databases"
	"github.com/supakjack/isekai-shop-api-tutorial/entities"
	_adminException "github.com/supakjack/isekai-shop-api-tutorial/pkg/admin/exception"
)

type adminRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewAdminRepositoryImpl(db databases.Database, logger echo.Logger) AdminRepository {
	return &adminRepositoryImpl{db, logger}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (*entities.Admin, error) {
	admin := new(entities.Admin)
	err := r.db.Connect().Create(adminEntity).Scan(admin).Error

	if err != nil {
		r.logger.Errorf("Creating admin failed: %s", err.Error())
		return nil, &_adminException.AdminCreating{AdminID: adminEntity.ID}
	}

	return admin, nil
}

func (r *adminRepositoryImpl) FindById(adminID string) (*entities.Admin, error) {
	admin := new(entities.Admin)

	err := r.db.Connect().First(admin, adminID).Error
	if err != nil {
		r.logger.Errorf("Failed to find admin by adminId : %s", err.Error())
		return nil, &_adminException.AdminNotFound{AdminID: adminID}
	}

	return admin, nil
}
