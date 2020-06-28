import React from "react"
import PropTypes from "prop-types"
import { useFormContext } from "react-hook-form"

import { Button } from "@Components/generic"

const SubmitButton = ({ children, size, fullWidth }) => {
  const { formState } = useFormContext()
  const { isSubmitting } = formState
  return (
    <Button
      type="submit"
      isDisabled={isSubmitting}
      size={size}
      fullWidth={fullWidth}
    >
      {children}
    </Button>
  )
}

SubmitButton.propTypes = {
  children: PropTypes.any.isRequired,
}

export default SubmitButton
