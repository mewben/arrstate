import React from "react"
import PropTypes from "prop-types"

const Button = ({ children, type = "button", isDisabled }) => {
  return (
    <button type={type} disabled={isDisabled}>
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
