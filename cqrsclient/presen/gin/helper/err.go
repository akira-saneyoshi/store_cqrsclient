package helper

import (
	"net/http"

	v1 "github.com/akira-saneyoshi/store_pb/pb/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// エラーレスポンス送信
func ErrResp(ctx *gin.Context, err any) {
	if pbErr, ok := err.(*v1.Error); ok {
		if pbErr.Type == "Domain Error" || pbErr.Type == "CRUD Error" {
			ctx.JSON(http.StatusBadRequest, pbErr)
		} else {
			ctx.JSON(http.StatusInternalServerError, pbErr)
		}
	} else {
		other, _ := err.(error)
		// エラーは*status.Statusに変換可能?
		if status, ok := status.FromError(other); ok {
			// ステータスはBad Request?
			if status.Code() == codes.InvalidArgument {
				ctx.JSON(http.StatusBadRequest,
					v1.Error{Type: "Validate Error", Message: status.Message()})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError,
				v1.Error{Type: "InterError Error", Message: other.Error()})
		}
	}
}
