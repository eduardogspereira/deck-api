package deck

import (
	"errors"
	"strings"

	"github.com/eduardogspereira/deck-api/domains/card"
	"github.com/gin-gonic/gin"
)

type createOptions struct {
	Shuffle     bool   `form:"shuffle"`
	CardsString string `form:"cards"`
	Cards       []string
}

func bindCreateOptions(c *gin.Context) (createOptions, error) {
	var opt createOptions
	var err error

	if err = c.ShouldBindQuery(&opt); err != nil {
		err = errors.New("invalid request")
		return opt, err
	}

	if len(opt.CardsString) == 0 {
		return opt, err
	}

	opt.Cards = strings.Split(opt.CardsString, ",")

	for _, c := range opt.Cards {
		if _, err := card.FromCode(c); err != nil {
			err = errors.New("invalid card code provided")
			return opt, err
		}
	}

	return opt, err
}
