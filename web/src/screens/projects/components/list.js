import React from "react"
import { useTranslation } from "react-i18next"

import { Empty, Portal, Button, Table, TBody, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useProjects } from "@Hooks"
import { ProjectForm } from "@Components/popups/project"
import ListItem from "./list-item"

const List = () => {
  const { t } = useTranslation()

  const renderAdd = () => {
    return (
      <Portal openByClickOn={<Button>{t("projects.add")}</Button>}>
        <ProjectForm />
      </Portal>
    )
  }

  const renderContent = content => {
    return (
      <div className="p-4">
        <Table>
          <thead>
            <tr>
              <Th fullWidth>{t("projects.name")}</Th>
              <Th align="right">{t("projects.area")}</Th>
            </tr>
          </thead>
          <TBody>
            {map(content?.list, (item, index) => {
              return <ListItem key={item._id} item={item} index={index} />
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
      getMethod={useProjects}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
