import isPromise from "is-promise"

import { extractError } from "@Utils/"

// middleware to catch all promise rejects
export default function errorMiddleware() {
  return next => action => {
    // If not a promise, continue on
    if (!isPromise(action.payload)) {
      return next(action)
    }

    // Dispatch initial pending promise, but catch error
    return next(action).catch(error => {
      return {
        error: true,
        data: extractError(error),
      }
    })
  }
}
