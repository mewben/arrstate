import React from "react"
import { Router } from "@reach/router"
import { useTranslation } from "react-i18next"

import { Loading, Error, Portal, Button } from "@Components/generic"
import { useProject } from "@Hooks"
import { AppBar, SubMenu, SubMenuItem } from "@Wrappers/layout"
import { ProjectForm } from "@Components/popups/project"
import { map } from "@Utils/lodash"
// import { Header, SubMenu } from "./components"
import ProjectOverview from "./project-overview"
import ProjectProperties from "./project-properties"

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
  ]

  const renderSubmenu = () => {
    return (
      <SubMenu>
        {map(submenu, (item, i) => {
          return (
            <SubMenuItem key={i} to={item.path}>
              {item.label}
            </SubMenuItem>
          )
        })}
      </SubMenu>
    )
  }

  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : (
    <>
      <AppBar title={data.name} backTo="/projects" submenu={renderSubmenu()}>
        <Portal openByClickOn={<Button>{t("projects.edit")}</Button>}>
          <ProjectForm model={data} />
        </Portal>
      </AppBar>
      <Router className="flex-1 overflow-y-scroll pb-28">
        <ProjectOverview path="/" project={data} />
        <ProjectProperties path="/properties" project={data} />
      </Router>
    </>
  )
}

export default ProjectSingle
