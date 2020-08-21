import React from "react"
import { useTranslation } from "react-i18next"

import { Portal, Button } from "@Components/generic"
import { PropertyForm } from "@Components/popups/property"
import { AppBar } from "@Wrappers/layout"

import { useProjects } from "@Hooks"
import { List } from "./components"

const PropertyList = () => {
  useProjects()
  const { t } = useTranslation()
  return (
    <>
      <AppBar title={t("properties.title")}>
        <Portal openByClickOn={<Button>{t("properties.new")}</Button>}>
          <PropertyForm />
        </Portal>
      </AppBar>
      <List />
    </>
  )
}

export default PropertyList
