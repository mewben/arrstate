import React from "react"
import PropTypes from "prop-types"
import cx from "clsx"

const Panel = ({ children, className, noPadding }) => {
  const withPadding = <div className="px-4 py-5 sm:p-6">{children}</div>
  return (
    <div
      className={cx("bg-white overflow-hidden rounded-lg shadow-sm", className)}
    >
      {noPadding ? children : withPadding}
    </div>
  )
}

Panel.propTypes = {
  children: PropTypes.any,
  className: PropTypes.string,
  noPadding: PropTypes.bool,
}

export default Panel
