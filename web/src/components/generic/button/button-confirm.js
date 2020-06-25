import React from "react"
import PropTypes from "prop-types"

import { Button } from "@Components/generic"

const ButtonConfirm = ({
  onConfirm,
  onReject,
  children = "Delete",
  confirmComponent = "Are you sure?",
  isDisabled,
}) => {
  const [isOn, setIsOn] = React.useState(false)

  const onClick = () => {
    if (isOn) {
      onConfirm()
    } else {
      setIsOn(true)
    }
  }

  const onBlur = () => {
    setIsOn(false)
  }

  return (
    <Button
      onClick={onClick}
      type="button"
      isDisabled={isDisabled}
      onBlur={onBlur}
    >
      {isOn ? confirmComponent : children}
    </Button>
  )
}

ButtonConfirm.propTypes = {
  children: PropTypes.any,
  confirmComponent: PropTypes.any,
  onConfirm: PropTypes.func.isRequired,
  isDisabled: PropTypes.bool,
}

export default ButtonConfirm
