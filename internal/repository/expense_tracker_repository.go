package repository

import (
	"context"

	"github.com/imkarthi24/sf-backend/internal/entities"
	"github.com/imkarthi24/sf-backend/internal/repository/common"
	"github.com/imkarthi24/sf-backend/internal/repository/scopes"
	"github.com/imkarthi24/sf-backend/pkg/db"
	"github.com/imkarthi24/sf-backend/pkg/errs"
)

type ExpenseTrackerRepository interface {
	Create(*context.Context, *entities.ExpenseTracker) *errs.XError
	Update(*context.Context, *entities.ExpenseTracker) *errs.XError
	Get(*context.Context, uint) (*entities.ExpenseTracker, *errs.XError)
	GetAll(*context.Context, string) ([]entities.ExpenseTracker, *errs.XError)
	Delete(*context.Context, uint) *errs.XError
}

type expenseTrackerRepository struct {
	txn      db.DBTransactionManager
	customDB common.CustomGormDB
}

func ProvideExpenseTrackerRepository(txn db.DBTransactionManager, customDB common.CustomGormDB) ExpenseTrackerRepository {
	return &expenseTrackerRepository{txn: txn, customDB: customDB}
}

func (etr *expenseTrackerRepository) Create(ctx *context.Context, expenseTracker *entities.ExpenseTracker) *errs.XError {
	res := etr.txn.Txn(ctx).Create(&expenseTracker)
	if res.Error != nil {
		return errs.NewXError(errs.DATABASE, "Unable to save expense tracker", res.Error)
	}
	return nil
}

func (etr *expenseTrackerRepository) Update(ctx *context.Context, expenseTracker *entities.ExpenseTracker) *errs.XError {
	return etr.customDB.Update(ctx, *expenseTracker)
}

func (etr *expenseTrackerRepository) Get(ctx *context.Context, id uint) (*entities.ExpenseTracker, *errs.XError) {
	expenseTracker := entities.ExpenseTracker{}
	res := etr.txn.Txn(ctx).
		Scopes(scopes.Channel(), scopes.IsActive()).
		Find(&expenseTracker, id)
	if res.Error != nil {
		return nil, errs.NewXError(errs.DATABASE, "Unable to find expense tracker", res.Error)
	}
	return &expenseTracker, nil
}

func (etr *expenseTrackerRepository) GetAll(ctx *context.Context, search string) ([]entities.ExpenseTracker, *errs.XError) {
	var expenseTrackers []entities.ExpenseTracker

	res := etr.txn.Txn(ctx).
		Scopes(scopes.Channel(), scopes.IsActive()).
		Scopes(db.Paginate(ctx)).
		Find(&expenseTrackers)
	if res.Error != nil {
		return nil, errs.NewXError(errs.DATABASE, "Unable to find expense trackers", res.Error)
	}
	return expenseTrackers, nil
}

func (etr *expenseTrackerRepository) Delete(ctx *context.Context, id uint) *errs.XError {
	expenseTracker := &entities.ExpenseTracker{Model: &entities.Model{ID: id, IsActive: false}}
	err := etr.customDB.Delete(ctx, expenseTracker)
	if err != nil {
		return err
	}
	return nil
}
