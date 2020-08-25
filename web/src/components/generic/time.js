import React from "react"

import { useCurrentContext } from "@Wrappers"
import { formatDate } from "@Utils"

export const Time = ({ d, format, dateOnly, timeOnly }) => {
  const {
    currentPerson: { locale },
  } = useCurrentContext()
  if (!d) {
    return null
  }

  let f = format
  if (!f) {
    // get from the person locale
    if (dateOnly) {
      f = locale?.dateFormat
    } else if (timeOnly) {
      f = locale?.timeFormat
    } else {
      f = locale?.dateFormat + " " + locale?.timeFormat
    }
  }
  return <span>{formatDate(d, f)}</span>
}
