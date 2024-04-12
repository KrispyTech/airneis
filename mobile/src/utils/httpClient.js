import axios from "axios"
import mobileConfig from "../../mobileConfig"

const makeRequest =
  method =>
  async (route, token, data = {}) => {
    if (token === "") {
      return await axios[method.toLowerCase()](`${mobileConfig.api.url}/${route}`, data, {})
    }

    return await axios[method.toLowerCase()](`${mobileConfig.api.url}/${route}`, data, {
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
