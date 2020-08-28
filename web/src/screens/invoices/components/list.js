import React from "react"
import { useTranslation } from "react-i18next"

import { Empty, Table, TBody, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useInvoices } from "@Hooks"
import ListItem from "./list-item"

const List = ({ propertyID }) => {
  const { t } = useTranslation()

  const renderContent = content => {
    return (
      <div className="p-4">
        <Table>
          <thead>
            <tr>
              <Th fullWidth>{t("invoices.name")}</Th>
              {!!propertyID && <Th>{t("invoices.property")}</Th>}
              <Th>{t("invoices.issuedTo")}</Th>
              <Th>{t("invoices.issueDate")}</Th>
              <Th>{t("invoices.dueDate")}</Th>
              <Th align="right">{t("invoices.total")} (Php)</Th>
              <Th>{t("invoices.status")}</Th>
            </tr>
          </thead>
          <TBody>
            {map(content?.list, item => {
              return (
                <ListItem key={item._id} item={item} propertyID={propertyID} />
              )
            })}
          </TBody>
        </Table>
      </div>
    )
  }

  const renderEmpty = () => {
    return (
      <div>
        <Empty />
      </div>
    )
  }

  const getMethodParams = {}
  if (propertyID === null || !!propertyID) {
    getMethodParams["propertyID"] = propertyID
  }

  return (
    <InfiniteScroll
      getMethod={useInvoices}
      getMethodParams={getMethodParams}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
