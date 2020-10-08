package controllers

import (
	"context"
	"strconv"

	"github.com/Admin/app/ent"
	"github.com/Admin/app/ent/registerstore"
	"github.com/gin-gonic/gin"
)

// RegisterStoreController defines the struct for the registerstore controller
type RegisterStoreController struct {
	client *ent.Client
	router gin.IRouter
}
type RegisterStore struct {
	Name string
}

// CreateRegisterStore handles POST requests for adding registerstore entities
// @Summary Create registerstore
// @Description Create registerstore
// @ID create-registerstore
// @Accept   json
// @Produce  json
// @Param Drug body ent.RegisterStore true "RegisterStore entity"
// @Success 200 {object} ent.RegisterStore
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registerstores [post]
func (ctl *DrugController) CreateRegisterStore(c *gin.Context) {
	obj := RegisterStore{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "drug binding failed",
		})
		return
	}

	rs, err := ctl.client.RegisterStore.
		Create().
		SetName(obj.Name).
		// SetValue(obj.Value).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, rs)
}

// GetRegisterStore GET requests to retrieve a registerstore entity
// @Summary Get a registerStore entity by ID
// @Description get registerStore by ID
// @ID get-registerStore
// @Produce  json
// @Param id path int true "registerstore ID"
// @Success 200 {object} ent.RegisterStore
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registerstores/{id} [get]
func (ctl *RegisterStoreController) GetRegisterStore(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	rs, err := ctl.client.RegisterStore.
		Query().
		Where(registerstore.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, rs)
}

// ListRegisterStore handles request to get a list of registerstore entities
// @Summary List registerstore entities
// @Description list registerstore entities
// @ID list-registerstore
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.RegisterStore
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registerstores [get]
func (ctl *RegisterStoreController) ListRegisterStore(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	registerstores, err := ctl.client.RegisterStore.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, registerstores)
}

// NewRegisterStoreController creates and registers handles for the user controller
func NewRegisterStoreController(router gin.IRouter, client *ent.Client) *RegisterStoreController {
	uc := &RegisterStoreController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitRegisterStoreController registers routes to the main engine
func (ctl *RegisterStoreController) register() {
	registerstores := ctl.router.Group("/registerstores")

	registerstores.GET("", ctl.ListRegisterStore)

	// CRUD
	registerstores.GET(":id", ctl.GetRegisterStore)

}
