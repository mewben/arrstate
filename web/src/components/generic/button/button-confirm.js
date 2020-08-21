import React from "react"
import PropTypes from "prop-types"
import { useTranslation } from "react-i18next"

import { Button } from "@Components/generic"

const ButtonConfirm = ({
  onConfirm,
  onReject,
  children,
  confirmComponent,
  isDisabled,
}) => {
  const { t } = useTranslation()
  const [isOn, setIsOn] = React.useState(false)

  if (!children) {
    children = t("delete")
  }

  if (!confirmComponent) {
    confirmComponent = t("confirm")
  }

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
    <Button onClick={onClick} isDisabled={isDisabled} onBlur={onBlur}>
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
