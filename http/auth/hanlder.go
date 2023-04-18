package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yotzapon/todo-service/http/helpers"
	"github.com/yotzapon/todo-service/internal/services"
	"github.com/yotzapon/todo-service/internal/xerrors"
	"github.com/yotzapon/todo-service/internal/xlogger"
)

var log = xlogger.Get()

type handler struct {
	Service services.AuthServiceInterface
}

func NewAuthHandler(serv services.AuthServiceInterface) handler {
	return handler{Service: serv}
}

// Login    		User login
// @Summary      	API is used to authenticate a user.
// @Description  	This API is used to authenticate a user and generate a unique token that can be used for accessing other APIs within the system.
// @Tags         	Auth
// @Accept       	json
// @Produce      	json
// @Param 			request body RequestAuthModel true "Login body request"
// @Success      	200  {object} ResponseAuthModel
// @Failure 		400,401 {object} helpers.ResponseError
// @Router       	/login [post]
func (h handler) Login(c *gin.Context) {
	log.Info("login handler")

	reqModel := &RequestAuthModel{}
	err := c.ShouldBindJSON(reqModel)
	if err != nil {
		log.Error("body request parsing error", err.Error())
		c.JSON(http.StatusBadRequest, helpers.ErrorJSON(xerrors.InvalidBodyRequestCode, xerrors.InvalidBodyRequestStatus.Error()))

		return
	}

	authEntity := reqModel.toAuthEntity()
	result, err := h.Service.Login(authEntity)
	if err != nil {
		log.Error("something went wrong", err.Error())
		c.JSON(http.StatusUnauthorized, helpers.ErrorJSON(xerrors.InvalidCredentialsCode, xerrors.InvalidCredentialsStatus.Error()))

		return
	}

	authEntity.UserID = result.UserID
	token, err := h.Service.GenerateJWT(authEntity)
	if err != nil {
		log.Error("something went wrong", err.Error())
		c.JSON(http.StatusUnauthorized, helpers.ErrorJSON(xerrors.InternalErrorCode, xerrors.InternalErrorStatus.Error()))

		return
	}

	c.JSON(http.StatusOK, &ResponseAuthModel{AccessToken: token})
}
