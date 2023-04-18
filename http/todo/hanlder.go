package todo

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"

	"github.com/yotzapon/todo-service/config"
	"github.com/yotzapon/todo-service/http/helpers"
	"github.com/yotzapon/todo-service/internal/constants"
	"github.com/yotzapon/todo-service/internal/entity"
	"github.com/yotzapon/todo-service/internal/services"
	"github.com/yotzapon/todo-service/internal/xerrors"
	"github.com/yotzapon/todo-service/internal/xlogger"
)

var log = xlogger.Get()

type handler struct {
	service services.TodoServiceInterface
	config  *config.Config
}

func NewTodoHandler(serv services.TodoServiceInterface, cfg *config.Config) handler {
	return handler{service: serv, config: cfg}
}

// CreateTodo    	Create the todo
// @Summary      	This action involves creating a new Todo item in the system
// @Description  	The API should allow the user to input the details of the new Todo item, including its title, description, and any other relevant information. Upon successful creation, the API should return the newly created Todo item's unique ID.
// @Tags         	Todo
// @Accept       	json
// @Produce      	json
// @Param 			request body RequestTodoModel true "Create Todo body request"
// @Success      	201  {object} helpers.Response{data=todo.ResponseTodoModel}
// @Failure 		400,500 {object} helpers.ResponseError
// @Router       	/v1/todos [post]
func (h handler) CreateTodo(c *gin.Context) {
	log.Info("create todo handler")
	authHeader := c.GetHeader("Authorization")

	reqModel := &RequestTodoModel{}
	err := c.ShouldBindJSON(reqModel)
	if err != nil {
		log.Errorf("body request parsing error: %v", err.Error())
		c.JSON(http.StatusBadRequest, helpers.ErrorJSON(xerrors.InvalidBodyRequestCode, xerrors.InvalidBodyRequestStatus.Error()))

		return
	}

	inputEntity := reqModel.toTodoEntity()
	user := helpers.DecodeAccessToken(authHeader, h.config)

	inputEntity.UserID = user.UserID
	response, err := h.service.CreateTodo(inputEntity)
	if err != nil {
		log.Errorf("something went wrong: %v", err.Error())
		c.JSON(http.StatusInternalServerError, helpers.ErrorJSON(xerrors.InternalErrorCode, xerrors.InternalErrorStatus.Error()))

		return
	}
	log.Info("created todo")

	c.JSON(http.StatusCreated, helpers.SuccessJSON(http.StatusCreated, toTodoResponse(response)))
}

// GetTodo    		Get the todo
// @Summary      	This action involves retrieving a specific Todo item from the system.
// @Description  	The API should allow the user to input the unique ID of the Todo item and return its details, including its title, description, and any other relevant information.
// @Tags         	Todo
// @Accept       	json
// @Produce      	json
// @Param        	ids    			query     string  false  "Todo ID can input more than one by use , (comma) separate"
// @Param        	isComplete    	query     string  false  "Filter the result by 'true' or 'false'"
// @Param        	orderCreated    query     string  false  "Order the Todo by created dated the results by 'desc' or 'asc'"
// @Param        	orderUpdated    query     string  false  "Order the Todo by updated dated the results by 'desc' or 'asc'"
// @Param        	limit    		query     number  false  "Limit the result"
// @Success      	200  {object} helpers.Response{data=[]todo.ResponseTodoModel}
// @Failure 		400,500 {object} helpers.ResponseError
// @Router       	/v1/todos [get]
func (h handler) GetTodo(c *gin.Context) {
	log.Info("get todo handler")

	authHeader := c.GetHeader("Authorization")
	todoIDs := strings.TrimSpace(c.Query("ids"))
	isComplete := strings.TrimSpace(c.Query("isComplete"))
	orderCreated := strings.TrimSpace(c.Query("orderCreated")) // asc, desc
	orderUpdated := strings.TrimSpace(c.Query("orderUpdated")) // asc, desc
	limit := strings.TrimSpace(c.Query("limit"))               // asc, desc

	user := helpers.DecodeAccessToken(authHeader, h.config)

	ids := strings.Split(todoIDs, ",")
	for _, v := range ids {
		if v == "" {
			ids = nil
		}
	}
	filterMap := make(map[string]interface{})
	orderMap := make(map[string]interface{})
	var limitValue int64

	filterMap["user_id"] = user.UserID
	if ids != nil {
		filterMap["id"] = ids
	}

	if isComplete != "" {
		boolValue, err := strconv.ParseBool(isComplete)
		if err != nil {
			log.Errorf("parse error: %v", err.Error())
			c.JSON(http.StatusBadRequest, helpers.ErrorJSON(xerrors.InternalErrorCode, xerrors.InternalErrorStatus.Error()))

			return
		}
		filterMap["completed"] = boolValue
	}

	if orderCreated != "" && slices.Contains(constants.ValidOrder, strings.ToUpper(orderCreated)) {
		orderMap["created_at"] = strings.ToUpper(orderCreated)
	}

	if orderUpdated != "" && slices.Contains(constants.ValidOrder, strings.ToUpper(orderUpdated)) {
		orderMap["updated_at"] = strings.ToUpper(orderUpdated)
	}

	if limit != "" {
		limitParse, errParseInt := strconv.ParseInt(limit, 10, 64)
		if errParseInt != nil {
			log.Errorf("parse error: %v", errParseInt.Error())
			c.JSON(http.StatusBadRequest, helpers.ErrorJSON(xerrors.InternalErrorCode, xerrors.InternalErrorStatus.Error()))

			return
		}

		limitValue = limitParse
	}

	response, err := h.service.GetTodo(filterMap, orderMap, int(limitValue))
	if err != nil {
		log.Errorf("something went wrong: %v", err.Error())
		c.JSON(http.StatusInternalServerError, helpers.ErrorJSON(xerrors.InternalErrorCode, xerrors.InternalErrorStatus.Error()))

		return
	}
	log.Info("created todo")

	c.JSON(http.StatusOK, helpers.SuccessJSON(http.StatusOK, toTodoResponses(response)))
}

