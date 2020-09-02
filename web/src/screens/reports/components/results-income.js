import React from "react"
import { useTranslation } from "react-i18next"
import { Empty, Table, TBody, Th, Td, Time } from "@Components/generic"
import acc from "accounting"

import { map } from "@Utils/lodash"
import { fromMoney } from "@Utils/money"

const ResultsIncome = ({ results }) => {
  const { t } = useTranslation()
  console.log("results", results)
  if (!results?.list?.length) {
    return null
  }

  return (
    <div className="p-4">
      <Table>
        <thead>
          <tr>
            <Th fullWidth>{t("reports.invoiceName")}</Th>
            <Th>{t("reports.paidAt")}</Th>
            <Th align="right">{t("reports.amount")}</Th>
          </tr>
        </thead>
        <TBody>
          {map(results?.list, item => {
            return <ListItem key={item._id} item={item} />
          })}
        </TBody>
      </Table>
    </div>
  )
}

const ListItem = ({ item }) => {
  return (
    <tr>
      <Td wrap>{item.name}</Td>
      <Td>
        <Time d={item?.paidAt} />
      </Td>
      <Td align="right">{acc.formatNumber(fromMoney(item.total), 2)}</Td>
    </tr>
  )
}

export default ResultsIncome
