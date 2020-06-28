import React from "react"
import { useFormContext } from "react-hook-form"
import { get } from "@Utils/lodash"

export const InputWrapper = ({ children }) => {
  return (
    <div className="mt-1 flex relative rounded-md shadow-sm">{children}</div>
  )
}

export const FieldLabel = ({ id, label, hint }) => {
  if (!label) {
    return null
  }
  return (
    <div className="flex justify-between">
      <label
        htmlFor={id}
        className="block text-sm font-medium leading-5 text-gray-700"
      >
        {label}
      </label>
      {!!hint && (
        <span className="text-sm leading-5 text-gray-500">{hint}</span>
      )}
    </div>
  )
}

export const FieldDescription = ({ description }) => {
  if (!description) {
    return null
  }

  return <p className="mt-2 text-sm text-gray-500">{description}</p>
}

export const FieldError = ({ name }) => {
  const { errors } = useFormContext()
  const errorMessage = get(errors, [name, "message"])
  if (!errorMessage) {
    return null
  }

  return <p className="mt-2 text-sm text-red-600">{errorMessage}</p>
}
