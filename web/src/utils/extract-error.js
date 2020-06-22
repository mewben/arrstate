import { get } from "@Utils/lodash"

// extract the error message from any error
export const extractError = error => {
  console.log("Error: ", error)
  let message = "Server error"

  if (error.response) {
    console.log("error response", error.response)
    message = get(
      error,
      "response.data.message",
      get(error, "response.statusText")
    )
  } else if (error.request) {
    console.log("errequest", error.request)
    message = "Request error"
  } else if (error.payload) {
    message = get(error, "payload.message")
  }
  return message
}
