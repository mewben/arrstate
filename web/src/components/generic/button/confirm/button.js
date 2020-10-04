import React from "react"
import PropTypes from "prop-types"
import { useTranslation } from "react-i18next"

import { Button } from "@Components/generic"
import { Confirm } from "./confirm"

export const ConfirmButton = ({
  button,
  confirmButton,
  isDelete,
  onSubmit,
  children,
  redirectOnSuccess,
  onSuccess,
  icon,
  confirmIcon,
  confirmText,
  successMessage,
  ...props
}) => {
  const { t } = useTranslation()
  const defaultButton = <Button color="red">{t("delete")}</Button>
  const defaultConfirmButton = <Button color="red">{t("confirm")}</Button>
  return (
    <Confirm
      button={button || defaultButton}
      confirmButton={confirmButton || defaultConfirmButton}
      onSubmit={onSubmit}
      redirectOnSuccess={redirectOnSuccess}
      onSuccess={onSuccess}
      successMessage={successMessage}
      {...props}
    />
  )
}

ConfirmButton.propTypes = {
  isDelete: PropTypes.bool,
  redirectOnSuccess: PropTypes.string,
  onSuccess: PropTypes.func,
  confirmText: PropTypes.string,
  successMessage: PropTypes.string,
  icon: PropTypes.string,
  confirmIcon: PropTypes.string,
}
