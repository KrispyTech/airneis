package lib

import (
	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/KrispyTech/airneis/src/controllers/cart"
	"github.com/KrispyTech/airneis/src/controllers/categories"
	"github.com/KrispyTech/airneis/src/controllers/contact"
	"github.com/KrispyTech/airneis/src/controllers/products"
	"github.com/KrispyTech/airneis/src/controllers/sessions"
	"github.com/KrispyTech/airneis/src/controllers/users"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

const (
	categoriesRoute        = "categories"
	categoriesRouteWithID  = "categories/:categoryID"
	productsRoute          = "products"
	productsRouteWithID    = "products/:productID"
	sessionsRoute          = "sessions"
	userRoute              = "users"
	cartRoute              = "cart"
	cartRouteWithProductID = "cart/:productID"
	contactRoute           = "contact/"
	contactRouteWithID     = "contact/:contactID"
)

func SetupRoutes(router fiber.Router, services Services) fiber.Router {
	SetProductsRoutes(router)
	SetCategoriesRoutes(router)
	SetUsersRoutes(router)
	SetSessionsRoutes(router, services.VaultClient)
	SetupCartRoutes(router, services.VaultClient)
	log.Info("Routes have been set")

	return router
}

func SetCategoriesRoutes(app fiber.Router) {
	app.Get(categoriesRoute, categories.GetCategories)
	app.Get(categoriesRouteWithID, categories.GetCategoryByID)
	app.Post(categoriesRoute, categories.CreateCategory)
	app.Patch(categoriesRouteWithID, categories.EditCategory)
	app.Delete(categoriesRouteWithID, categories.DeleteCategory)
}

func SetContactRoutes(app fiber.Router) {
	app.Get(contactRoute, contact.GetContactTickets)
	app.Get(contactRouteWithID, contact.GetContactByID)
	app.Post(contactRoute, contact.CreateContactTicket)
	app.Patch(contactRouteWithID, contact.EditContactTicket)
	app.Delete(contactRouteWithID, contact.DeleteContactTicket)
}

func SetProductsRoutes(app fiber.Router) {
	app.Get(productsRoute, products.GetProducts)
	app.Get(productsRouteWithID, products.GetProductByID)
	app.Post(productsRoute, products.CreateProduct)
	app.Patch(productsRouteWithID, products.EditProduct)
	app.Delete(productsRouteWithID, products.DeleteProduct)
}

func SetSessionsRoutes(router fiber.Router, vlt vault.VaultClient) fiber.Router {
	router.Post(sessionsRoute, func(ctx *fiber.Ctx) error {
		return sessions.ValidateSession(ctx, vlt)
	})

	router.Delete(sessionsRoute, func(ctx *fiber.Ctx) error {
		return sessions.DestroySession(ctx)
	})

	return router
}

func SetUsersRoutes(router fiber.Router) {
	router.Post(userRoute, users.CreateUser)
}

func SetupCartRoutes(router fiber.Router, vc vault.VaultClient) {
	router.Use(authMiddleware(vc))
	router.Get(cartRoute, cart.GetCart)
	router.Post(cartRouteWithProductID, cart.AddProductToCart)
	router.Patch(cartRouteWithProductID, cart.UpdateCartQuantity)
	router.Delete(cartRouteWithProductID, cart.DeleteProductFromCart)
}
