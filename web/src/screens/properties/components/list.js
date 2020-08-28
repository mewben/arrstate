import React from "react"
import { useTranslation } from "react-i18next"

import { Empty, Portal, Button, Table, TBody, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useProperties } from "@Hooks"
import { PropertyForm } from "@Components/popups/property"
import ListItem from "./list-item"

const List = ({ projectID }) => {
  const { t } = useTranslation()

  const renderAdd = () => {
    if (!projectID) return null

    return (
      <Portal openByClickOn={<Button>{t("properties.add")}</Button>}>
        <PropertyForm projectID={projectID} />
      </Portal>
    )
  }

  const renderContent = content => {
    return (
      <div className="p-4">
        <div>{renderAdd()}</div>
        <Table>
          <thead>
            <tr>
              <Th fullWidth>{t("properties.code")}</Th>
              <Th>{t("properties.type")}</Th>
              {!projectID && <Th>{t("properties.project")}</Th>}
              <Th align="right">{t("properties.area")} (sq.m)</Th>
              <Th align="right">{t("properties.price")} (Php)</Th>
              <Th align="right">{t("properties.priceAddon")} (Php)</Th>
              <Th>{t("properties.status")}</Th>
            </tr>
          </thead>
          <TBody>
            {map(content?.list, item => {
              return (
                <ListItem key={item._id} item={item} projectID={projectID} />
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
        {renderAdd()}
      </div>
    )
  }

  const getMethodParams = {}
  if (projectID === null || !!projectID) {
    getMethodParams["projectID"] = projectID
  }

  return (
    <InfiniteScroll
      getMethod={useProperties}
      getMethodParams={getMethodParams}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
