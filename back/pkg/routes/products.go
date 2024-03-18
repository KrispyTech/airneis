package routes

import (
	controllers "github.com/KrispyTech/airneis/src/controllers/products"
	"github.com/gofiber/fiber/v2"
)

const (
	productsRoute       = "products"
	productsRouteWithID = "products/:productID"
)

func SetupProductsRoutes(app fiber.Router) {
	app.Get(productsRoute, controllers.GetProducts)
	app.Get(productsRouteWithID, controllers.GetProductById)
	app.Post(productsRoute, controllers.CreateProduct)
	app.Patch(productsRouteWithID, controllers.UpdateProduct)
	app.Delete(productsRouteWithID, controllers.DeleteProduct)
}
