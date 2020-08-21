import React from "react"
import { useTranslation } from "react-i18next"

import { Empty, Portal, Button, Table, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { usePeople } from "@Hooks"
import { PersonForm } from "@Components/popups/people"
import ListItem from "./list-item"

const List = () => {
  const { t } = useTranslation()

  const renderAdd = () => {
    return (
      <Portal openByClickOn={<Button>{t("people.add")}</Button>}>
        <PersonForm />
      </Portal>
    )
  }

  const renderContent = content => {
    return (
      <div className="p-4">
        <Table>
          <thead>
            <tr>
              <Th fullWidth>{t("people.name")}</Th>
              <Th>{t("people.role")}</Th>
            </tr>
          </thead>
          <tbody className="bg-white">
            {map(content?.list, (item, index) => {
              return <ListItem key={item._id} item={item} index={index} />
            })}
          </tbody>
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
      getMethod={usePeople}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
