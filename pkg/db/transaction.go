package db

import (
	"context"

	"github.com/imkarthi24/sf-backend/internal/utils"
	"github.com/imkarthi24/sf-backend/pkg/constants"
	"github.com/imkarthi24/sf-backend/pkg/util"
	"gorm.io/gorm"
)

type DBTransactionManager interface {
	Txn(ctx *context.Context, opts ...*TransactionOptions) *gorm.DB
	Commit(ctx *context.Context)
	Rollback(ctx *context.Context)
	ExecuteStoredProc(ctx *context.Context, name string, params map[string]interface{}) ([]ResultSet, error)
}

type txnManager struct {
	*StoredProcExecutor
	db *gorm.DB
}

func ProvideDBTransactionManager(db *gorm.DB) DBTransactionManager {
	return &txnManager{
		db:                 db,
		StoredProcExecutor: &StoredProcExecutor{db: db},
	}
}

func (txn *txnManager) Txn(ctx *context.Context, opts ...*TransactionOptions) *gorm.DB {

	if util.ReadValueFromContext(ctx, constants.TRANSACTION_KEY) == nil {
		return txn.createTransaction(ctx)
	}

	transactionObj := util.ReadValueFromContext(ctx, constants.TRANSACTION_KEY)
	return transactionObj.(*gorm.DB)

}

func (txn *txnManager) Commit(ctx *context.Context) {
	transaction := txn.Txn(ctx)
	transaction.Commit()
}

func (txn *txnManager) Rollback(ctx *context.Context) {
	transaction := txn.Txn(ctx)
	transaction.Rollback()
}

func (txn *txnManager) ExecuteStoredProc(ctx *context.Context, name string, params map[string]interface{}) ([]ResultSet, error) {
	return txn.StoredProcExecutor.CallStoredProcedure(ctx, name, params)
}

func (txn *txnManager) createTransaction(ctx *context.Context) *gorm.DB {

	transaction := txn.prepareTransaction(ctx)
	newCtx := util.NewContextWithValue(ctx, constants.TRANSACTION_KEY, transaction)
	*ctx = newCtx

	return transaction

}

func (txn *txnManager) prepareTransaction(ctx *context.Context) *gorm.DB {

	session := utils.GetSession(ctx)
	if session == nil {
		return txn.db.Begin()
	}

	//Set SessionData to transaction to fetch in BeforeCreate/AfterCreate hooks
	transaction := txn.db.Set(constants.USER_ID, session.UserId).
		Set(constants.CHANNEL_ID, session.ChannelId)

	return transaction.Begin()

}
