import React from "react"

import Filters from "./components/filters"

const Income = () => {
  const onSubmit = formData => {
    console.log("formData", formData)
    const format = "yyyy-LL-dd"
    const payload = {
      ...formData,
      range: [
        formData.range[0].toFormat(format),
        formData.range[1].toFormat(format),
      ],
    }
    console.log("payload", payload)
  }
  return (
    <div>
      <Filters onSubmit={onSubmit} />
      results
    </div>
  )
}

export default Income
