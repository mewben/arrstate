import React from "react"
import cx from "clsx"

import { extractError } from "@Utils"

const Error = ({ error, className }) => {
  if (!error) {
    return null
  }

  return (
    <div className={cx("text-red-500", className)}>{extractError(error)}</div>
  )
}

export default Error
