package v1

import (
	"context"
	"exam/api-gateway/genproto/reyting"
	l "exam/api-gateway/pkg/logger"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary create reyting
// @Description this func creates ranking
// @Tags reyting
// @Accept json
// @Produce json
// @Param ranking body reyting.Ranking true "Ranking"
// @Success 200 {object} reyting.Empty
// @Router /v1/create-ranking [post]
func (h *handlerV1) CreateReyting(c *gin.Context) {
	var (
		r reyting.Ranking
	)
	err := c.ShouldBindJSON(&r)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	res, err := h.serviceManager.RankingService().CreateRanking(ctx, &reyting.Ranking{
		Name:        r.Name,
		Ranking:     r.Ranking,
		Description: r.Description,
		PostId:      r.PostId,
		CustomerId:  r.CustomerId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create ranking", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// @Summary get reyting
// @Description this func get rankings of posts
// @Tags reyting
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} reyting.Rankings
// @Router /v1/getranking-post/{id} [get]
func (h *handlerV1) GetRanking(c *gin.Context) {

	s_id := c.Param("id")
	id, err := strconv.ParseInt(s_id, 10, 64)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	res, err := h.serviceManager.RankingService().GetRankings(ctx, &reyting.Id{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get ranking", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}

// @Summary get reyting
// @Description this func get rankings of posts
// @Tags reyting
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} reyting.Rankings
// @Router /v1/getranking-customer/{id} [get]
func (h *handlerV1) GetRankingOfCustomer(c *gin.Context) {

	s_id := c.Param("id")
	id, err := strconv.ParseInt(s_id, 10, 64)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	res, err := h.serviceManager.RankingService().GetRankingsByCustomerId(ctx, &reyting.Id{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get ranking", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, res)
}
