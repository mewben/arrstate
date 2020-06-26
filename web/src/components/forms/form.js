import React from "react"
import PropTypes from "prop-types"
import { useForm, FormContext } from "react-hook-form"

const Form = ({ model, validationSchema, onSubmit, children, ...props }) => {
  const methods = useForm({
    defaultValues: model,
    validationSchema,
    ...props,
  })

  return (
    <FormContext {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)}>{children}</form>
    </FormContext>
  )
}

Form.propTypes = {
  onSubmit: PropTypes.func.isRequired,
  model: PropTypes.object,
  validationSchema: PropTypes.object,
}

export default Form
