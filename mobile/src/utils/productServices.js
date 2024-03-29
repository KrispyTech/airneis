import apiRoutes from "./constants/apiRoutes"
import { get } from "./httpClient"

export const getProductById = async id => await get(apiRoutes.productId(id))
