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

type createResourceRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Owner string `json:"owner"`
	Name string `json:"name"`
	Amount string `json:"amount"`
	MineID string `json:"mineID"`

}

func createResourceHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createResourceRequest
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
		
		
		parsedAmount,_ :=  strconv.Atoi(req.Amount)

		parsedMineID := req.MineID


		msg := types.NewMsgCreateResource(
			owner,
			parsedName,
			parsedMineID,
			parsedAmount,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setResourceRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	Creator string `json:"creator"`
	Name string `json:"name"`
	Amount string `json:"amount"`
	Price string `json:"price"`
	Selling string `json:"selling"`
	MinPrice string `json:"minPrice"`
	
}

func setResourceHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setResourceRequest
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
		parsedAmount,_ := strconv.Atoi(req.Amount)

		parsedPrice := req.Price
		
		parsedSelling := req.Selling
		
		parsedMinPrice := req.MinPrice
		

		msg := types.NewMsgSetResource(
			creator,
			req.ID,
			parsedName,
			parsedAmount,
			parsedPrice,
			parsedSelling,
			parsedMinPrice,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteResourceRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	ID 		string `json:"id"`
}

func deleteResourceHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteResourceRequest
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
		msg := types.NewMsgDeleteResource(req.ID, creator)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
