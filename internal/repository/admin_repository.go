package repository

import (
	"context"

	"github.com/imkarthi24/sf-backend/internal/repository/common"
	"github.com/imkarthi24/sf-backend/pkg/db"
	"github.com/imkarthi24/sf-backend/pkg/errs"
)

type AdminRepository interface {
	SwitchBranch(ctx *context.Context, params map[string]interface{}) *errs.XError
}

type adminRepository struct {
	txn      db.DBTransactionManager
	customDB common.CustomGormDB
}

func ProvideAdminRepository(txn db.DBTransactionManager, customDB common.CustomGormDB) AdminRepository {
	return &adminRepository{txn: txn, customDB: customDB}
}

func (ur *adminRepository) SwitchBranch(ctx *context.Context, params map[string]interface{}) *errs.XError {
	_, err := ur.txn.ExecuteStoredProc(ctx, "SwitchChannelForRecord", params)
	if err != nil {
		return errs.NewXError(errs.DATABASE, "Unable to execute func SwitchChannelForRecord", err)
	}
	return nil
}
