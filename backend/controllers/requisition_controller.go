package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Admin/app/ent"
	"github.com/Admin/app/ent/drug"
	"github.com/Admin/app/ent/registerstore"
	"github.com/Admin/app/ent/requisition"
	"github.com/Admin/app/ent/user"
	"github.com/gin-gonic/gin"
)

// RequisitionController defines the struct for the Requisition controller
type RequisitionController struct {
	client *ent.Client
	router gin.IRouter
}

type Requisition struct {
	User          int
	Registerstore int
	Drug          int
	Added         string
}

// CreateRequisition handles POST requests for adding Requisition entities
// @Summary Create Requisition
// @Description Create Requisition
// @ID create-Requisition
// @Accept   json
// @Produce  json
// @Param Requisition body ent.Requisition true "Requisition entity"
// @Success 200 {object} ent.Requisition
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requisitions [post]
func (ctl *RequisitionController) CreateRequisition(c *gin.Context) {
	obj := Requisition{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Requisition binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	rs, err := ctl.client.RegisterStore.
		Query().
		Where(registerstore.IDEQ(int(obj.Registerstore))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "RegisterStore not found",
		})
		return
	}

	d, err := ctl.client.Drug.
		Query().
		Where(drug.IDEQ(int(obj.Drug))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "drug not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Added)
	pv, err := ctl.client.Requisition.
		Create().
		SetAddedTime(time).
		SetUser(u).
		SetRegisterstore(rs).
		SetDrug(d).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pv)
}

// GetRequisition handles GET requests to retrieve a Requisition entity
// @Summary Get a Requisition entity by ID
// @Description get Requisition by ID
// @ID get-Requisition
// @Produce  json
// @Param id path int true "Requisition ID"
// @Success 200 {object} ent.Requisition
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /RequisitionsRequisitionRequisition/{id} [get]
func (ctl *RequisitionController) GetRequisition(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pv, err := ctl.client.Requisition.
		Query().
		Where(requisition.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pv)
}

// ListRequisition handles request to get a list of RequisitionRequisition entities
// @Summary List RequisitionRequisition entities
// @Description list RequisitionRequisition entities
// @ID list-RequisitionRequisition
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Requisition
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Requisition [get]
func (ctl *RequisitionController) ListRequisition(c *gin.Context) {
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

	Requisition, err := ctl.client.Requisition.
		Query().
		WithUser().
		WithRegisterstore().
		WithDrug().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Requisition)
}

// DeleteRequisition handles DELETE requests to delete a Requisition entity
// @Summary Delete a Requisition entity by ID
// @Description get Requisition by ID
// @ID delete-Requisition
// @Produce  json
// @Param id path int true "Requisition ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /RequisitionsRequisition/{id} [delete]
func (ctl *RequisitionController) DeleteRequisition(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Requisition.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateRequisition handles PUT requests to update a Requisition entity
// @Summary Update a Requisition entity by ID
// @Description update Requisition by ID
// @ID update-Requisition
// @Accept   json
// @Produce  json
// @Param id path int true "Requisition ID"
// @Param Requisition body ent.Requisition true "Requisition entity"
// @Success 200 {object} ent.Requisition
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Requisitions/{id} [put]
func (ctl *RequisitionController) UpdateRequisition(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Requisition{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Requisition binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Requisition.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewRequisitionController creates and registers handles for the Requisition controller
func NewRequisitionController(router gin.IRouter, client *ent.Client) *RequisitionController {
	uc := &RequisitionController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitRequisitionController registers routes to the main engine
func (ctl *RequisitionController) register() {
	requisitions := ctl.router.Group("/requisitions")

	requisitions.GET("", ctl.ListRequisition)

	// CRUD
	requisitions.POST("", ctl.CreateRequisition)
	requisitions.GET(":id", ctl.GetRequisition)
	requisitions.PUT(":id", ctl.UpdateRequisition)
	requisitions.DELETE(":id", ctl.DeleteRequisition)
}
