import axios from "axios"
import config from "@/utils/config"

const makeRequest =
  method =>
  async (route, token, data = {}) => {
    if (token === "") {
      return await axios[method.toLowerCase()](`${config.api.url}/${route}`, data, {})
    }

    return await axios[method.toLowerCase()](`${config.api.url}/${route}`, data, {
      headers: {
        authorization: `Bearer ${token}`
      }
    })
  }

export const get = makeRequest("GET")
export const post = makeRequest("POST")
export const del = makeRequest("DELETE")
export const patch = makeRequest("PATCH")
export const put = makeRequest("PUT")
