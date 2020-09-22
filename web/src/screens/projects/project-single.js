import React from "react"
import { Router } from "@reach/router"
import { useTranslation } from "react-i18next"

import { Loading, Error, Portal, Button } from "@Components/generic"
import { useProject } from "@Hooks"
import { AppBar } from "@Wrappers/layout"
import { ProjectForm } from "@Components/popups/project"
// import { Header, SubMenu } from "./components"
import ProjectOverview from "./project-overview"
import ProjectProperties from "./project-properties"
import ProjectFiles from "./project-files"

const ProjectSingle = ({ projectID }) => {
  const { t } = useTranslation()
  const { status, data, error } = useProject(projectID)

  const submenu = [
    {
      label: t("projects.menu.overview"),
      path: `/projects/${projectID}`,
    },
    {
      label: t("projects.menu.properties"),
      path: `/projects/${projectID}/properties`,
    },
    {
      label: t("projects.menu.clients"),
      path: `/projects/${projectID}/clients`,
    },
    {
      label: t("projects.menu.agents"),
      path: `/projects/${projectID}/agents`,
    },
    {
      label: t("projects.menu.files"),
      path: `/projects/${projectID}/files`,
    },
  ]

  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : (
    <>
      <AppBar title={data.name} backTo="/projects" submenu={submenu}>
        <Portal openByClickOn={<Button>{t("projects.edit")}</Button>}>
          <ProjectForm model={data} />
        </Portal>
      </AppBar>
      <Router className="flex-1 overflow-y-scroll pb-28">
        <ProjectOverview path="/" project={data} />
        <ProjectProperties path="/properties" project={data} />
        <ProjectFiles path="/files" project={data} />
      </Router>
    </>
  )
}

export default ProjectSingle
