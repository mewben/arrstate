import React from "react"
import { useTranslation } from "react-i18next"
import { useMutation, queryCache } from "react-query"

import { Empty, Portal, Button, Table, TBody, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { UploadFileButton } from "@Components/files"
import { requestApi } from "@Utils"
import { map } from "@Utils/lodash"
import { useFiles } from "@Hooks"
import ListItem from "./list-item"

const List = ({ entityType, entityID }) => {
  const { t } = useTranslation()
  const getMethodParams = { entityType, entityID }

  const [upload, { reset, error }] = useMutation(formData => {
    return requestApi(`/api/files`, "POST", { data: formData })
  })

  const onUpload = data => {
    reset()
    upload(data)
  }

  const onUploadComplete = result => {
    queryCache.invalidateQueries(["files", getMethodParams])
  }

  const renderAdd = () => {
    return (
      <UploadFileButton
        entityType={entityType}
        entityID={entityID}
        onUpload={onUpload}
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
