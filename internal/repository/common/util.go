package common

import (
	"context"

	"github.com/imkarthi24/sf-backend/internal/utils"
	"github.com/imkarthi24/sf-backend/pkg/constants"
	"github.com/imkarthi24/sf-backend/pkg/db"
	"github.com/imkarthi24/sf-backend/pkg/errs"
	"gorm.io/gorm"
)

type CustomGormDB struct {
	txn db.DBTransactionManager
}

func ProvideCustomGormDB(txn db.DBTransactionManager) CustomGormDB {
	return CustomGormDB{txn: txn}
}

func (customDB *CustomGormDB) Delete(ctx *context.Context, model interface{}) *errs.XError {

	session := utils.GetSession(ctx)
	if session == nil {
		return errs.NewXError(errs.SMTP_ERROR, "Unable to get user session", nil)
	}

	res := customDB.txn.Txn(ctx, WithPrepareTransactionOption(ctx)).Model(model).
		Updates(map[string]interface{}{
			"is_active":     false,
			"updated_by_id": session.UserId,
		})

	if res.Error != nil {
		return errs.NewXError(errs.DATABASE, "Unable to delete", res.Error)
	}
	return nil
}

func (customDB *CustomGormDB) Update(ctx *context.Context, model interface{}) *errs.XError {

	// updateMap := util.PrepareEntityForUpdate(model)

	// val := util.ReadValueFromContext(ctx, constants.SESSION)
	// if val != nil {
	// 	updateMap["updated_by_id"] = val.(*models.Session).UserId
	// }

	res := customDB.txn.Txn(ctx, WithPrepareTransactionOption(ctx)).Model(model).Save(model)
	if res.Error != nil {
		return errs.NewXError(errs.DATABASE, "Unable to update entity", res.Error)
	}

	return nil
}

func (customDB *CustomGormDB) Create(ctx *context.Context, value interface{}) *errs.XError {

	res := customDB.txn.Txn(ctx, WithPrepareTransactionOption(ctx)).Create(value)
	if res.Error != nil {
		return errs.NewXError(errs.DATABASE, "Unable to create entity", res.Error)
	}

	return nil
}

func (customDB *CustomGormDB) BatchCreate(ctx *context.Context, value []interface{}) *errs.XError {

	res := customDB.txn.Txn(ctx, WithPrepareTransactionOption(ctx)).CreateInBatches(value, 10)
	if res.Error != nil {
		return errs.NewXError(errs.DATABASE, "Unable to batch create entity", res.Error)
	}

	return nil
}

// WithPrepareTransactionOption prepares the transaction with user and channel info
func WithPrepareTransactionOption(ctx *context.Context) db.TransactionOption {

	return func(db *gorm.DB) {
		session := utils.GetSession(ctx)
		if session == nil {
			return
		}

		db.Set(constants.USER_ID, session.UserId).
			Set(constants.CHANNEL_ID, session.ChannelId)
	}

}
