import apiRoutes from "../constants/apiRoutes"
import {get} from "../httpClient"

export const getCategories = () => get(apiRoutes.categories)
export const getCategoryById = id => get(apiRoutes.categoryId(id))
