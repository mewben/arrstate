import React from "react"
import PropTypes from "prop-types"
import { useFormContext } from "react-hook-form"
import cx from "clsx"

import { random } from "@Utils/lodash"
import InputWrapper from "./input-wrapper"

export const BaseTextField = ({
  name,
  id,
  type = "text",
  inputClassName,
  endAddon,
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
        "form-input",
        !!endAddon ? "rounded-none rounded-l-md" : "",
        inputClassName
      )}
      {...props}
    />
  )
}

const TextField = ({ name, ...props }) => {
  const id = `${name}.${random(1, 100)}`

  return (
    <InputWrapper name={name} id={id} {...props}>
      <BaseTextField name={name} id={id} {...props} />
    </InputWrapper>
  )
}

TextField.propTypes = {
  name: PropTypes.string.isRequired,
  type: PropTypes.string,
  label: PropTypes.any,
}

export default TextField
