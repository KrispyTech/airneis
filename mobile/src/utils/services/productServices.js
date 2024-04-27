import apiRoutes from "../constants/apiRoutes"
import {get} from "../httpClient"

export const getProductById = id => get(apiRoutes.productId(id))
