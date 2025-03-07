package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/dione/internal/reports/usecase"
)

type ReportsHandler struct {
	getSaleById         *usecase.GetSalesById
	getSalesByUser      *usecase.GetSalesByUserIDUsecase
}

func NewGetSaleByIdHandler(
	getSaleById *usecase.GetSalesById,
	getSalesByUser *usecase.GetSalesByUserIDUsecase,
) *ReportsHandler {
	return &ReportsHandler{
		getSaleById:         getSaleById,
		getSalesByUser:      getSalesByUser,
	}
}

func (h *ReportsHandler) GetSalesById(c *gin.Context) {
	saleID := c.Param("id")

	report, err := h.getSaleById.Execute(saleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

func (h *ReportsHandler) GetSalesByUserId(c *gin.Context) {
	userID := c.Param("id")
	sales, err := h.getSalesByUser.Execute(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if sales == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no sales found for the given user ID"})
		return
	}
	c.JSON(http.StatusOK, sales)
}
