package main

import (
	"spiropoulos94/FRONTIDA_STO_SPITI_backend/utils"
)

func main() {
	utils.SetupDatabase()

	SetupRouter()
}
