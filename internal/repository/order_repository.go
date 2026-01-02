package repository

import (
	"context"

	"github.com/imkarthi24/sf-backend/internal/entities"
	"github.com/imkarthi24/sf-backend/internal/repository/common"
	"github.com/imkarthi24/sf-backend/pkg/db"
	"github.com/imkarthi24/sf-backend/pkg/errs"
)

type OrderRepository interface {
	Create(*context.Context, *entities.Order) *errs.XError
	Update(*context.Context, *entities.Order) *errs.XError
	Get(*context.Context, uint) (*entities.Order, *errs.XError)
	GetAll(*context.Context, string) ([]entities.Order, *errs.XError)
	Delete(*context.Context, uint) *errs.XError
}

type orderRepository struct {
	txn      db.DBTransactionManager
	customDB common.CustomGormDB
}

func ProvideOrderRepository(txn db.DBTransactionManager, customDB common.CustomGormDB) OrderRepository {
	return &orderRepository{txn: txn, customDB: customDB}
}

func (or *orderRepository) Create(ctx *context.Context, order *entities.Order) *errs.XError {
	res := or.txn.Txn(ctx).Create(&order)
	if res.Error != nil {
		return errs.NewXError(errs.DATABASE, "Unable to save order", res.Error)
	}
	return nil
}

func (or *orderRepository) Update(ctx *context.Context, order *entities.Order) *errs.XError {
	return or.customDB.Update(ctx, *order)
}

func (or *orderRepository) Get(ctx *context.Context, id uint) (*entities.Order, *errs.XError) {
	order := entities.Order{}
	res := or.txn.Txn(ctx).Preload("Customer").Preload("OrderItems").Find(&order, id)
	if res.Error != nil {
		return nil, errs.NewXError(errs.DATABASE, "Unable to find order", res.Error)
	}
	return &order, nil
}

func (or *orderRepository) GetAll(ctx *context.Context, search string) ([]entities.Order, *errs.XError) {
	var orders []entities.Order
	res := or.txn.Txn(ctx).Model(&entities.Order{}).Preload("Customer").Preload("OrderItems").Find(&orders)
	if res.Error != nil {
		return nil, errs.NewXError(errs.DATABASE, "Unable to find orders", res.Error)
	}
	return orders, nil
}

func (or *orderRepository) Delete(ctx *context.Context, id uint) *errs.XError {
	order := &entities.Order{Model: &entities.Model{ID: id, IsActive: false}}
	err := or.customDB.Delete(ctx, order)
	if err != nil {
		return err
	}
	return nil
}
