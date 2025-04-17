package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"gitlab.com/e-market724/back-end/api_gateway/genproto/product_service"
)

func (h *Handlers) CategoryCreate(ctx *gin.Context) {

	var req product_service.CategoryCreateReq

	ctx.BindJSON(&req)

	resp, err := h.services.GetProductSevice().CreateCategory(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)
}

func (h *Handlers) CategoryGet(ctx *gin.Context) {

	var req product_service.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.services.GetProductSevice().GetCategory(context.Background(), &req)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *Handlers) CategoriesGet(ctx *gin.Context) {

	var req product_service.GetListReq

	ctx.BindJSON(&req)

	resp, err := h.services.GetProductSevice().GetCategories(context.Background(), &req)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *Handlers) CategoryUpdate(ctx *gin.Context) {

	var req product_service.CategoryUpdateReq

	ctx.BindJSON(&req)

	resp, err := h.services.GetProductSevice().UpdateCategory(context.Background(), &req)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)
}


func (h *Handlers) CategoryDelete(ctx *gin.Context) {

	var req product_service.DeleteReq

	req.Id = ctx.Param("id")

	resp, err := h.services.GetProductSevice().DeleteCategory(context.Background(), &req)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)
}


func (h *Handlers)SubCategoryCreate(ctx *gin.Context) {

	var req product_service.SubCategoryCreateReq

	ctx.BindJSON(&req)

	resp, err := h.services.GetProductSevice().SubCreateCategory(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)
}

func (h *Handlers) SubCategoryGet(ctx *gin.Context) {

	var req product_service.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.services.GetProductSevice().GetSubCategory(context.Background(), &req)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *Handlers) SubCategoriesGet(ctx *gin.Context) {

	var req product_service.GetListReq

	ctx.BindJSON(&req)

	resp, err := h.services.GetProductSevice().GetSubCategories(context.Background(), &req)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *Handlers) SubCategoryUpdate(ctx *gin.Context) {

	var req product_service.SubCategoryUpdateReq

	ctx.BindJSON(&req)

	resp, err := h.services.GetProductSevice().UpdateSubCategory(context.Background(), &req)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)
}


func (h *Handlers) SubCategoryDelete(ctx *gin.Context) {

	var req product_service.DeleteReq

	req.Id = ctx.Param("id")

	resp, err := h.services.GetProductSevice().DeleteSubCategory(context.Background(), &req)

	if err != nil {

		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)
}