// UpdateTodo    	Update the todo
// @Summary      	This action involves updating an existing Todo item in the system.
// @Description  	The API should allow the user to input the unique ID of the Todo item and the updated details, such as its title, description, and any other relevant information.
// @Tags         	Todo
// @Accept       	json
// @Produce      	json
// @Param        	id    			query     string  false  "Todo ID needs to update"
// @Param 			request body RequestTodoModel true "Update Todo body request"
// @Success      	200  {object} helpers.Response{data=[]todo.ResponseTodoModel}
// @Failure 		400,500 {object} helpers.ResponseError
// @Router       	/v1/todos [put]
func (h handler) UpdateTodo(c *gin.Context) {
	log.Info("updated todo handler")
	authHeader := c.GetHeader("Authorization")
	todoID := c.Query("id")

	reqModel := &RequestTodoModel{}
	err := c.ShouldBindJSON(reqModel)
	if err != nil {
		log.Errorf("body request parsing error: %v", err.Error())
		c.JSON(http.StatusBadRequest, helpers.ErrorJSON(xerrors.InvalidBodyRequestCode, xerrors.InvalidBodyRequestStatus.Error()))

		return
	}

	inputEntity := reqModel.toTodoEntity()
	inputEntity.ID = todoID
	user := helpers.DecodeAccessToken(authHeader, h.config)

	inputEntity.UserID = user.UserID
	response, err := h.service.UpdateTodo(inputEntity)
	if err != nil {
		log.Errorf("something went wrong: %v", err.Error())
		c.JSON(http.StatusBadRequest, helpers.ErrorJSON(xerrors.InternalErrorCode, xerrors.InternalErrorStatus.Error()))

		return
	}
	log.Info("updated todo")

	c.JSON(http.StatusOK, helpers.SuccessJSON(http.StatusOK, toTodoResponse(response)))
}

// PatchTodo    	Patch the todo
// @Summary      	This action involves marking a specific Todo item as complete in the system.
// @Description  	The API should allow the user to input the unique ID of the Todo item and update its "completed" status to "true". Upon successful completion, the API should return the updated Todo item's unique ID.
// @Tags         	Todo
// @Accept       	json
// @Produce      	json
// @Param        	id    			query     string  false  "Todo ID need to mark as complete"
// @Success      	200  {object} helpers.Response{data=[]todo.ResponseTodoModel}
// @Failure 		400,500 {object} helpers.ResponseError
// @Router       	/v1/todos [patch]
func (h handler) PatchTodo(c *gin.Context) {
	log.Info("patch todo handler")
	authHeader := c.GetHeader("Authorization")
	todoID := c.Query("id")

	if strings.TrimSpace(todoID) == "" {
		log.Error("require param missing")
		c.JSON(http.StatusBadRequest, helpers.ErrorJSON(xerrors.InvalidBodyRequestCode, xerrors.InvalidBodyRequestStatus.Error()))

		return
	}

	user := helpers.DecodeAccessToken(authHeader, h.config)
	inputEntity := entity.TodoEntity{
		ID:          strings.TrimSpace(todoID),
		IsCompleted: true,
		UserID:      user.UserID,
	}

	response, err := h.service.MarkCompleteTodo(inputEntity)
	if err != nil {
		log.Errorf("something went wrong: %v", err.Error())
		c.JSON(http.StatusBadRequest, helpers.ErrorJSON(xerrors.InternalErrorCode, xerrors.InternalErrorStatus.Error()))

		return
	}
	log.Info("patched todo")

	c.JSON(http.StatusOK, helpers.SuccessJSON(http.StatusOK, toTodoResponse(response)))
}

// DeleteTodo    	Delete the todo
// @Summary      	This action involves deleting a specific Todo item from the system.
// @Description  	The API should allow the user to input the unique ID of the Todo item and delete it from the database.
// @Tags         	Todo
// @Accept       	json
// @Produce      	json
// @Param        	id    			query     string  false  "Todo ID needs to delete"
// @Success      	200  {object} helpers.Response{data=[]todo.ResponseTodoModel}
// @Failure 		400,500 {object} helpers.ResponseError
// @Router       	/v1/todos [delete]
func (h handler) DeleteTodo(c *gin.Context) {
	log.Info("delete todo handler")
	authHeader := c.GetHeader("Authorization")
	todoID := c.Query("id")

	if strings.TrimSpace(todoID) == "" {
		log.Error("require param missing")
		c.JSON(http.StatusBadRequest, helpers.ErrorJSON(xerrors.InvalidBodyRequestCode, xerrors.InvalidBodyRequestStatus.Error()))

		return
	}

	user := helpers.DecodeAccessToken(authHeader, h.config)
	inputEntity := entity.TodoEntity{
		ID:     strings.TrimSpace(todoID),
		UserID: user.UserID,
	}

	_, err := h.service.DeleteTodo(inputEntity)
	if err != nil {
		log.Errorf("something went wrong: %v", err.Error())
		c.JSON(http.StatusInternalServerError, helpers.ErrorJSON(xerrors.InternalErrorCode, xerrors.InternalErrorStatus.Error()))

		return
	}
	log.Info("deleted todo")

	c.JSON(http.StatusNoContent, helpers.SuccessJSON(http.StatusNoContent, "ok"))
}
