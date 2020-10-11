import React from "react"
import { useTranslation } from "react-i18next"
import { queryCache } from "react-query"

import { Empty, Portal, Button, Table, TBody, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { UploadFileButton } from "@Components/files"
import { map } from "@Utils/lodash"
import { useFiles } from "@Hooks"
import ListItem from "./list-item"

const List = ({ entityType, entityID }) => {
  const { t } = useTranslation()
  const getMethodParams = { entityType, entityID }

  const onUploadComplete = (result, cb) => {
    queryCache.invalidateQueries(["files", getMethodParams])
    cb()
  }

  const renderAdd = () => {
    return (
      <UploadFileButton
        entityType={entityType}
        entityID={entityID}
        onUploadComplete={onUploadComplete}
      >
        <Button>{t("files.add")}</Button>
      </UploadFileButton>
    )
  }

  const renderContent = content => {
    return (
      <div className="p-4">
        <div>{renderAdd()}</div>
        <Table>
          <thead>
            <tr>
              <Th fullWidth>{t("files.filename")}</Th>
              <Th>{t("files.type")}</Th>
              <Th align="right">{t("files.size")}</Th>
              <Th> </Th>
            </tr>
          </thead>
          <TBody>
            {map(content?.list, item => {
              return <ListItem key={item._id} item={item} />
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

  return (
    <InfiniteScroll
      getMethod={useFiles}
      getMethodParams={getMethodParams}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
