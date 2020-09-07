import React from "react"
import cx from "clsx"

import { initialsName } from "@Utils"

const PersonBadge = ({ person, size = "sm" }) => {
  const initials = initialsName(person?.name)

  const cl = cx(
    "flex items-center justify-center bg-cool-gray-200 text-cool-gray-500 uppercase",
    size === "sm" && "w-6 h-6 rounded-sm text-xs",
    size === "2xl" && "w-40 h-40 rounded-lg border-4 border-white text-5xl"
  )

  return (
    <div className={cl}>
      <span>{initials}</span>
    </div>
  )
}

export default PersonBadge
