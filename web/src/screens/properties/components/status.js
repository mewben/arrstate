import React from "react"

import { Badge } from "@Components/generic"

const Status = ({ status }) => {
  let color = ""
  if (status === "acquired") {
    color = "red"
  } else if (status === "ongoing") {
    color = "yellow"
  } else if (status === "available") {
    color = "green"
  }
  return <Badge text={status} color={color} />
}

export default Status
