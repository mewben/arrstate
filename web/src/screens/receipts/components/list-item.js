import React from "react"
import { Link } from "gatsby"
import acc from "accounting"
import dayjs from "dayjs"

import { useProperty } from "@Hooks"
import { Td, Time } from "@Components/generic"
import { t } from "@Utils/t"
import { fromMoney } from "@Utils/money"
import { PropertyWrapper, PersonWrapper } from "@Wrappers"
import { fullName } from "@Utils"

// Receipts listitem
const ListItem = ({ item, propertyID }) => {
  return (
    <tr>
      <Td wrap>
        <Link
          to={`/receipts/${item._id}`}
          className="font-medium text-cool-gray-700 hover:text-blue-500"
        >
          {item.receiptNo}
        </Link>
        <br />
        <Link
          to={`/invoices/${item._id}`}
          className="text-cool-gray-700 hover:text-blue-500"
        >
          {item.name}
        </Link>
      </Td>
      {!!propertyID && (
        <Td>
          {!!item.propertyID && (
            <PropertyWrapper propertyID={item.propertyID}>
              {({ property }) => {
                return (
                  <Link
                    to={`/properties/${property._id}`}
                    className="text-gray-700 hover:text-blue-500"
                  >
                    {property.name}
                  </Link>
                )
              }}
            </PropertyWrapper>
          )}
        </Td>
      )}
      <Td>
        {!!item?.to?._id && (
          <PersonWrapper personID={item?.to?._id}>
            {({ person }) => {
              return (
                <div className="text-gray-700">
                  {fullName(person?.givenName, person?.familyName)}
                </div>
              )
            }}
          </PersonWrapper>
        )}
      </Td>
      <Td>
        <Time d={item?.issueDate} dateOnly />
      </Td>
      <Td>
        <Time d={item?.paidAt} dateOnly />
      </Td>
      <Td align="right">
        <div className="text-green-500 font-medium">
          {acc.formatNumber(fromMoney(item.total), 2)}
        </div>
      </Td>
    </tr>
  )
}

export default ListItem
