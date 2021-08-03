package config

import "gorm.io/gorm"

type appCtx struct {
	db *gorm.DB
}

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}
