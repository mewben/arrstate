import React from "react"

import { LANGUAGES } from "@Enums/languages"
import { map } from "@Utils/lodash"
import SelectField from "../select-field"

const LanguageSelectField = ({ ...props }) => {
  const options = React.useMemo(() => {
    return map(LANGUAGES, (item, k) => {
      return {
        value: k,
        label: item.nativeName,
      }
    })
  }, [])

  return <SelectField options={options} {...props} />
}

export default LanguageSelectField
