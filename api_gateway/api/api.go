package api

import (
	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
	"gitlab.com/e-market724/back-end/api_gateway/api/handlers"
	"gitlab.com/e-market724/back-end/api_gateway/redis"
	"gitlab.com/e-market724/back-end/api_gateway/services"
)

type Options struct {
	Services services.ServiceManagerI
	Redis redis.RedisRepoI
	Log logger.LoggerI
}

func Api(o Options)*gin.Engine {

	engine := gin.Default()

	h:=handlers.NewHandler(o.Log,o.Services,o.Redis)

	api := engine.Group("/api")

	pd:=api.Group("/pd")


	//product
	pd.POST("product",h.ProductCreate)
	pd.GET("product",h.ProductGet)
	pd.POST("products",h.ProductsGet)
	pd.POST("product_update",h.ProductUpdate)
	pd.POST("product_delete",h.ProductDelete)


	//product_imge
	pd.POST("product_imge",h.ProductImgCreate)
	pd.GET("product_imge",h.ProductImgGet)
	pd.POST("product_imges",h.ProductImgsGet)
	pd.POST("product_imge_update",h.ProductImgUpdate)
	pd.POST("product_imge_delete",h.ProductImgDelete)


	//category
	pd.POST("category",h.CategoryCreate)
	pd.GET("category",h.CategoryGet)
	pd.POST("categories",h.CategoriesGet)
	pd.POST("category_update",h.CategoryUpdate)
	pd.POST("category_delete",h.CategoryDelete)


	//subcategory
	pd.POST("subcategory",h.SubCategoryCreate)
	pd.GET("subcategory",h.SubCategoryGet)
	pd.POST("subcategories",h.SubCategoriesGet)
	pd.POST("subcategory_update",h.SubCategoryUpdate)
	pd.POST("subcategory_delete",h.SubCategoryDelete)

	engine.Run(":8080")

	return engine
}
