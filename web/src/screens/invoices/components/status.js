import React from "react"

import { Badge } from "@Components/generic"

const Status = ({ status }) => {
  let color = ""
  if (status === "overdue") {
    color = "red"
  } else if (status === "pending") {
    color = "yellow"
  } else if (status === "paid") {
    color = "green"
  }
  return <Badge text={status} color={color} />
}

export default Status
