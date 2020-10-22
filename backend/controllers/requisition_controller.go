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

// RequisitionController defines the struct for the user controller
type RequisitionController struct {
	client *ent.Client
	router gin.IRouter
}

type Requisition struct {
	User          int
	Registerstore int
	Drug          int
	Amount        int
	Added         string
}

// CreateRequisition handles POST requests for adding requisition entities
// @Summary Create requisition
// @Description Create requisition
// @ID create-requisition
// @Accept   json
// @Produce  json
// @Param requisition body ent.Requisition true "Requisition entity"
// @Success 200 {object} ent.Requisition
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requisitions [post]
func (ctl *RequisitionController) CreateRequisition(c *gin.Context) {
	obj := Requisition{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "requisition binding failed",
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
			"error": "registerstore not found",
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
	rq, err := ctl.client.Requisition.
		Create().
		SetUser(u).
		SetRegisterstore(rs).
		SetDrug(d).
		SetAmount(obj.Amount).
		SetAddedTime(time).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   rq,
	})
}

// GetRequisition handles GET requests to retrieve a requisition entity
// @Summary Get a requisition entity by ID
// @Description get requisition by ID
// @ID get-requisition
// @Produce  json
// @Param id path int true "Requisition ID"
// @Success 200 {object} ent.Requisition
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requisitions/{id} [get]
func (ctl *RequisitionController) GetRequisition(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	rq, err := ctl.client.Requisition.
		Query().
		Where(requisition.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, rq)
}

// ListRequisition handles request to get a list of requisition entities
// @Summary List requisition entities
// @Description list requisition entities
// @ID list-requisition
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Requisition
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requisitions [get]
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

	requisitions, err := ctl.client.Requisition.
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

	c.JSON(200, requisitions)
}

// DeleteRequisition handles DELETE requests to delete a requisition entity
// @Summary Delete a requisition entity by ID
// @Description get requisition by ID
// @ID delete-requisition
// @Produce  json
// @Param id path int true "Requisition ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requisitions/{id} [delete]
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

// UpdateRequisition handles PUT requests to update a requisition entity
// @Summary Update a requisition entity by ID
// @Description update requisition by ID
// @ID update-requisition
// @Accept   json
// @Produce  json
// @Param id path int true "Requisition ID"
// @Param requisition body ent.Requisition true "Requisition entity"
// @Success 200 {object} ent.Requisition
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /requisitions/{id} [put]
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
	uq, err := ctl.client.Requisition.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, uq)
}

// NewRequisitionController creates and registers handles for the requisition controller
func NewRequisitionController(router gin.IRouter, client *ent.Client) *RequisitionController {
	rc := &RequisitionController{
		client: client,
		router: router,
	}
	rc.register()
	return rc
}

// InitRequisitionController registers routes to the main engine
func (ctl *RequisitionController) register() {
	requisitions := ctl.router.Group("/requisitions")

	requisitions.GET("", ctl.ListRequisition)

	// CRUD
	requisitions.POST("", ctl.CreateRequisition)
	requisitions.GET(":id", ctl.GetRequisition)
	requisitions.DELETE(":id", ctl.DeleteRequisition)

}
