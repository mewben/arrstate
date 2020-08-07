import React from "react"
import acc from "accounting"

import { Table, Th, Td } from "@Components/generic"
import { fromMoney } from "@Utils"
import { map } from "@Utils/lodash"

const Summary = ({ block, invoice, isReceipt }) => {
  return (
    <div className="flex w-full justify-end pb-8">
      <div className="w-1/2">
        <Table className="border border-gray-200">
          <tbody>
            <tr>
              <Td py="py-3">SubTotal:</Td>
              <Td align="right" py="py-3">
                Php {acc.formatNumber(fromMoney(invoice.subTotal), 2)}
              </Td>
            </tr>
            {map(invoice?.addOrLess, item => {
              return (
                <tr key={item._id}>
                  <Td py="py-3">{item.name}:</Td>
                  <Td align="right" py="py-3">
                    {item.value}
                  </Td>
                </tr>
              )
            })}
          </tbody>
          <tfoot>
            <tr>
              <Th fullWidth>
                <div className="text-sm font-medium text-green-500">
                  {isReceipt ? "Amount Paid:" : "Amount Due:"}
                </div>
              </Th>
              <Th align="right">
                <div className="text-sm font-bold text-green-500">
                  Php {acc.formatNumber(fromMoney(invoice.total), 2)}
                </div>
              </Th>
            </tr>
          </tfoot>
        </Table>
      </div>
    </div>
  )
}

export default Summary
