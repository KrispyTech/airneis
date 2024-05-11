import { apiRoutes } from "@/utils/constants"
import { del, get, patch, post } from "@/utils/httpClient"

export const getUsers = token => get(apiRoutes.users, token)
export const getUsersById = id => get(apiRoutes.userId(id))
export const createUser = (user, token) => post(apiRoutes.users, token, user)
export const deleteUser = (id, token) => del(apiRoutes.userId(id), token)
export const updateUser = (id, token, updatedUserFields) => patch(apiRoutes.userId(id), token, updatedUserFields)
