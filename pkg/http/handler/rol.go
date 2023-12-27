package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	domain "mg-inc-devops/pkg/domain"
	"mg-inc-devops/pkg/responses"
	services "mg-inc-devops/pkg/services/rol/interface"
)

type RolHandler struct {
	rolUseCase services.RolUseCase
}

type ResponseRol struct {
	ID               uint      `json:"rol_id" copier:"must"`
	RolNombre        string    `json:"rol_nombre" copier:"must"`
	RolDescripcion   string    `json:"rol_descripcion" copier:"must"`
	RolEsactivo      bool      `json:"rol_esactivo" copier:"must"`
	RolEliminado     bool      `json:"rol_eliminado" copier:"must"`
	RolFecharegistro time.Time `json:"rol_fecharegistro" copier:"must"`
}

func NewRolHandler(usecase services.RolUseCase) *RolHandler {
	return &RolHandler{
		rolUseCase: usecase,
	}
}

// FindAll godoc
//	@summary		Get all rol
//	@description	Get all rol
//	@tags			rol
//	@security		ApiKeyAuth
//	@Param			Authorization	header	string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@id				FindAllRol
//	@produce		json
//	@Router			/rol [get]
//	@response		200	{object}	[]ResponseRol	"OK"
func (cr *RolHandler) FindAll(c *gin.Context) {
	rol, err := cr.rolUseCase.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	} else {
		response := []ResponseRol{}
		copier.Copy(&response, &rol)

		responses.SuccessJSON(c, 200, response)
	}
}

func (cr *RolHandler) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	rol, err := cr.rolUseCase.FindByID(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := ResponseRol{}
		copier.Copy(&response, &rol)

		responses.SuccessJSON(c, 200, response)
	}
}

func (cr *RolHandler) Save(c *gin.Context) {
	var rol domain.Rol

	//if exists set values
	if err := c.BindJSON(&rol); err != nil {
		responses.ErrorJSON(c, 400, err)
		c.Abort()
		return
	}

	rol, err := cr.rolUseCase.Save(c.Request.Context(), rol)

	if err != nil {
		responses.ErrorJSON(c, 200, err)
		c.Abort()
	} else {
		response := ResponseRol{}
		copier.Copy(&response, &rol)

		responses.SuccessJSON(c, 200, response)
	}

}

func (cr *RolHandler) Update(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	oldrol, err := cr.rolUseCase.FindByID(c.Request.Context(), uint(id))
	var newrol domain.Rol
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		//if exists set values
		if err := c.BindJSON(&newrol); err != nil {
			responses.ErrorJSON(c, 400, err)
			c.Abort()
			return
		}

		//Set values update
		oldrol.RolNombre = newrol.RolNombre
		oldrol.RolDescripcion = newrol.RolDescripcion
		oldrol.RolEsactivo = newrol.RolEsactivo
		// oldrol.RolEliminado = newrol.RolEliminado

		rol, err := cr.rolUseCase.Save(c.Request.Context(), oldrol)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			response := ResponseRol{}
			copier.Copy(&response, &rol)

			responses.SuccessJSON(c, 200, response)
		}
	}
}

func (cr *RolHandler) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	ctx := c.Request.Context()
	rol, err := cr.rolUseCase.FindByID(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if rol == (domain.Rol{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Rol is not booking yet",
		})
		return
	}

	cr.rolUseCase.Delete(ctx, rol)

	c.JSON(http.StatusOK, gin.H{"message": "Rol is deleted successfully"})
}

func (cr *RolHandler) DeleteRolUser(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	ctx := c.Request.Context()
	rol, err := cr.rolUseCase.FindByID(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if rol == (domain.Rol{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Rol is not booking yet",
		})
		return
	}

	cr.rolUseCase.DeleteRolUser(ctx, rol)

	responses.SuccessJSON(c, 200, "Rol is deleted successfully")
}
