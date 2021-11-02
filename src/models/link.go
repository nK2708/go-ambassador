package models

type Link struct {
	Model
	Code     string    `json:"code"`
	UserId   uint      `json:"user_id"`
	User     User      `json:"user" gorm:"foreignKey:UserId;preload:false;"`
	Products []Product `json:"products" gorm:"many2many:link_products;preload:false;"`
	Orders   []Order   `json:"orders,omitempty" gorm:"-"`
}
