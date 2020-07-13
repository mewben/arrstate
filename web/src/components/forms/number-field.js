import React from "react"
import { Controller, useFormContext } from "react-hook-form"
import NumberFormat from "react-number-format"
import cx from "clsx"

import { random } from "@Utils/lodash"
import InputWrapper from "./input-wrapper"

const MaskedNumberInput = ({
  value,
  onChange,
  endAddon,
  inputClassName,
  max = 9999999,
  ...props
}) => {
  return (
    <NumberFormat
      type="number"
      onValueChange={({ floatValue }) => onChange(floatValue)}
      value={value}
      thousandSeparator=","
      className={cx(
        "form-input relative block w-full bg-transparent focus:z-10 transition ease-in-out duration-150 sm:text-sm sm:leading-5",
        !!endAddon ? "rounded-none rounded-l-md" : "",
        inputClassName
      )}
      max={max}
      {...props}
    />
  )
}

const NumberField = ({ name, ...props }) => {
  const id = `${name}.${random(1, 100)}`

  return (
    <InputWrapper name={name} id={id} {...props}>
      <Controller
        name={name}
        render={({ onChange, value }) => {
          return (
            <MaskedNumberInput
              value={value}
              onChange={onChange}
              id={id}
              {...props}
            />
          )
        }}
      />
    </InputWrapper>
  )
}

const NumberField2 = ({ name, ...props }) => {
  return (
    <div>
      <Controller
        name={name}
        render={({ onChange, value }) => {
          console.log("vvv", value, typeof value)

          const handleChange = val => {
            // const val = e.target.value
            console.log("typeof", typeof val)
            onChange(val)
            // onChange(parseFloat(val))
            // onChange(val !== "" ? parseFloat(val) : null)
          }

          return (
            <MaskedNumberInput
              value={value}
              onChange={handleChange}
              {...props}
            />
          )
          // const handleChange = e => {
          //   const val = e.target.value
          //   console.log("typeof", typeof val)
          //   onChange(val !== "" ? parseFloat(val) : null)
          // }
          // return <input value={value} onChange={handleChange} {...props} />
        }}
      />
    </div>
  )
}

export default NumberField
