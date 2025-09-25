package controller

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDataHander godoc
// @Summary      get data by url
// @Description  get data by input url as result of html parsing
// @Tags         data
// @Accept       json
// @Produce      json
// @Param        url   query      string  true  "URL"
// @Success      200  {string}  string "data"
// @Failure      400  {string}  string "no url is provided"
// @Failure      500  {string}  string "internal server error"
// @Router 		/data [get]
func (s *Server) GetDataHander(c *gin.Context) { // в usecase <- c.gin.Context() (создается на уровне входа, т.е)
	ctx := c.Request.Context()
	url := c.Query("url")
	fmt.Println(url)
	if url == "" {
		s.l.Error("GetDataHander c.Param",
			slog.Any("error", "no url is provided"),
			slog.Int("status", http.StatusBadRequest))
		c.JSON(http.StatusBadRequest, gin.H{"error": "no url is provided"})
		return
	}
	data, err := s.u.GetParsingDataV1(ctx, url)
	if err != nil {
		s.l.Error("GetDataHander s.u.GetParsingDataV1",
			slog.Any("error", err.Error()),
			slog.Int("status", http.StatusInternalServerError))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error of getting parsing data"})
		return
	}
	s.l.Info("GetDataHander",
		slog.String("message", "get data by url successful"),
		slog.Int("status", http.StatusOK))
	c.String(http.StatusOK, data)
}
