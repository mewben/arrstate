import React from "react"
import PropTypes from "prop-types"
import { Controller } from "react-hook-form"
import Autocomplete from "@material-ui/lab/Autocomplete"

import { FieldLabel, FieldError } from "./field"
import { keyBy, random, isObject } from "@Utils/lodash"

const SelectField = ({
  name,
  label,
  options = [],
  disableClearable,
  selectOnFocus = true,
  multiple,
  ...props
}) => {
  const optionsMap = React.useMemo(() => {
    return keyBy(options, "value")
  }, [options])

  return (
    <div>
      <FieldLabel label={label} />
      <Controller
        as={
          <Autocomplete
            id={`${name}.${random(0, 1000)}`}
            options={options}
            multiple={multiple}
            disableClearable={disableClearable}
            selectOnFocus={selectOnFocus}
            size="small"
            getOptionLabel={option =>
              isObject(option) ? option.label : optionsMap[option]?.label
            }
            getOptionSelected={(option, value) => option.value === value}
            renderInput={params => (
              <div
                className="mt-1 flex relative rounded-md shadow-sm"
                ref={params.InputProps.ref}
              >
                <input
                  type="text"
                  {...params.inputProps}
                  className="form-input relative block w-full bg-transparent focus:z-10 transition ease-in-out duration-150 sm:text-sm sm:leading-5"
                />
              </div>
            )}
          />
        }
        onChange={([, option]) => option?.value}
        name={name}
        defaultValue={multiple ? [] : null}
      />
      <FieldError name={name} />
    </div>
  )
}

SelectField.propTypes = {
  name: PropTypes.string.isRequired,
  options: PropTypes.array,
  label: PropTypes.string,
}

export default SelectField
