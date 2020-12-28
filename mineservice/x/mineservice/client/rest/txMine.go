package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/leonvanderhaeghen/mineservice/x/mineservice/types"
	//"github.com/syndtr/goleveldb/leveldb/util"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createMineRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Owner string `json:"owner"`
	PlayerID string `json:"owner"`
	Name string `json:"name"`
	Price string `json:"price"`
	Selling string `json:"selling"`
	Efficiency string `json:"efficiency"`
	Resources []string `json:"resources"`
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
		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedName := req.Name
		parsedPlayerID := req.PlayerID
		
		parsedPrice,_ := sdk.ParseCoins(req.Price)
				
		parsedSelling := false
		
		parsedEfficiency,_ := strconv.Atoi(req.Efficiency)
				
		parsedResources := req.Resources
		
		parsedUraniumCost,_ := strconv.Atoi(req.UraniumCost)
		

		msg := types.NewMsgCreateMine(
			owner,
			parsedPlayerID,
			parsedName,
			parsedPrice,
			parsedSelling,
			parsedEfficiency,
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
	Name string `json:"name"`
	Price string `json:"price"`
	Owner string `json:"owner"`
	Selling string `json:"selling"`
	Efficiency string `json:"efficiency"`
	Invetory string `json:"invetory"`
	Resources []string `json:"resources"`
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
		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedName := req.Name
		
		parsedPrice,_ := sdk.ParseCoins(req.Price)
				
		parsedSelling := false
		
		parsedEfficiency,_ := strconv.Atoi(req.Efficiency)
		
		parsedInvetory := req.Invetory
		
		parsedResources := req.Resources
		
		parsedUraniumCost,_ := strconv.Atoi(req.UraniumCost)
		

		msg := types.NewMsgSetMine(
			owner,
			req.ID,
			parsedName,
			parsedPrice,
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
	Owner string `json:"owner"`
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
		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgDeleteMine(req.ID, owner)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type sellMineRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	Owner string `json:"owner"`
	Price string `json:"price"`
}

func sellMineHandler(cliCtx context.CLIContext) http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request) {
		var req sellMineRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedPrice,_ := sdk.ParseCoins(req.Price)
		parsedID := req.ID
		msg := types.NewMsgSellMine(owner,parsedID,parsedPrice)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w,http.StatusBadRequest,err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})

	}
}

type buyMineRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	BuyerID 		string `json:"buyerid"`
	SellerID 		string `json:"sellerid"`
	Owner string `json:"owner"`
	Price string `json:"price"`
}

func buyMineHandler(cliCtx context.CLIContext) http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request) {
		var req buyMineRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedPrice,_ := sdk.ParseCoins(req.Price)
		parsedID := req.ID
		parsedBuyerID := req.BuyerID
		parsedSellerID := req.SellerID
		msg := types.NewMsgBuyMine(parsedID,parsedBuyerID,parsedSellerID,owner,parsedPrice)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w,http.StatusBadRequest,err.Error())
			return
		}
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})

	}
}