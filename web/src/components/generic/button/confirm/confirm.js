import React, { cloneElement, useRef, useState } from "react"
import { navigate } from "gatsby"
import PropTypes from "prop-types"
import { Portal, Alert } from "@Components/generic"
import { Confirm as ConfirmPopup } from "@Components/popups/generic"

export const Confirm = ({
  onSubmit,
  onSuccess,
  redirectOnSuccess,
  successMessage,
  popup,
  button,
  confirmButton,
}) => {
  const [isConfirm, setIsConfirm] = useState(false)
  const [isLoading, setIsLoading] = useState(false)
  const timeout = useRef(0)

  const hideConfirm = () => {
    setIsConfirm(false)

    if (timeout.current) {
      clearTimeout(timeout.current)
    }
  }

  const showConfirm = () => {
    setIsConfirm(true)

    timeout.current = setTimeout(hideConfirm, 5000)
  }

  const remove = () => {
    setIsLoading(true)

    if (onSubmit) {
      onSubmit((err, res) => {
        if (err) {
          // Alert.error(err)
        } else {
          if (redirectOnSuccess) {
            navigate(redirectOnSuccess, { replace: true })
          }

          if (onSuccess) {
            onSuccess(res)
          }

          if (successMessage) {
            // Alert.success(successMessage)
          }
        }

        hideConfirm()

        setIsLoading(false)
      })
    }
  }

  const content = cloneElement(!isConfirm ? button : confirmButton, {
    onClick: !isConfirm ? showConfirm : remove,
    onBlur: hideConfirm,
    isLoading,
  })

  if (!popup) {
    return content
  }

  return (
    <Portal openByClickOn={content}>
      <ConfirmPopup onConfirm={remove}>{popup}</ConfirmPopup>
    </Portal>
  )
}

Confirm.propTypes = {
  redirectOnSuccess: PropTypes.string,
  successMessage: PropTypes.string,
  onSubmit: PropTypes.func,
  onSuccess: PropTypes.func,
  popup: PropTypes.any,
  button: PropTypes.element,
  confirmButton: PropTypes.element,
}
