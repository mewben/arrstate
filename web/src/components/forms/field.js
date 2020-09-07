import React from "react"
import { useFormContext } from "react-hook-form"
import { get } from "@Utils/lodash"
import cx from "clsx"

export const FieldLabel = ({ id, label, leftLabel, hint }) => {
  if (!label && !leftLabel) {
    return null
  }

  const cl = cx(
    label ? "flex justify-between" : "",
    leftLabel ? "w-1/3 break-all" : ""
  )

  return (
    <div className={cl}>
      <label
        htmlFor={id}
        className="block text-sm font-medium leading-5 text-gray-700"
      >
        {label || leftLabel}
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
  const errorMessage = get(errors, `${name}.message`)
  if (!errorMessage) {
    return null
  }

  return <p className="mt-2 text-sm text-red-600">{errorMessage}</p>
}
