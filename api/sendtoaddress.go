package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cmingxu/wallet-keeper/keeper"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (api *ApiServer) SendToAddress(c *gin.Context) {
	value, _ := c.Get(KEEPER_KEY) // sure about the presence of this value
	keeper := value.(keeper.Keeper)

	address, addrFound := c.GetQuery("address")
	amountArg, amountFound := c.GetQuery("amount")
	if !addrFound || !amountFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "both address and amount should present",
		})
		return
	}

	amount, err := strconv.ParseFloat(amountArg, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})
		return
	}

	err = keeper.SendToAddress(address, amount)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": address,
		})
	}
}
