package repository

import (
	"context"

	"github.com/imkarthi24/sf-backend/internal/entities"
	"github.com/imkarthi24/sf-backend/internal/repository/common"
	"github.com/imkarthi24/sf-backend/pkg/db"
	"github.com/imkarthi24/sf-backend/pkg/errs"
)

type OrderItemRepository interface {
	Create(*context.Context, *entities.OrderItem) *errs.XError
	Update(*context.Context, *entities.OrderItem) *errs.XError
	Get(*context.Context, uint) (*entities.OrderItem, *errs.XError)
	GetAll(*context.Context, string) ([]entities.OrderItem, *errs.XError)
	Delete(*context.Context, uint) *errs.XError
}

type orderItemRepository struct {
	txn      db.DBTransactionManager
	customDB common.CustomGormDB
}

func ProvideOrderItemRepository(txn db.DBTransactionManager, customDB common.CustomGormDB) OrderItemRepository {
	return &orderItemRepository{txn: txn, customDB: customDB}
}

func (oir *orderItemRepository) Create(ctx *context.Context, orderItem *entities.OrderItem) *errs.XError {
	res := oir.txn.Txn(ctx).Create(&orderItem)
	if res.Error != nil {
		return errs.NewXError(errs.DATABASE, "Unable to save order item", res.Error)
	}
	return nil
}

func (oir *orderItemRepository) Update(ctx *context.Context, orderItem *entities.OrderItem) *errs.XError {
	return oir.customDB.Update(ctx, *orderItem)
}

func (oir *orderItemRepository) Get(ctx *context.Context, id uint) (*entities.OrderItem, *errs.XError) {
	orderItem := entities.OrderItem{}
	res := oir.txn.Txn(ctx).Preload("Order").Find(&orderItem, id)
	if res.Error != nil {
		return nil, errs.NewXError(errs.DATABASE, "Unable to find order item", res.Error)
	}
	return &orderItem, nil
}

func (oir *orderItemRepository) GetAll(ctx *context.Context, search string) ([]entities.OrderItem, *errs.XError) {
	var orderItems []entities.OrderItem
	res := oir.txn.Txn(ctx).Model(&entities.OrderItem{}).Preload("Order").Find(&orderItems)
	if res.Error != nil {
		return nil, errs.NewXError(errs.DATABASE, "Unable to find order items", res.Error)
	}
	return orderItems, nil
}

func (oir *orderItemRepository) Delete(ctx *context.Context, id uint) *errs.XError {
	orderItem := &entities.OrderItem{Model: &entities.Model{ID: id, IsActive: false}}
	err := oir.customDB.Delete(ctx, orderItem)
	if err != nil {
		return err
	}
	return nil
}
