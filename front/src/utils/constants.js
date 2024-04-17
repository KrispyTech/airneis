export const apiRoutes = {
  categories: "api/categories",
  categoryId: id => `${apiRoutes.categories}/${id}`,
  products: "api/products",
  productId: id => `${apiRoutes.products}/${id}`
}
