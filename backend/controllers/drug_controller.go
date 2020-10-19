package controllers

import (
	"context"
	"strconv"

	"github.com/Admin/app/ent"
	"github.com/Admin/app/ent/drug"
	"github.com/gin-gonic/gin"
)

// DrugController defines the struct for the Drug controller
type DrugController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDrug handles POST requests for adding drug entities
// @Summary Create drug
// @Description Create drug
// @ID create-drug
// @Accept   json
// @Produce  json
// @Param Drug body ent.Drug true "Drug entity"
// @Success 200 {object} ent.Drug
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugs [post]
func (ctl *DrugController) CreateDrug(c *gin.Context) {
	obj := ent.Drug{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "drug binding failed",
		})
		return
	}

	d, err := ctl.client.Drug.
		Create().
		SetName(obj.Name).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetDrug GET requests to retrieve a drug entity
// @Summary Get a drug entity by ID
// @Description get drug by ID
// @ID get-drug
// @Produce  json
// @Param id path int true "Drug ID"
// @Success 200 {object} ent.Drug
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugs/{id} [get]
func (ctl *DrugController) GetDrug(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	d, err := ctl.client.Drug.
		Query().
		Where(drug.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, d)
}

// ListDrug handles request to get a list of drug entities
// @Summary List drug entities
// @Description list drug entities
// @ID list-drug
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Drug
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugs [get]
func (ctl *DrugController) ListDrug(c *gin.Context) {
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

	drugs, err := ctl.client.Drug.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, drugs)
}

// NewDrugController creates and registers handles for the user controller
func NewDrugController(router gin.IRouter, client *ent.Client) *DrugController {
	dc := &DrugController{
		client: client,
		router: router,
	}
	dc.register()
	return dc
}

// InitDrugController registers routes to the main engine
func (ctl *DrugController) register() {
	drugs := ctl.router.Group("/drugs")

	drugs.GET("", ctl.ListDrug)
	drugs.POST("", ctl.CreateDrug)
	drugs.GET(":id", ctl.GetDrug)

}
