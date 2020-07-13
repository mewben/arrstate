import React from "react"
import cx from "clsx"

import { map } from "@Utils/lodash"
import InputWrapper from "./input-wrapper"

const InputGroup = ({ name, id, label, children }) => {
  return (
    <InputWrapper name={name} label={label} id={id}>
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
  )
}

export default InputGroup
