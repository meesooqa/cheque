package repositories

import (
	"github.com/meesooqa/cheque/common/common_db"
	"github.com/meesooqa/cheque/common/models"
)

type OperatorRepository struct {
	common_db.BaseRepository[models.Operator]
}

func NewOperatorRepository() *OperatorRepository {
	return &OperatorRepository{common_db.BaseRepository[models.Operator]{}}
}
