import React from "react"
import PropTypes from "prop-types"
import cx from "clsx"

const Th = ({ children, align = "left" }) => {
  const cxAlign =
    align === "center"
      ? "text-center"
      : align === "right"
      ? "text-right"
      : "text-left"

  return (
    <th
      className={cx(
        "px-6 py-3 border-b border-gray-200 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider",
        cxAlign
      )}
    >
      {children}
    </th>
  )
}

Th.propTypes = {
  children: PropTypes.any,
  align: PropTypes.oneOf(["left", "center", "right"]),
}

export default Th
