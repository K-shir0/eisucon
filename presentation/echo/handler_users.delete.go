package echo

import (
	"net/http"
	"prc_hub_back/application/user"
	"prc_hub_back/domain/model/jwt"

	"github.com/labstack/echo/v4"
)

// (DELETE /users)
func (s *Server) DeleteUsers(ctx echo.Context) error {
	// Get jwt claim
	jcc, err := jwt.CheckProvided(ctx)
	if err != nil {
		return JSONMessage(ctx, http.StatusUnauthorized, err.Error())
	}

	// Delete user
	err = user.Delete(s.db, jcc.Id, jcc.Id)
	if err != nil {
		return JSONMessage(ctx, user.ErrToCode(err), err.Error())
	}

	return JSONMessage(ctx, http.StatusNoContent, "success")
}
