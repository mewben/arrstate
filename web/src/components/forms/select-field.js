import React from "react"
import PropTypes from "prop-types"
import { Controller } from "react-hook-form"
import TextField from "@material-ui/core/TextField"
import Autocomplete from "@material-ui/lab/Autocomplete"
import UnfoldMoreIcon from "@material-ui/icons/UnfoldMore"

import { FieldLabel, FieldError } from "./field"
import { keyBy, random, isObject, map } from "@Utils/lodash"

const SelectField = ({
  name,
  label,
  options = [],
  placeholder,
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
            renderInput={params => {
              return (
                <div className="mt-1 flex relative rounded-md shadow-sm">
                  <TextField
                    {...params}
                    className="form-select relative bg-transparent focus:z-10 transition ease-in-out duration-150 sm:text-sm sm:leading-5"
                    placeholder={placeholder}
                    InputProps={{
                      ...params.InputProps,
                      disableUnderline: true,
                    }}
                    disableUnderline
                  />
                </div>
              )
            }}
            renderTags={(value, getTagProps) => {
              return value.map((option, index) => {
                return <Tag label={option} {...getTagProps({ index })} />
              })
            }}
            popupIcon={<UnfoldMoreIcon fontSize="small" />}
            forcePopupIcon={false}
          />
        }
        onChange={([, option]) => {
          if (!multiple) {
            return option?.value
          } else {
            return map(option, opt => {
              if (isObject(opt)) {
                return opt.value
              } else {
                return opt
              }
            })
          }
        }}
        name={name}
        defaultValue={multiple ? [] : null}
      />
      <FieldError name={name} />
    </div>
  )
}

const Tag = ({ label, onDelete, ...props }) => {
  return (
    <span
      {...props}
      className="flex items-center px-2.5 py-1.5 rounded-sm text-sm font-medium leading-4 bg-gray-100 border border-gray-200 text-gray-800 m-0.5"
    >
      {label}
      <button
        type="button"
        onClick={onDelete}
        className="flex-shrink-0 ml-1.5 inline-flex text-gray-400 p-1 focus:outline-none focus:text-gray-800 hover:text-gray-800"
        aria-label="Remove small badge"
      >
        <svg
          className="h-2 w-2"
          stroke="currentColor"
          fill="none"
          viewBox="0 0 8 8"
        >
          <path strokeLinecap="round" strokeWidth="1.5" d="M1 1l6 6m0-6L1 7" />
        </svg>
      </button>
    </span>
  )
}

SelectField.propTypes = {
  name: PropTypes.string.isRequired,
  options: PropTypes.array,
  label: PropTypes.string,
}

export default SelectField
