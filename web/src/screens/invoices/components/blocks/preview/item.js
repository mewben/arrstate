import React from "react"
import acc from "accounting"

import { Table, Th, Td } from "@Components/generic"
import { fromMoney } from "@Utils"

export const ItemWrapper = ({ children }) => {
  return (
    <div className="py-8">
      <Table className="border border-gray-200 shadow-none">
        <thead>
          <tr>
            <Th fullWidth>Item</Th>
            <Th align="right">Amount (Php)</Th>
            <Th align="right">Quantity</Th>
            <Th align="right">Tax</Th>
            <Th align="right">Discount</Th>
            <Th align="right">Total (Php)</Th>
          </tr>
        </thead>
        <tbody>{children}</tbody>
      </Table>
    </div>
  )
}

const Item = ({ block, ...props }) => {
  return (
    <tr>
      <Td className="max-w-full" wrap>
        <div className="text-gray-800 font-medium">{block.title}</div>
        <div>{block.description}</div>
      </Td>
      <Td align="right">{acc.formatNumber(fromMoney(block.amount), 2)}</Td>
      <Td align="right">{block.quantity}</Td>
      <Td align="right">{block.tax ? `${block.tax}%` : 0}</Td>
      <Td align="right">{block.discount || 0}</Td>
      <Td align="right" className="text-gray-900 font-medium">
        {acc.formatNumber(fromMoney(block.total), 2)}
      </Td>
    </tr>
  )
}

export default Item
