package domain

import (
	"context"
	"time"
)

type User struct {
	ID              uint      `gorm:"primary_key;column:us_id;autoIncrement;" json:"us_id"`
	UsUsuario       string    `gorm:"column:us_usuario;type:VARCHAR;size:50;" json:"us_usuario"`
	UsNombre        string    `gorm:"column:us_nombre;type:VARCHAR;size:50;" json:"us_nombre"`
	UsApellido      string    `gorm:"column:us_apellido;type:VARCHAR;size:50;" json:"us_apellido"`
	UsCorreo        string    `gorm:"column:us_correo;type:VARCHAR;size:100;" json:"us_correo"`
	UsClave         string    `gorm:"column:us_clave;type:VARCHAR;size:50;" json:"us_clave"`
	UsEsactivo      bool      `gorm:"column:us_esactivo;type:BOOL;" json:"us_esactivo"`
	UsEliminado     bool      `gorm:"column:us_eliminado;type:BOOL;" json:"us_eliminado"`
	UsFecharegistro time.Time `gorm:"column:us_fecharegistro;autoCreateTime:true;" json:"us_fecharegistro"`
	Roles           []Rol     `gorm:"many2many:cx_rol_usuarios;association_jointable_foreignkey:rol_id" json:"us_roles"`
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Tokens struct {
	Token Token `json:"token"`
}
type Token struct {
	Expires_in    int64  `json:"expires_in"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id uint) (User, error)
	FindByIDWithRole(ctx context.Context, id uint) (User, error)
	Save(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, user User) error
	FindByUserName(ctx context.Context, username string) (User, error)
}

// TableName sets the insert table name for this struct type
func (c *User) TableName() string {
	return "usuario"
}
