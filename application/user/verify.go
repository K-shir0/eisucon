package user

import (
	"github.com/jmoiron/sqlx"
	"prc_hub_back/domain/model/jwt"
)

func Verify(db *sqlx.DB, email string, password string) (token string, verify bool, err error) {
	u, err := GetByEmail(db, email)
	if err != nil {
		return "", false, err
	}
	verify, err = u.Verify(password)
	if err != nil {
		return "", false, err
	}
	token, err = jwt.GenerateToken(jwt.GenerateTokenParam{Id: u.Id, Email: u.Email, Admin: u.Admin})
	if err != nil {
		return "", false, err
	}
	return token, verify, nil
}
