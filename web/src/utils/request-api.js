import axios from "axios"

import { authToken } from "@Providers"

// const backend = process.env.GATSBY_BACKEND

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
