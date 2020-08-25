import React from "react"
import PropTypes from "prop-types"
import { Controller } from "react-hook-form"
import TextField from "@material-ui/core/TextField"
import Paper from "@material-ui/core/Paper"
import Autocomplete from "@material-ui/lab/Autocomplete"
import UnfoldMoreIcon from "@material-ui/icons/UnfoldMore"
import cx from "clsx"

import { keyBy, random, isObject, map } from "@Utils/lodash"
import { FieldLabel, FieldError } from "../field"

const SelectField = ({
  name,
  label,
  leftLabel,
  hint,
  options = [],
  placeholder,
  disableClearable,
  selectOnFocus = true,
  multiple,
  containerClass = "col-span-12",
  ...props
}) => {
  const { optionsMap, id } = React.useMemo(() => {
    return {
      optionsMap: keyBy(options, "value"),
      id: `${name}.${random(0, 1000)}`,
    }
  }, [options])

  const cxContainer = cx(
    containerClass,
    leftLabel ? "flex space-x-4" : "",
    hint ? "items-start" : "items-center"
  )
  const cxInput = cx(leftLabel ? "sm:w-2/3" : "")

  return (
    <div className={cxContainer}>
      <FieldLabel label={label} leftLabel={leftLabel} id={id} hint={hint} />
      <div className={cxInput}>
        <Controller
          name={name}
          defaultValue={multiple ? [] : null}
          render={({ onChange, value }) => {
            return (
              <Autocomplete
                id={id}
                value={value}
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
                        className="form-input relative bg-transparent focus:z-10 transition ease-in-out duration-150 sm:text-sm sm:leading-5"
                        placeholder={placeholder}
                        InputProps={{
                          ...params.InputProps,
                          disableUnderline: true,
                        }}
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
                // forcePopupIcon={false}
                PaperComponent={CustomPaper}
                onChange={(e, option, reason) => {
                  if (!multiple) {
                    onChange(option?.value || null)
                  } else {
                    onChange(
                      map(option, opt => {
                        if (isObject(opt)) {
                          return opt.value
                        } else {
                          return opt
                        }
                      })
                    )
                  }
                }}
                {...props}
              />
            )
          }}
        />
        <FieldError name={name} />
      </div>
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

const CustomPaper = ({ children }) => {
  return (
    <Paper
      variant="outlined"
      className="text-sm mt-1 rounded-md bg-white shadow-lg"
    >
      {children}
    </Paper>
  )
}

SelectField.propTypes = {
  name: PropTypes.string.isRequired,
  options: PropTypes.array,
  label: PropTypes.string,
}

export default SelectField
