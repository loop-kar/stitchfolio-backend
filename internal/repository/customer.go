package repository

import (
	"context"

	"github.com/imkarthi24/sf-backend/internal/entities"
	"github.com/imkarthi24/sf-backend/internal/repository/common"
	"github.com/imkarthi24/sf-backend/pkg/db"
	"github.com/imkarthi24/sf-backend/pkg/errs"
)

type CustomerRepository interface {
	Create(*context.Context, *entities.Customer) *errs.XError
	Update(*context.Context, *entities.Customer) *errs.XError
	Get(*context.Context, uint) (*entities.Customer, *errs.XError)
	GetAll(*context.Context, string) ([]entities.Customer, *errs.XError)
	Delete(*context.Context, uint) *errs.XError
	CustomerAutoComplete(*context.Context, string) ([]entities.Customer, *errs.XError)
}

type customerRepository struct {
	txn      db.DBTransactionManager
	customDB common.CustomGormDB
}

func ProvideCustomerRepository(txn db.DBTransactionManager, customDB common.CustomGormDB) CustomerRepository {
	return &customerRepository{txn: txn, customDB: customDB}
}

func (cr *customerRepository) Create(ctx *context.Context, customer *entities.Customer) *errs.XError {
	res := cr.txn.Txn(ctx).Create(&customer)
	if res.Error != nil {
		return errs.NewXError(errs.DATABASE, "Unable to save customer", res.Error)
	}
	return nil
}

func (cr *customerRepository) Update(ctx *context.Context, customer *entities.Customer) *errs.XError {
	return cr.customDB.Update(ctx, *customer)
}

func (cr *customerRepository) Get(ctx *context.Context, id uint) (*entities.Customer, *errs.XError) {
	customer := entities.Customer{}
	res := cr.txn.Txn(ctx).Find(&customer, id)
	if res.Error != nil {
		return nil, errs.NewXError(errs.DATABASE, "Unable to find customer", res.Error)
	}
	return &customer, nil
}

func (cr *customerRepository) GetAll(ctx *context.Context, search string) ([]entities.Customer, *errs.XError) {
	customers := new([]entities.Customer)
	res := cr.txn.Txn(ctx).Find(customers)
	if res.Error != nil {
		return nil, errs.NewXError(errs.DATABASE, "Unable to find customers", res.Error)
	}
	return *customers, nil
}

func (cr *customerRepository) Delete(ctx *context.Context, id uint) *errs.XError {
	customer := &entities.Customer{Model: &entities.Model{ID: id, IsActive: false}}
	err := cr.customDB.Delete(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}

func (cr *customerRepository) CustomerAutoComplete(ctx *context.Context, search string) ([]entities.Customer, *errs.XError) {
	customers := new([]entities.Customer)
	res := cr.txn.Txn(ctx).Where("name LIKE ?", "%"+search+"%").Find(customers)
	if res.Error != nil {
		return nil, errs.NewXError(errs.DATABASE, "Unable to find customers", res.Error)
	}
	return *customers, nil
}
