import React from "react"
import { Link } from "gatsby"
import acc from "accounting"
import { useTranslation } from "react-i18next"

import { Td } from "@Components/generic"
import { fromMoney } from "@Utils/money"
import { ProjectWrapper } from "@Wrappers"
import Status from "./status"

// Properties listitem
// t('properties.lot')
// t('properties.house')
const ListItem = ({ item, projectID }) => {
  const { t } = useTranslation()

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
        <span>{t(`properties.${item.type}`)}</span>
      </Td>
      {!projectID && (
        <Td>
          {!!item.projectID && (
            <ProjectWrapper
              projectID={item.projectID}
              loadingVariant="skeleton"
            >
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
