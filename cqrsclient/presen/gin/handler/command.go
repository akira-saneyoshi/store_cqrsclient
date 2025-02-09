package handler

import (
	"cqrsclient/presen/gin/helper"

	"github.com/gin-gonic/gin"
)

// 更新リクエストハンドラ
type CommandHandler struct {
	helper *helper.CommandHelper
}

// コンストラクタ
func NewCommandHandler(helper *helper.CommandHelper) *CommandHandler {
	return &CommandHandler{helper: helper}
}

// カテゴリ登録をする
// @tags カテゴリ
// @Summary 登録する
// @Description 新しいカテゴリを登録する
// @ID add-category
// @Accept application/json
// @Produce json
// @Param newcategory body v1.CategoryUpParam true "crud:1(更新種別),name(カテゴリ名)"
// @Success 200 {object} v1.Category
// @Failure 400 {object} v1.Error
// @Failure 500 {object} v1.Error
// @Router /category/add [post]
func (ins *CommandHandler) CreateCategory(ctx *gin.Context) {
	ins.helper.CreateCategory(ctx)
}

// 商品カテゴリを変更する
// @tags カテゴリ
// @Summary 変更する
// @Description 指定された識別子のカテゴリを変更する
// @ID update-category
// @Accept application/json
// @Produce json
// @Param updatecategory body v1.CategoryUpParam true "crud:2(更新種別),id(カテゴリID),name(カテゴリ名)"
// @Success 200 {object} v1.Category
// @Failure 400 {object} v1.Error
// @Failure 500 {object} v1.Error
// @Router /category/update [put]
func (ins *CommandHandler) UpdateCategory(ctx *gin.Context) {
	ins.helper.UpdateCategory(ctx)
}

// カテゴリを削除する
// @tags カテゴリ
// @Summary 削除する
// @Description 指定された識別子でカテゴリを削除する
// @ID delete-category
// @Accept application/json
// @Produce json
// @Param deletecategory body v1.CategoryUpParam true "crud:3(更新種別),id(カテゴリID)"
// @Success 200 {object} v1.Category
// @Failure 400 {object} v1.Error
// @Failure 500 {object} v1.Error
// @Router /category/delete [delete]
func (ins *CommandHandler) DeleteCategory(ctx *gin.Context) {
	ins.helper.DeleteCategory(ctx)
}

// 商品を登録する
// @tags 商品
// @Summary 登録する
// @Description 新商品を登録する
// @ID add-product
// @Accept application/json
// @Produce json
// @Param newproduct body v1.ProductUpParam true "crud:1(更新種別),name(商品名),price(単価),categoryId(商品カテゴリ番号)"
// @Success 200 {object} v1.Product
// @Failure 400 {object} v1.Error
// @Failure 500 {object} v1.Error
// @Router /product/add [post]
func (ins *CommandHandler) CreateProduct(ctx *gin.Context) {
	ins.helper.CreateProduct(ctx)
}

// 商品を変更する
// @tags 商品
// @Summary 変更する
// @Description 指定された識別子の商品を変更する
// @ID update-product
// @Accept application/json
// @Produce json
// @Param updateproduct body v1.ProductUpParam true "crud:2(更新種別),id(商品ID),name(商品名),price(単価),categoryId(商品カテゴリ番号)"
// @Success 200 {object} v1.Product
// @Failure 400 {object} v1.Error
// @Failure 500 {object} v1.Error
// @Router /product/update [put]
func (ins *CommandHandler) UpdateProduct(ctx *gin.Context) {
	ins.helper.UpdateProduct(ctx)
}

// 商品を削除する
// @tags 商品
// @Summary 削除する
// @Description 指定された識別子で商品を削除する
// @ID delete-product
// @Accept application/json
// @Produce json
// @Param deleteproduct body v1.ProductUpParam true "crud:3(更新種別),id(商品ID)"
// @Success 200 {object} v1.Product
// @Failure 400 {object} v1.Error
// @Failure 500 {object} v1.Error
// @Router /product/delete [delete]
func (ins *CommandHandler) DeleteProduct(ctx *gin.Context) {
	ins.helper.DeleteProduct(ctx)
}
