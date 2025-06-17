package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	domain "mg-inc-devops/pkg/domain"
	"mg-inc-devops/pkg/responses"
	services "mg-inc-devops/pkg/services/devops/interface"
)

type DevopsHandler struct {
	devopsUseCase services.DevopsUseCase
}

type ResponseDevops struct {
	ID            int32     `json:"id" copier:"must"`
	Message       string    `json:"message" copier:"must"`
	To            string    `json:"to" copier:"must"`
	From          string    `json:"from" copier:"must"`
	TimeToLifeSec int32     `json:"timeToLifeSec" copier:"must"`
	DateCreate    time.Time `json:"date_create" copier:"must"`
	DateDeleted   time.Time `json:"date_deleted" copier:"must"`
}

func NewDevopsHandler(usecase services.DevopsUseCase) *DevopsHandler {
	return &DevopsHandler{
		devopsUseCase: usecase,
	}
}

// FindAll godoc
//	@summary		Get all devops
//	@description	Get all devops
//	@tags			devops
//	@security		ApiKeyAuth
//	@Param			Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@id				FindAll
//	@produce		json
//	@Router			/devops [get]
//	@response		200	{object}	[]ResponseDevops	"OK"
func (cr *DevopsHandler) FindAll(c *gin.Context) {
	devops, err := cr.devopsUseCase.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	} else {
		response := []ResponseDevops{}
		copier.Copy(&response, &devops)

		responses.SuccessJSON(c, 200, response)
	}
}

//	@tags		devops
//	@security	ApiKeyAuth
//	@Param		Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Param		id				path	int		true	" ID"
//	@Router		/devops/{id} [get]
func (cr *DevopsHandler) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	devops, err := cr.devopsUseCase.FindByID(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := ResponseDevops{}
		copier.Copy(&response, &devops)

		responses.SuccessJSON(c, 200, response)
	}
}

//	@Accept		json
//	@Produce	json
//	@Param		devopsInput	body	domain.Devops	true	"Devops Data"
//	@tags		devops
//	@Success	200	{object}	handler.ResponseDevops
//	@security	ApiKeyAuth
//	@Param		Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Router		/devops/save [post]
func (cr *DevopsHandler) Save(c *gin.Context) {

	var devops domain.Devops
	//var vars map[string]interface{}

	//if exists set values
	if err := c.BindJSON(&devops); err != nil {
		responses.ErrorJSON(c, 400, err)
		c.Abort()
		return
	}

	devops, err := cr.devopsUseCase.Save(c.Request.Context(), devops)

	if err != nil {
		responses.ErrorJSON(c, 200, err)
		c.Abort()
	} else {

		responses.SuccessJSON(c, 200, devops)
	}

}

//	@Success	200	{string}	string	"ok"
//	@tags		devops
//	@Router		/devops/saveall [post]
//	@security	ApiKeyAuth
//	@Param		Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Param		Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)

func (cr *DevopsHandler) SaveAll(c *gin.Context) {
	var devops []domain.Devops

	//if exists set values
	if err := c.BindJSON(&devops); err != nil {
		responses.ErrorJSON(c, 400, err)
		c.Abort()
		return
	}

	_result, err := cr.devopsUseCase.SaveAll(c.Request.Context(), devops)

	if err != nil {
		responses.ErrorJSON(c, 200, err)
		c.Abort()
	} else {
		// response := ResponseDevops{}
		// copier.Copy(&response, &devops)

		c.JSON(http.StatusOK, gin.H{"result": _result})
	}

}

//	@Accept		json
//	@Produce	json
//	@Param		id			path		int				true	" ID"
//	@Param		devopsInput	body		domain.Devops	true	"Devops Data"
//	@Success	200			{object}	ResponseDevops
//	@tags		devops
//	@security	ApiKeyAuth
//	@Param		Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Router		/devops/update/{id} [put]
func (cr *DevopsHandler) Update(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	olddevops, err := cr.devopsUseCase.FindByID(c.Request.Context(), uint(id))
	var newdevops domain.Devops
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		//if exists set values
		if err := c.BindJSON(&newdevops); err != nil {
			responses.ErrorJSON(c, 400, err)
			c.Abort()
			return
		}

		//Set values update
		olddevops.Message = newdevops.Message
		olddevops.To = newdevops.To
		olddevops.From = newdevops.From
		olddevops.TimeToLifeSec = newdevops.TimeToLifeSec

		// olddevops.ProEliminado = newdevops.ProEliminado

		devops, err := cr.devopsUseCase.Save(c.Request.Context(), olddevops)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			response := ResponseDevops{}
			copier.Copy(&response, &devops)

			responses.SuccessJSON(c, 200, response)
		}
	}
}

//	@Success	200	{string}	string	"Devops is deleted successfully"
//	@tags		devops
//	@Param		id	path	int	true	" ID"
//	@Router		/devops/{id} [delete]
//	@security	ApiKeyAuth
//	@Param		Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
func (cr *DevopsHandler) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	ctx := c.Request.Context()
	devops, err := cr.devopsUseCase.FindByID(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if devops == (domain.Devops{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Devops is not booking yet",
		})
		return
	}

	cr.devopsUseCase.Delete(ctx, devops)

	responses.SuccessJSON(c, 200, "Devops is deleted successfully")
}
