import { apiRoutes } from "@/utils/constants"
import { del, get, patch, post } from "@/utils/httpClient"

export const getProducts = () => get(apiRoutes.products)
export const getProductById = id => get(apiRoutes.productId(id))
export const createProduct = (product, token) => post(apiRoutes.products, token, product)
export const deleteProduct = (id, token) => del(apiRoutes.productId(id), token)
export const updateProduct = (id, updatedProduct, token) => patch(apiRoutes.productId(id), token, updatedProduct)
