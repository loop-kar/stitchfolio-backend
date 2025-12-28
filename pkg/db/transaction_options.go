package db

type TransactionOptions struct {
	Model      interface{}
	IsPaginate bool
}

func DefaultTransactionOptions(model interface{}) *TransactionOptions {
	return &TransactionOptions{
		Model:      model,
		IsPaginate: true,
	}
}
