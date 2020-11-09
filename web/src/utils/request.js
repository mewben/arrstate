import axios from "axios"

import { authToken } from "@Providers"

export const requestApi = async (
  path = "",
  method = "GET",
  { data, params, noToken = false } = {}
) => {
  let headers = {
    "Content-Type": "application/json",
  }

  if (!noToken) {
    // attach bearer token
    headers["Authorization"] = authToken.authorizationString
  }

  return axios({
    method,
    url: path,
    headers,
    data,
    params,
    responseType: "json",
  })
}

export const publicFetcher = (url, params) => {
  return requestApi(url, "GET", { params, noToken: true }).then(res => res.data)
}

export const privateFetcher = (url, params) => {
  return requestApi(url, "GET", { params }).then(res => res.data)
}
