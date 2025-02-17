package constant

import (
	db "github.com/lkplanwise-api/db/sqlc"
	"github.com/lkplanwise-api/models"
)

func NewAccountResponse(account db.Account) models.AccountResponse {
	return models.AccountResponse{
		Id:        account.Id,
		UserName:  account.UserName,
		FullName:  account.FirstName.String + " " + account.LastName.String,
		Email:     account.Email,
		CreatedAt: account.CreatedAt.Time,
		CreatedBy: account.CreatedBy.String,
	}
}
