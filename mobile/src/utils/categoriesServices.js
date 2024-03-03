import apiRoutes from "./constants/apiRoutes"
import mobileConfig from "../../mobileConfig"
import axios from "axios"

export const getCategories = async () => await axios.get(mobileConfig.api.url + apiRoutes.categories)
export const getCategoryById = async id => await axios.get(`${mobileConfig.api.url}${apiRoutes.categories}/${id}`)
