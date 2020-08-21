import React from "react"
import { useTranslation } from "react-i18next"

import { Portal, Button } from "@Components/generic"
import { ProjectForm } from "@Components/popups/project"
import { AppBar } from "@Wrappers/layout"
import { List } from "./components"

const ProjectList = () => {
  const { t } = useTranslation()

  return (
    <>
      <AppBar title={t("projects.title")}>
        <Portal openByClickOn={<Button>{t("projects.new")}</Button>}>
          <ProjectForm />
        </Portal>
      </AppBar>
      <List />
    </>
  )
}

export default ProjectList
