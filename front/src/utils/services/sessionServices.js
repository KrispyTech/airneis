import { apiRoutes } from "@/utils/constants"
import { del, post } from "@/utils/httpClient"

export const createSession = (email, password, token) => post(apiRoutes.session, token, { email, password })
export const deleteSession = () => del(apiRoutes.session)
