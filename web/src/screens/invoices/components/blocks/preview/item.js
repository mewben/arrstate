import React from "react"
import acc from "accounting"
import { useTranslation } from "react-i18next"

import { Table, useTable, Th, Td } from "@Components/generic"
import { fromMoney } from "@Utils"

export const ItemWrapper = ({ children }) => {
  return (
    <div className="py-8">
      <Table className="border border-gray-200 shadow-none">
        <ItemHead />
        <tbody>{children}</tbody>
      </Table>
    </div>
  )
}

const ItemHead = () => {
  const { t } = useTranslation()
  const { renderExtraTh } = useTable()
  return (
    <thead>
      <tr>
        <Th fullWidth>{t("blocks.item.item")}</Th>
        <Th align="right">{t("blocks.item.amount")} (Php)</Th>
        <Th align="right">{t("blocks.item.quantity")}</Th>
        {renderExtraTh()}
        <Th align="right">{t("blocks.item.total")} (Php)</Th>
      </tr>
    </thead>
  )
}

const Item = ({ block, ...props }) => {
  const { registerExtraTh, renderExtraTd } = useTable()
  React.useEffect(() => {
    registerExtraTh(block.addOrLess)
  }, [])
  return (
    <tr>
      <Td className="max-w-full" wrap>
        <div className="text-gray-800 font-medium">{block.title}</div>
        <div>{block.description}</div>
      </Td>
      <Td align="right">{acc.formatNumber(fromMoney(block.amount), 2)}</Td>
      <Td align="right">{block.quantity}</Td>
      {renderExtraTd(block.addOrLess)}
      <Td align="right" className="text-gray-900 font-medium">
        {acc.formatNumber(fromMoney(block.total), 2)}
      </Td>
    </tr>
  )
}

export default Item
