import React from "react"
import { Link } from "gatsby"
import { Router } from "@reach/router"

import { Loading, Error, Portal, Button } from "@Components/generic"
import { useProject } from "@Hooks"
import { AppBar, SubMenu, SubMenuItem } from "@Wrappers/layout"
import { ProjectForm } from "@Components/popups/project"
import { map } from "@Utils/lodash"
// import { Header, SubMenu } from "./components"
import ProjectOverview from "./project-overview"
import ProjectProperties from "./project-properties"

const ProjectSingle = ({ projectID }) => {
  const { status, data, error } = useProject(projectID)

  const submenu = [
    {
      label: "Overview",
      path: `/projects/${projectID}`,
    },
    {
      label: "Properties",
      path: `/projects/${projectID}/properties`,
    },
    {
      label: "Clients",
      path: `/projects/${projectID}/clients`,
    },
    {
      label: "Agents",
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
        <Portal openByClickOn={<Button>Edit Project</Button>}>
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
