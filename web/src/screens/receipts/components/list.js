import React from "react"
import { useTranslation } from "react-i18next"

import { Empty, Table, TBody, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useInvoices } from "@Hooks"
import { INVOICE_STATUS } from "@Enums"
import ListItem from "./list-item"

const List = ({ propertyID }) => {
  const { t } = useTranslation()

  const renderContent = content => {
    return (
      <div className="p-4">
        <Table>
          <thead>
            <tr>
              <Th fullWidth>{t("receipts.no")}</Th>
              <Th>{t("receipts.property")}</Th>
              <Th>{t("receipts.issuedTo")}</Th>
              <Th>{t("receipts.issuedDate")}</Th>
              <Th>{t("receipts.paidDate")}</Th>
              <Th align="right">{t("receipts.total")} (Php)</Th>
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

  const getMethodParams = {
    status: INVOICE_STATUS.PAID,
  }
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
