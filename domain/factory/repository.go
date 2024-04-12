package factory

import "github.com/jefflssantos/imersao/gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
