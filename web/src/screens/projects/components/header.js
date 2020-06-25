import React from "react"

import ProjectForm from "./form"

const Header = ({ project }) => {
  const [isOpenForm, setIsOpenForm] = React.useState(false)

  return (
    <div>
      <h1>{project.name}</h1>
      <button onClick={() => setIsOpenForm(true)}>Edit Project</button>
      {isOpenForm && (
        <ProjectForm onClose={() => setIsOpenForm(false)} model={project} />
      )}
    </div>
  )
}

export default Header
