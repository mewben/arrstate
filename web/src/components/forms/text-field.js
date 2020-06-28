import React from "react"
import PropTypes from "prop-types"
import { useFormContext } from "react-hook-form"
import cx from "clsx"

import { random } from "@Utils/lodash"
import { InputWrapper, FieldLabel, FieldError, FieldDescription } from "./field"

export const DisconnectedTextField = ({ type = "text", label, ...props }) => {
  return (
    <div>
      <FieldLabel label={label} />
      <InputWrapper>
        <input type={type} {...props} />
      </InputWrapper>
    </div>
  )
}

export const BaseTextField = ({
  name,
  id,
  type = "text",
  className,
  hasEndAddon,
  ...props
}) => {
  const { register } = useFormContext()
  return (
    <input
      id={id}
      type={type}
      name={name}
      ref={register}
      className={cx(
        "form-input relative block w-full bg-transparent focus:z-10 transition ease-in-out duration-150 sm:text-sm sm:leading-5",
        hasEndAddon ? "rounded-none rounded-l-md" : "",
        className
      )}
      {...props}
    />
  )
}

const TextField = ({ name, type, label, description, endAddon, ...props }) => {
  const id = `${name}.${random(1, 100)}`

  return (
    <div>
      <FieldLabel id={id} label={label} />
      <InputWrapper>
        <BaseTextField
          id={id}
          name={name}
          type={type}
          hasEndAddon={!!endAddon}
          {...props}
        />
        {!!endAddon && (
          <span className="inline-flex items-center px-3 rounded-r-md border border-l-0 border-gray-300 bg-gray-50 text-gray-500 sm:text-sm">
            {endAddon}
          </span>
        )}
      </InputWrapper>
      <FieldDescription description={description} />
      <FieldError name={name} />
    </div>
  )
}

TextField.propTypes = {
  name: PropTypes.string.isRequired,
  type: PropTypes.string,
  label: PropTypes.any,
}

export default TextField
