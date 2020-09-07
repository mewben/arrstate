import React from "react"
import { useTranslation } from "react-i18next"

import { Portal, Button } from "@Components/generic"
import { PersonForm } from "@Components/popups/person"
import { AppBar } from "@Wrappers/layout"
import { List } from "./components"

const PersonList = () => {
  const { t } = useTranslation()
  return (
    <>
      <AppBar title={t("people.title")}>
        <Portal openByClickOn={<Button>{t("people.new")}</Button>}>
          <PersonForm />
        </Portal>
      </AppBar>
      <List />
    </>
  )
}

export default PersonList
