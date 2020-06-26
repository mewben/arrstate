import React from "react"
import PropTypes from "prop-types"
import { useFormContext } from "react-hook-form"

import { get } from "@Utils/lodash"
import FormLabel from "./form-label"

export const DisconnectedTextField = ({ type = "text", label, ...props }) => {
  return (
    <div>
      <FormLabel label={label} />
      <div>
        <input type={type} {...props} />
      </div>
    </div>
  )
}

const TextField = ({ name, type = "text", label, ...props }) => {
  const { register, errors } = useFormContext() // retrieve all hook methods
  const errorMsg = get(errors, [name, "message"])
  return (
    <div>
      <FormLabel label={label} />
      <div>
        <input type={type} name={name} ref={register} {...props} />
      </div>
      {!!errorMsg && <div style={{ color: "red" }}>{errorMsg}</div>}
    </div>
  )
}

TextField.propTypes = {
  name: PropTypes.string.isRequired,
  type: PropTypes.string,
  label: PropTypes.any,
}

export default TextField
