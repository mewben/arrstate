import React from "react"

import { Loading, Error } from "@Components/generic"
import { useCountries } from "@Hooks"
import { map } from "@Utils/lodash"
import { t } from "@Utils/t"
import SelectField from "../select-field"

const CountriesSelectField = ({ ...props }) => {
  const { status, data, error } = useCountries()

  const options = React.useMemo(() => {
    return map(data?.countries, item => {
      return {
        value: item,
        label: t(`countries.${item}`),
      }
    })
  }, [data?.countries])

  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : (
    <SelectField options={options} {...props} />
  )
}

export default CountriesSelectField
