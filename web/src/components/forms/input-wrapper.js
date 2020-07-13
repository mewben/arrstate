import React from "react"

import { FieldLabel, FieldError, FieldDescription } from "./field"

const InputWrapper = ({
  id,
  name,
  label,
  description,
  children,
  containerClass = "col-span-12",
  startAddon,
  startAddonInline,
  endAddon,
  endAddonInline,
}) => {
  return (
    <div className={containerClass}>
      <FieldLabel id={id} label={label} />
      <div className="mt-1 flex relative rounded-md shadow-sm">
        {!!startAddon && (
          <span className="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 sm:text-sm">
            {startAddon}
          </span>
        )}
        {!!startAddonInline && (
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <span className="text-gray-500 sm:text-sm sm:leading-5">
              {startAddonInline}
            </span>
          </div>
        )}
        {children}
        {!!endAddon && (
          <span className="inline-flex items-center px-3 rounded-r-md border border-l-0 border-gray-300 bg-gray-50 text-gray-500 sm:text-sm">
            {endAddon}
          </span>
        )}
        {!!endAddonInline && (
          <div className="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
            <span className="text-gray-500 sm:text-sm sm:leading-5">
              {endAddonInline}
            </span>
          </div>
        )}
      </div>
      <FieldDescription description={description} />
      <FieldError name={name} />
    </div>
  )
}

export default InputWrapper
