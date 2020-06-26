import React from "react"

import { extractError } from "@Utils"

const Error = ({ error }) => {
  if (!error) {
    return null
  }

  return <div style={{ color: "red" }}>{extractError(error)}</div>
}

export default Error
