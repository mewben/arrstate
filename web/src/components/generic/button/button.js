import React from "react"
import PropTypes from "prop-types"

const Button = ({ children, type = "button", isDisabled, ...props }) => {
  return (
    <button type={type} disabled={isDisabled} {...props}>
      {children}
    </button>
  )
}

Button.propTypes = {
  children: PropTypes.any.isRequired,
  type: PropTypes.string,
  isDisabled: PropTypes.bool,
}

export default Button
