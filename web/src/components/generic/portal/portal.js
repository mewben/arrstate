import React from "react"
import PropTypes from "prop-types"
import Dialog from "@material-ui/core/Dialog"
import Drawer from "@material-ui/core/Drawer"

const Portal = ({
  openByClickOn,
  defaultOpen,
  mode = "drawer",
  anchor = "right",
  children,
  ...props
}) => {
  const [isOpen, setIsOpen] = React.useState(defaultOpen)

  const openPortal = () => {
    setIsOpen(true)
  }

  const closePortal = () => {
    setIsOpen(false)
  }

  const content = React.cloneElement(children, { onClose: closePortal })

  const renderContent = () => {
    if (mode === "modal") {
      // todo
      return content
    } else {
      return (
        <Drawer open={isOpen} anchor={anchor} onClose={closePortal}>
          {content}
        </Drawer>
      )
    }
  }

  return (
    <>
      {openByClickOn &&
        React.cloneElement(openByClickOn, {
          onClick: openPortal,
        })}
      {renderContent()}
    </>
  )
}

Portal.propTypes = {
  mode: PropTypes.oneOf(["drawer", "modal"]),
  anchor: PropTypes.oneOf(["left", "right", "top", "bottom"]),
  openByClickOn: PropTypes.any,
  children: PropTypes.any,
  defaultOpen: PropTypes.bool,
}

export default Portal
