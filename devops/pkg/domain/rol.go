package domain

import (
	"context"
	"time"
)

// CXROL struct is a row record of the CX_ROL table in the devops database
type Rol struct {
	ID               uint      `gorm:"primary_key;column:rol_id;autoIncrement;" json:"rol_id"`
	RolNombre        string    `gorm:"column:rol_nombre;type:VARCHAR;size:20;" json:"rol_nombre"`
	RolDescripcion   string    `gorm:"column:rol_descripcion;type:VARCHAR;size:100;" json:"rol_descripcion"`
	RolEsactivo      bool      `gorm:"column:rol_esactivo;type:BOOL;" json:"rol_esactivo"`
	RolEliminado     bool      `gorm:"column:rol_eliminado;type:BOOL;" json:"rol_eliminado"`
	RolFecharegistro time.Time `gorm:"column:rol_fecharegistro;autoCreateTime:true;" json:"rol_fecharegistro"`
}

// TableName sets the insert table name for this struct type
func (c *Rol) TableName() string {
	return "rol"
}

type RolRepository interface {
	FindAll(ctx context.Context) ([]Rol, error)
	FindByID(ctx context.Context, id uint) (Rol, error)
	Save(ctx context.Context, rol Rol) (Rol, error)
	Update(ctx context.Context, rol Rol) (Rol, error)
	Delete(ctx context.Context, rol Rol) error
	DeleteRolUser(ctx context.Context, rol Rol) error
}
