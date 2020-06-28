import React from "react"
import PropTypes from "prop-types"
import cx from "clsx"

const Button = ({
  children,
  type = "button",
  size = "md",
  fullWidth,
  isDisabled,
  ...props
}) => {
  let cxSize =
    size === "xl"
      ? "px-6 py-3 text-base"
      : size === "lg"
      ? "px-4 py-2 text-base"
      : size === "md"
      ? "px-4 py-2 text-sm"
      : size === "sm"
      ? "px-3 py-2 text-sm"
      : "px-2.5 py-1.5 text-xs"
  return (
    <span
      className={cx(
        "inline-flex rounded-md shadow-sm",
        fullWidth ? "w-full" : ""
      )}
    >
      <button
        type={type}
        disabled={isDisabled}
        className={cx(
          "inline-flex items-center border border-transparent leading-4 font-medium rounded text-white bg-cool-gray-600 hover:bg-cool-gray-500 focus:outline-none focus:border-cool-gray-700 focus:shadow-outline-cool-gray active:bg-cool-gray-700 transition ease-in-out duration-150",
          cxSize,
          fullWidth ? "w-full justify-center" : ""
        )}
        {...props}
      >
        {children}
      </button>
    </span>
  )
}

Button.propTypes = {
  children: PropTypes.any.isRequired,
  type: PropTypes.string,
  isDisabled: PropTypes.bool,
  size: PropTypes.oneOf(["xs", "sm", "md", "lg", "xl"]),
}

export default Button
