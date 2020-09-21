import React from "react"
import { Link } from "gatsby"
import PropTypes from "prop-types"
import cx from "clsx"

const Button = ({
  children,
  variant = "text",
  type = "button",
  to,
  size = "md",
  color = "cool-gray",
  circle,
  fullWidth,
  isDisabled,
  isLoading,
  ...props
}) => {
  let Component = "button"
  if (to) {
    Component = Link
  }

  return (
    <Component
      type={!to ? type : undefined}
      to={to ? to : undefined}
      disabled={isDisabled || isLoading}
      className={cx("btn", color, size, {
        "btn-contained": variant === "contained",
        "btn-outlined": variant === "outlined",
        "btn-text": variant === "text",
        "btn-circle": circle,
        "w-full": fullWidth,
      })}
      {...props}
    >
      {children}
    </Component>
  )
}

const Button2 = ({
  children,
  to,
  type = "button",
  size = "md",
  color = "cool-gray",
  fullWidth,
  isDisabled,
  ...props
}) => {
  let Component = "button"
  if (to) {
    Component = Link
  }
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

  let cxColor =
    color === "white"
      ? "text-cool-gray-900 bg-white border border-cool-gray-200 hover:bg-cool-gray-100 focus:border-cool-gray-700 focus:shadow-outline-cool-gray active:bg-cool-gray-200"
      : "text-white bg-cool-gray-600 hover:bg-cool-gray-500 focus:border-cool-gray-700 focus:shadow-outline-cool-gray active:bg-cool-gray-700"
  return (
    <span
      className={cx(
        "inline-flex rounded-md shadow-sm",
        fullWidth ? "w-full" : ""
      )}
    >
      <Component
        type={!to ? type : undefined}
        to={to ? to : undefined}
        disabled={isDisabled}
        className={cx(
          "inline-flex items-center border border-transparent leading-5 font-medium rounded focus:outline-none transition ease-in-out duration-150",
          cxSize,
          cxColor,
          fullWidth ? "w-full justify-center" : ""
        )}
        {...props}
      >
        {children}
      </Component>
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
