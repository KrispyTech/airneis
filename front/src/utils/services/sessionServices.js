import { apiRoutes } from "@/utils/constants"
import { del, post } from "@/utils/httpClient"

export const createSession = (email, password) => post(apiRoutes.session, null, { email, password })
export const deleteSession = () => del(apiRoutes.session)
