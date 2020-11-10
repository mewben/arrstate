import React from "react"
import acc from "accounting"
import { Axis, Tooltip, LineAdvance } from "bizcharts"
import { Panel } from "@Components/generic"
import { Chart } from "@Components/chart"
import { startOfMonth, endOfMonth, addToDate, format } from "@Utils/date"
import { fromMoney } from "@Utils/money"
import { isEmpty } from "@Utils/lodash"

const SalesWidget = ({ data = [] }) => {
  if (isEmpty(data)) {
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
      amount: 0,
      count: 0,
      ...mapData[id],
    })
    startMonth = addToDate(startMonth, "1", "days")
  }

  const scale = {
    amount: {
      min: 0,
    },
  }

  return (
    <div className="col-span-12">
      <Panel withPadding>
        <h1 className="font-medium text-base pb-8">Sales</h1>
        <Chart autoFit height={300} data={prep} scale={scale}>
          <LineAdvance shape="smooth" point area position="dateLabel*amount" />
          <Axis
            name="amount"
            label={{ formatter: val => acc.formatNumber(val) }}
          />
          <Tooltip showMarkers showCrosshairs>
            {(title, items) => {
              return (
                <div className="p-2 text-cool-gray-900 leading-4">
                  <h4 className="text-cool-gray-500">{title}</h4>
                  <div className="text-green-600 font-medium text-sm">
                    Php {acc.formatNumber(items[0]?.data?.amount)}
                  </div>
                  <div className="font-medium">
                    {acc.formatNumber(items[0]?.data?.count)} Sales
                  </div>
                </div>
              )
            }}
          </Tooltip>
        </Chart>
      </Panel>
    </div>
  )
}

export default SalesWidget
