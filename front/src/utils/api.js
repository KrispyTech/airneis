import axios from "axios"
import config from "@/utils/config"

// Don't use this function instead use a function according to your method
const request =
  method =>
  async (route, token, data = {}) =>
    await axios[method.toLowerCase()](`${config.api.url}/${route}`, data, {
      header: {
        authorization: `BEARER ${token}`
      }
    })

export const getRequest = request("GET")
export const postRequest = request("POST")
export const deleteRequest = request("DELETE")
export const patchRequest = request("PATCH")
export const putRequest = request("PUT")
