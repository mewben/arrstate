import React from "react"
import PropTypes from "prop-types"
import cx from "clsx"

const Td = ({ children, align = "left", ...props }) => {
  const cxAlign =
    align === "center"
      ? "text-center"
      : align === "right"
      ? "text-right"
      : "text-left"

  return (
    <td
      className={cx(
        "px-6 py-4 whitespace-no-wrap border-b border-gray-200 text-sm leading-5 text-gray-500",
        cxAlign
      )}
      {...props}
    >
      {children}
    </td>
  )
}

Td.propTypes = {
  children: PropTypes.any,
  align: PropTypes.oneOf(["left", "center", "right"]),
}

export default Td
