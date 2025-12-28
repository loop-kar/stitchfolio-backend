package task

import (
	"context"

	"github.com/imkarthi24/sf-backend/pkg/constants"
	"github.com/imkarthi24/sf-backend/pkg/errs"
	"github.com/imkarthi24/sf-backend/pkg/util"
	"gorm.io/gorm"
)

func disposeTransaction(ctx *context.Context, err *errs.XError) {
	transaction := getTransactionFromContext(ctx)

	if transaction == nil {
		return
	}

	if err != nil {
		transaction.Rollback()
		return
	}

	transaction.Commit()

}

func getTransactionFromContext(ctx *context.Context) *gorm.DB {
	transaction := util.ReadValueFromContext(ctx, constants.TRANSACTION_KEY)
	if transaction == nil {
		return nil
	}
	return transaction.(*gorm.DB)
}
