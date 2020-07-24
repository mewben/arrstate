import React from "react"

import { useMeContext } from "@Wrappers"
import { formatDate } from "@Utils"

export const Time = ({ d, format, dateOnly, timeOnly }) => {
  const {
    currentUser: { person },
  } = useMeContext()
  if (!d) {
    return null
  }

  let f = format
  if (!f) {
    // get from the person locale
    if (dateOnly) {
      f = person?.locale?.dateFormat
    } else if (timeOnly) {
      f = person?.locale?.timeFormat
    } else {
      f = person?.locale?.dateFormat + " " + person?.locale?.timeFormat
    }
  }
  return <span>{formatDate(d, f)}</span>
}
