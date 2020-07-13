import React from "react"
import PropTypes from "prop-types"
import { useForm, FormProvider } from "react-hook-form"
import { yupResolver } from "@hookform/resolvers"

const Form = ({ model, validationSchema, onSubmit, children, ...props }) => {
  const methods = useForm({
    defaultValues: model,
    resolver: yupResolver(validationSchema),
    ...props,
  })

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)} noValidate>
        {children}
      </form>
    </FormProvider>
  )
}

Form.propTypes = {
  onSubmit: PropTypes.func.isRequired,
  model: PropTypes.object,
  validationSchema: PropTypes.object,
}

export default Form
