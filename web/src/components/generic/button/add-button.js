import React from "react"
import PropTypes from "prop-types"
import cx from "clsx"

const AddButton = ({ text = "Add", onClick, icon, to, children }) => {
  return (
    <span className={cx("inline-flex rounded-md shadow-sm")}>
      <button type="button" onClick={onClick}>
        {text}
      </button>
    </span>
  )
}

AddButton.propTypes = {
  text: PropTypes.string,
  icon: PropTypes.element,
  to: PropTypes.string,
  children: PropTypes.element,
}

export default AddButton
