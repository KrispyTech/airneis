import { apiRoutes } from "@/utils/constants"
import { del, get, patch, post } from "@/utils/httpClient"

export const getCategories = () => get(apiRoutes.categories)
export const getCategoryById = id => get(apiRoutes.categoryId(id))
export const createCategory = (category, token) => post(apiRoutes.categories, token, category)
export const deleteCategory = (id, token) => del(apiRoutes.categoryId(id), token)
export const updateCategory = (id, updatedCategory, token) => patch(apiRoutes.categoryId(id), token, updatedCategory)
