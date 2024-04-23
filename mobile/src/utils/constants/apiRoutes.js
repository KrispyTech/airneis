const apiRoutes = {
  categories: "categories",
  categoryId: id => `${apiRoutes.categories}/${id}`,
  products: "products",
  productId: id => `${apiRoutes.products}/${id}`
}

export default apiRoutes
