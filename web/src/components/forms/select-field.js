import React from "react"
import PropTypes from "prop-types"
import { useFormContext, Controller } from "react-hook-form"
import Autocomplete from "@material-ui/lab/Autocomplete"
import FormLabel from "./form-label"

import { get, random } from "@Utils/lodash"

const SelectField = ({
  name,
  label,
  options = [],
  disableClearable,
  selectOnFocus = true,
  multiple,
  ...props
}) => {
  const { errors } = useFormContext()
  const errorMsg = get(errors, [name, "message"])

  return (
    <div>
      <FormLabel label={label} />
      <div>
        <Controller
          as={
            <Autocomplete
              id={`${name}.${random(0, 1000)}`}
              options={options}
              multiple={multiple}
              disableClearable={disableClearable}
              selectOnFocus={selectOnFocus}
              size="small"
              getOptionLabel={option => option.label}
              renderInput={params => (
                <div ref={params.InputProps.ref}>
                  <input type="text" {...params.inputProps} />
                </div>
              )}
            />
          }
          onChange={([, option]) => option}
          name={name}
          defaultValue={multiple ? [] : null}
        />
      </div>
      {!!errorMsg && <div style={{ color: "red" }}>{errorMsg}</div>}
    </div>
  )
}

SelectField.propTypes = {
  name: PropTypes.string.isRequired,
  options: PropTypes.array,
  label: PropTypes.string,
}

export default SelectField
