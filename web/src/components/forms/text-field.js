import React from "react"
import PropTypes from "prop-types"
import { useFormContext } from "react-hook-form"

import get from "lodash/get"

const TextField = ({ name, type = "text", label, ...props }) => {
  const { register, errors } = useFormContext() // retrieve all hook methods
  const errorMsg = get(errors, [name, "message"])
  return (
    <div>
      {!!label && <div>{label}</div>}
      <div>
        <input type={type} name={name} ref={register} {...props} />
      </div>
      {!!errorMsg && <div>{errorMsg}</div>}
    </div>
  )
}

TextField.propTypes = {
  name: PropTypes.string.isRequired,
  type: PropTypes.string,
  label: PropTypes.any,
}

export default TextField
