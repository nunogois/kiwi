package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleRoutes(router *gin.Engine) {
	store := router.Group("/store")
	{
		store.GET("/*key", storeGet)
		store.POST("/*key", storePost)
		store.PATCH("/*key", storePatch)
		store.DELETE("/*key", storeDelete)
	}
}

func storeGet(ctx *gin.Context) {
	key := ctx.Param("key")
	value := get(key)
	if value == nil {
		ctx.Status(http.StatusOK)
	} else {
		ctx.JSON(http.StatusOK, value)
	}
}

func storePost(ctx *gin.Context) {
	key := ctx.Param("key")
	var json interface{}
	if err := ctx.BindJSON(&json); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	value := set(key, json)
	if value == nil {
		ctx.Status(http.StatusOK)
	} else {
		ctx.JSON(http.StatusOK, value)
	}
}

func storePatch(ctx *gin.Context) {
	key := ctx.Param("key")
	var json map[string]interface{}
	if err := ctx.BindJSON(&json); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	value := patch(key, json)
	if value == nil {
		ctx.Status(http.StatusOK)
	} else {
		ctx.JSON(http.StatusOK, value)
	}
}

func storeDelete(ctx *gin.Context) {
	key := ctx.Param("key")
	del(key)
	ctx.Status(http.StatusOK)
}
