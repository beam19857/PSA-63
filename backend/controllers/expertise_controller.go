package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/beam19857/app/ent"
	"github.com/beam19857/app/ent/expertise"
	"github.com/gin-gonic/gin"
)

// ExpertiseController defines the struct for the expertise controller
type ExpertiseController struct {
	client *ent.Client
	router gin.IRouter
}
type Expertise struct {
	ExpertiseID int
	ExpertiseName string	
}

// CreateExpertise handles POST requests for adding expertise entities
// @Summary Create expertise
// @Description Create expertise
// @ID create-expertise
// @Accept   json
// @Produce  json
// @Param expertise body ent.Expertise true "Expertise entity"
// @Success 200 {object} ent.Expertise
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /expertises [post]
func (ctl *ExpertiseController) CreateExpertise(c *gin.Context) {
	obj := ent.Expertise{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "expertise binding failed",
		})
		return
	}

	u, err := ctl.client.Expertise.
		Create().
		SetExpertiseName(obj.ExpertiseName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetExpertise handles GET requests to retrieve a expertise entity
// @Summary Get a expertise entity by ID
// @Description get expertise by ID
// @ID get-expertise
// @Produce  json
// @Param id path int true "Expertise ID"
// @Success 200 {object} ent.Expertise
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /expertises/{id} [get]
func (ctl *ExpertiseController) GetExpertise(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Expertise.
		Query().
		Where(expertise.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListExpertise handles request to get a list of expertise entities
// @Summary List expertise entities
// @Description list expertise entities
// @ID list-expertise
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Expertise
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /expertises [get]
func (ctl *ExpertiseController) ListExpertise(c *gin.Context) {
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

	expertises, err := ctl.client.Expertise.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, expertises)
}

// DeleteExpertise handles DELETE requests to delete a expertise entity
// @Summary Delete a expertise entity by ID
// @Description get expertise by ID
// @ID delete-expertise
// @Produce  json
// @Param id path int true "Expertise ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /expertises/{id} [delete]
func (ctl *ExpertiseController) DeleteExpertise(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Expertise.
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

// UpdateExpertise handles PUT requests to update a expertise entity
// @Summary Update a expertise entity by ID
// @Description update expertise by ID
// @ID update-expertise
// @Accept   json
// @Produce  json
// @Param id path int true "Expertise ID"
// @Param expertise body ent.Expertise true "Expertise entity"
// @Success 200 {object} ent.Expertise
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /expertises/{id} [put]
func (ctl *ExpertiseController) UpdateExpertise(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Expertise{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "expertise binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Expertise.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewExpertiseController creates and registers handles for the expertise controller
func NewExpertiseController(router gin.IRouter, client *ent.Client) *ExpertiseController {
	uc := &ExpertiseController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitExpertiseController registers routes to the main engine
func (ctl *ExpertiseController) register() {
	expertises := ctl.router.Group("/expertises")

	expertises.GET("", ctl.ListExpertise)

	// CRUD
	expertises.POST("", ctl.CreateExpertise)
	expertises.GET(":id", ctl.GetExpertise)
	expertises.PUT(":id", ctl.UpdateExpertise)
	expertises.DELETE(":id", ctl.DeleteExpertise)
}
