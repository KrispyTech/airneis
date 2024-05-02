import apiRoutes from "../constants/apiRoutes"
import { post } from "../httpClient"

export const createSession = async data => await post(apiRoutes.sessions, "", data)
export const destroySession = async () => await post(apiRoutes.sessions, "", {})
