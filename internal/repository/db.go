package repository

import (
	"context"

	"github.com/imkarthi24/sf-backend/internal/repository/common"
	"github.com/imkarthi24/sf-backend/pkg/db"
	"gorm.io/gorm"
)

// DB returns the gorm DB instance from the transaction manager with prepared transaction options
func DB(ctx *context.Context, transactionManager db.DBTransactionManager) *gorm.DB {
	return transactionManager.Txn(ctx, common.WithPrepareTransactionOption(ctx))
}
