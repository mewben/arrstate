import React from "react"
import cx from "clsx"

const Badge = ({ text, color }) => {
  let variant = "text-gray-800 bg-gray-100"
  if (color === "green") {
    variant = "text-green-800 bg-green-100"
  } else if (color === "red") {
    variant = "text-red-800 bg-red-100"
  } else if (color === "yellow") {
    variant = "text-yellow-800 bg-yellow-100"
  }
  return (
    <span
      className={cx(
        "flex-shrink-0 inline-block px-2 py-0.5 text-xs leading-4 font-medium capitalize rounded-full",
        variant
      )}
    >
      {text}
    </span>
  )
}

export default Badge
