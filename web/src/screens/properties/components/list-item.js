import React from "react"
import { Link } from "gatsby"
import acc from "accounting"

import { Td } from "@Components/generic"
import { t } from "@Utils/t"
import { fromMoney } from "@Utils/money"
import { ProjectWrapper } from "@Wrappers"
import Status from "./status"

// Properties listitem
const ListItem = ({ item, projectID }) => {
  return (
    <tr>
      <Td wrap>
        <Link
          to={`/properties/${item._id}`}
          className="font-medium text-cool-gray-700 hover:text-blue-500"
        >
          {item.name}
        </Link>
      </Td>
      <Td>
        <span>{t(`${item.type}`)}</span>
      </Td>
      {!projectID && (
        <Td>
          {!!item.projectID && (
            <ProjectWrapper projectID={item.projectID}>
              {({ project }) => {
                return (
                  <Link
                    to={`/projects/${project._id}`}
                    className="text-gray-700 hover:text-blue-500"
                  >
                    {project.name}
                  </Link>
                )
              }}
            </ProjectWrapper>
          )}
        </Td>
      )}
      <Td align="right">{acc.formatNumber(item.area, 2)}</Td>
      <Td align="right">{acc.formatNumber(fromMoney(item.price), 2)}</Td>
      <Td align="right">{acc.formatNumber(fromMoney(item.priceAddon), 2)}</Td>
      <Td>
        <Status status={item.status} />
      </Td>
    </tr>
  )
}

export default ListItem
