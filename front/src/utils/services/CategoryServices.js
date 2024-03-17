import { apiRoutes } from "@/utils/constants"
import { del, get, patch, post } from "@/utils/httpClient"

export const getCategories = async () => await get(apiRoutes.categories)
export const getCategoryById = async id => await get(apiRoutes.categoryId(id))
export const createCategory = async (category, token) => await post(apiRoutes.categories, token, category)
export const deleteCategory = async (id, token) => await del(apiRoutes.categoryId(id), token)
export const updateCategory = async (id, updatedCategory, token) => await patch(apiRoutes.categoryId(id), token, updatedCategory)
