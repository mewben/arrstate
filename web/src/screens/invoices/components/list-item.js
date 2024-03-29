import React from "react"
import { Link } from "gatsby"
import acc from "accounting"

import { useProperty } from "@Hooks"
import { Time, Td } from "@Components/generic"
import { fromMoney } from "@Utils/money"
import { PropertyWrapper, PersonWrapper } from "@Wrappers"
import { fullName } from "@Utils"
import Status from "./status"

// Properties listitem
const ListItem = ({ item, propertyID }) => {
  return (
    <tr>
      <Td wrap>
        <Link
          to={`/invoices/${item._id}`}
          className="font-medium text-cool-gray-700 hover:text-blue-500"
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
                <div className="text-gray-700">{fullName(person?.name)}</div>
              )
            }}
          </PersonWrapper>
        )}
      </Td>
      <Td>{!!item?.issueDate && <Time d={item?.issueDate} dateOnly />}</Td>
      <Td>{!!item?.dueDate && <Time d={item?.dueDate} dateOnly />}</Td>
      <Td align="right">{acc.formatNumber(fromMoney(item.total), 2)}</Td>
      <Td>
        <Status status={item.status} />
      </Td>
    </tr>
  )
}

const Project = ({ id }) => {
  const { status, data } = useProperty(id)
  if (status !== "success") return null
  return (
    <Link to={`/projects/${id}`} className="text-gray-700 hover:text-blue-500">
      {data.name}
    </Link>
  )
}

export default ListItem
