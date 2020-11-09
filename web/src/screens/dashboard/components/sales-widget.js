import React, { useState, useEffect } from "react"

import { Chart, LineAdvance } from "bizcharts"
import { Panel } from "@Components/generic"
import { startOfMonth, endOfMonth, addToDate, format } from "@Utils/date"
import { fromMoney } from "@Utils/money"
import { isEmpty } from "@Utils/lodash"

const SalesWidget = ({ data = [] }) => {
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  if (isEmpty(data)) {
    return null
  }

  if (!mounted) {
    return null
  }

  // prepare data
  // transform data into map
  const mapData = {}
  data.forEach(({ _id, amount, count }) => {
    const key = `${_id.year}-${_id.month}-${_id.day}`
    mapData[key] = { amount: fromMoney(amount), count }
  })

  const now = new Date()
  let startMonth = startOfMonth(now)
  const endMonth = endOfMonth(now)

  const prep = []
  const idFormat = "yyyy-M-d"
  const labelFormat = "LLL d"
  while (startMonth < endMonth) {
    const id = format(startMonth, idFormat)
    prep.push({
      dateLabel: format(startMonth, labelFormat),
      id: format(startMonth, idFormat),
      ...mapData[id],
    })
    startMonth = addToDate(startMonth, "1", "days")
  }

  return (
    <div className="col-span-12">
      <Panel withPadding>
        <h1 className="font-medium pb-8">Sales</h1>
        <Chart autoFit height={300} data={prep}>
          <LineAdvance
            shape="smooth"
            point
            area
            //position="month*temperature"
            position="dateLabel*amount"
            // color="count"
          />
        </Chart>
      </Panel>
    </div>
  )
}

export default SalesWidget
