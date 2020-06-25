import React from "react"
import PropTypes from "prop-types"

const Portal = ({ openByClickOn, defaultOpen, children }) => {
  const [isOpen, setIsOpen] = React.useState(defaultOpen)

  const openPortal = () => {
    setIsOpen(true)
  }

  const closePortal = () => {
    setIsOpen(false)
  }

  return (
    <>
      {openByClickOn &&
        React.cloneElement(openByClickOn, {
          onClick: openPortal,
        })}
      {isOpen && React.cloneElement(children, { onClose: closePortal })}
    </>
  )
}

Portal.propTypes = {
  openByClickOn: PropTypes.any,
  children: PropTypes.any,
  defaultOpen: PropTypes.bool,
}

export default Portal
