import React from "react"
import { ProjectForm, List } from "./components"

const ProjectList = () => {
  const [isOpenForm, setIsOpenForm] = React.useState(false)

  return (
    <div>
      <div>
        <h1>Projects</h1>
        <button onClick={() => setIsOpenForm(true)}>New Project</button>
      </div>
      <List />
      {isOpenForm && <ProjectForm onClose={() => setIsOpenForm(false)} />}
    </div>
  )
}

export default ProjectList
