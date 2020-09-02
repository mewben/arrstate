import React from "react"
import { useMutation } from "react-query"

import { requestApi } from "@Utils"
import Filters from "./components/filters"
import ResultsIncome from "./components/results-income"

const Income = () => {
  const [results, setResults] = React.useState()
  const [fetch, { reset, isLoading, error }] = useMutation(
    params => {
      return requestApi("/api/reports/income", "GET", {
        params,
      })
    },
    {
      onSuccess: ({ data }) => setResults(data),
    }
  )

  const onSubmit = formData => {
    reset()
    const format = "LL-dd-yyyy"
    fetch({
      from: formData.range[0].toFormat(format),
      to: formData.range[1].toFormat(format),
    })
  }

  return (
    <div>
      <Filters onSubmit={onSubmit} isLoading={isLoading} />
      <ResultsIncome results={results} />
    </div>
  )
}

export default Income
