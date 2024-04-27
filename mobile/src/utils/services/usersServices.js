import { post } from "../httpClient"
import apiRoutes from "../constants/apiRoutes"

export const createUsers = async data => await post(apiRoutes.users, data)
