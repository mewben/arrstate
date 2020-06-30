import React from "react"
import Menu from "@material-ui/core/Menu"

const Dropdown = ({
  children,
  openByClickOn,
  align = "right",
  maxHeight = 280,
  minWidth = 200,
}) => {
  const [anchorEl, setAnchorEl] = React.useState()

  const onOpen = e => {
    setAnchorEl(e.currentTarget)
  }

  const onClose = () => {
    setAnchorEl(null)
  }

  return (
    <>
      {!!openByClickOn &&
        React.cloneElement(openByClickOn, {
          onClick: onOpen,
        })}
      <Menu
        // keepMounted
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={onClose}
        PaperProps={{
          style: { maxHeight, minWidth },
        }}
        anchorOrigin={{
          vertical: "bottom",
          horizontal: align,
        }}
        transformOrigin={{
          vertical: "top",
          horizontal: align,
        }}
        getContentAnchorEl={null}
        // eslint-disable-next-line
        autoFocus={false}
        disableAutoFocusItem={true} // to be able to type "T"
        variant="menu"
      >
        {React.cloneElement(children, { onClose: onClose })}
      </Menu>
    </>
  )
}

export default Dropdown
