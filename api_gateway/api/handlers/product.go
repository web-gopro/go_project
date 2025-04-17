package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"gitlab.com/e-market724/back-end/api_gateway/genproto/product_service"
)

func (h *Handlers) ProductCreate(ctx *gin.Context) {

	var req product_service.ProductCreateReq

	ctx.BindJSON(&req)

	resp, err := h.services.GetProductSevice().CreateProduct(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)

}


func (h *Handlers)ProductGet(ctx *gin.Context){

	var req product_service.GetByIdReq

	req.Id=ctx.Param("id")

	resp,err:=h.services.GetProductSevice().GetProduct(context.Background(),&req)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)


}


func (h *Handlers)ProductsGet(ctx *gin.Context){

	var req product_service.GetListReq

	ctx.BindJSON(&req)

	resp,err:=h.services.GetProductSevice().GetProducts(context.Background(),&req)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)

}



func (h *Handlers)ProductUpdate(ctx *gin.Context){

	var req product_service.ProductUpdateReq

	ctx.BindJSON(&req)

	resp,err:=h.services.GetProductSevice().UpdateProduct(context.Background(),&req)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)

}


func (h *Handlers)ProductDelete(ctx *gin.Context){

	var req product_service.DeleteReq

	req.Id=ctx.Param("id")

	resp,err:=h.services.GetProductSevice().DeleteProduct(context.Background(),&req)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)


}



func (h *Handlers) ProductImgCreate(ctx *gin.Context) {

	var req product_service.ProductImage

	ctx.BindJSON(&req)

	resp, err := h.services.GetProductSevice().CreateProductImg(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)

}


func (h *Handlers)ProductImgGet(ctx *gin.Context){

	var req product_service.GetByIdReq

	req.Id=ctx.Param("id")

	resp,err:=h.services.GetProductSevice().GetProductImg(context.Background(),&req)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)


}


func (h *Handlers)ProductImgsGet(ctx *gin.Context){

	var req product_service.GetListReq

	ctx.BindJSON(&req)

	resp,err:=h.services.GetProductSevice().GetProductImgs(context.Background(),&req)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, resp)

}



func (h *Handlers)ProductImgUpdate(ctx *gin.Context){

	var req product_service.ProductImage

	ctx.BindJSON(&req)

	resp,err:=h.services.GetProductSevice().UpdateProductImg(context.Background(),&req)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)

}


func (h *Handlers)ProductImgDelete(ctx *gin.Context){

	var req product_service.DeleteReq

	req.Id=ctx.Param("id")

	resp,err:=h.services.GetProductSevice().DeleteProductImg(context.Background(),&req)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(201, resp)


}
