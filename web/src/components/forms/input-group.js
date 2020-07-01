import React from "react"
import cx from "clsx"

import { map } from "@Utils/lodash"
import { FieldLabel, InputWrapper } from "./field"

const InputGroup = ({ label, children }) => {
  return (
    <>
      <FieldLabel label={label} />
      <InputWrapper>
        <div className="flex w-full">
          {map(children, (child, i) => {
            return (
              <div
                key={i}
                className={cx("flex-1 min-w-0", {
                  "-ml-px": i !== 0,
                })}
              >
                {child}
              </div>
            )
          })}
        </div>
      </InputWrapper>
    </>
  )
}

export default InputGroup
