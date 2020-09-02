import React from "react"
import PropTypes from "prop-types"
import cx from "clsx"

const Td = ({ children, className, align = "left", py, wrap, ...props }) => {
  const cxAlign =
    align === "center"
      ? "text-center"
      : align === "right"
      ? "text-right"
      : "text-left"

  return (
    <td
      className={cx(
        "px-6 text-sm leading-5 text-cool-gray-500",
        cxAlign,
        wrap ? "whitespace-normal" : "whitespace-no-wrap",
        py ? py : "py-4",
        className
      )}
      {...props}
    >
      {children}
    </td>
  )
}

const Td2 = ({ children, className, align = "left", ...props }) => {
  const cxAlign =
    align === "center"
      ? "text-center"
      : align === "right"
      ? "text-right"
      : "text-left"

  return (
    <div
      className={cx(
        "table-cell px-6 py-4 whitespace-no-wrap2 border-b border-gray-200 text-sm leading-5 text-gray-500",
        cxAlign,
        className
      )}
      {...props}
    >
      <div className="flex relative break-words whitespace-no-wrap">
        {children}
      </div>
    </div>
  )
}

Td.propTypes = {
  children: PropTypes.any,
  align: PropTypes.oneOf(["left", "center", "right"]),
}

export default Td
