package heroes

import "github.com/gin-gonic/gin"

type HeroHandler struct{}

func (h HeroHandler) GetAllHero(repo HeroRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		heroes, err := repo.GetAllHero(ctx)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
		c.JSON(200, heroes)
	}
}
