package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createMineRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Name string `json:"name"`
	Price string `json:"price"`
	Owner string `json:"owner"`
	Selling string `json:"selling"`
	Efficiency string `json:"efficiency"`
	Invetory string `json:"invetory"`
	Resources string `json:"resources"`
	UraniumCost string `json:"uraniumCost"`
	
}

func createMineHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createMineRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedName := req.Name
		
		parsedPrice := req.Price
		
		parsedOwner := req.Owner
		
		parsedSelling := req.Selling
		
		parsedEfficiency := req.Efficiency
		
		parsedInvetory := req.Invetory
		
		parsedResources := req.Resources
		
		parsedUraniumCost := req.UraniumCost
		

		msg := types.NewMsgCreateMine(
			creator,
			parsedName,
			parsedPrice,
			parsedOwner,
			parsedSelling,
			parsedEfficiency,
			parsedInvetory,
			parsedResources,
			parsedUraniumCost,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setMineRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	Creator string `json:"creator"`
	Name string `json:"name"`
	Price string `json:"price"`
	Owner string `json:"owner"`
	Selling string `json:"selling"`
	Efficiency string `json:"efficiency"`
	Invetory string `json:"invetory"`
	Resources string `json:"resources"`
	UraniumCost string `json:"uraniumCost"`
	
}

func setMineHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setMineRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedName := req.Name
		
		parsedPrice := req.Price
		
		parsedOwner := req.Owner
		
		parsedSelling := req.Selling
		
		parsedEfficiency := req.Efficiency
		
		parsedInvetory := req.Invetory
		
		parsedResources := req.Resources
		
		parsedUraniumCost := req.UraniumCost
		

		msg := types.NewMsgSetMine(
			creator,
			req.ID,
			parsedName,
			parsedPrice,
			parsedOwner,
			parsedSelling,
			parsedEfficiency,
			parsedInvetory,
			parsedResources,
			parsedUraniumCost,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteMineRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	ID 		string `json:"id"`
}

func deleteMineHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteMineRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgDeleteMine(req.ID, creator)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
