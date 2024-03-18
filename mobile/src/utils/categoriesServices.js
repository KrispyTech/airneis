import apiRoutes from "./constants/apiRoutes"
import { get } from "./httpClient"

export const getCategories = async () => await get(apiRoutes.categories)
export const getCategoryById = async id => await get(apiRoutes.categoryId(id))
