package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers mineservice-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/mineservice/mine", createMineHandler(cliCtx)).Methods("POST")
		//r.HandleFunc("/mineservice/mine", BuyMineHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/mineservice/mine", listMineHandler(cliCtx, "mineservice")).Methods("GET")
		r.HandleFunc("/mineservice/mine/{key}", getMineHandler(cliCtx, "mineservice")).Methods("GET")
		r.HandleFunc("/mineservice/mine", setMineHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/mineservice/mine", deleteMineHandler(cliCtx)).Methods("DELETE")

		
}
