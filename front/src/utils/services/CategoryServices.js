import { deleteRequest, getRequest, patchRequest, postRequest } from "@/utils/api"
import { apiRoutes } from "@/utils/constants"

export const getCategories = async () => await getRequest(apiRoutes.categories)
export const getCategoryById = async id => await getRequest(apiRoutes.categoryId(id))
export const createCategory = async (category, token) => await postRequest(apiRoutes.categories, token, category)
export const deleteCategory = async (id, token) => await deleteRequest(apiRoutes.categoryId(id), token)
export const updateCategory = async (token, id, updatedCategory) => await patchRequest(apiRoutes.categoryId(id), token, updatedCategory)
