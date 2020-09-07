import React from "react"

import { Loading, Error } from "@Components/generic"
import { usePeople } from "@Hooks"
import { map, sortBy } from "@Utils/lodash"
import { fullName } from "@Utils"
import SelectField from "../select-field"

const PeopleSelectField = ({ role = [], ...props }) => {
  const { status, data, error, isFetching } = usePeople(role)

  const options = React.useMemo(() => {
    const items = map(data?.list, item => {
      return {
        value: item._id,
        label: fullName(item.name),
      }
    })
    return sortBy(items, "label")
  }, [data?.list])

  return status === "loading" || isFetching ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : (
    <SelectField options={options} {...props} />
  )
}

export default PeopleSelectField
