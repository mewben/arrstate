import React from "react"
import cx from "clsx"

import { initialsName, uploadURL } from "@Utils"

const PersonBadge = ({ person, size = "sm", canEdit }) => {
  const [hovered, setHovered] = React.useState(false)
  const initials = initialsName(person?.name)
  const imageURL = React.useMemo(() => {
    return uploadURL(person.avatar)
  }, [person])

  const cl = cx(
    "flex items-center justify-center bg-cool-gray-200 text-cool-gray-500 uppercase",
    size === "sm" && "w-6 h-6 rounded-sm text-xs",
    size === "2xl" && "w-40 h-40 rounded-lg border-4 border-white text-5xl"
  )

  let content = <span>{initials}</span>
  if (imageURL) {
    content = <img src={imageURL} alt={person._id} />
  }

  if (canEdit) {
    const onMouseOver = () => {
      setHovered(!hovered)
    }

    content = (
      <button
        type="button"
        className="font-medium hover:text-indigo-500 focus:outline-none focus:underline transition duration-150 ease-in-out"
        onMouseOver={onMouseOver}
      >
        {hovered ? "upload" : content}
      </button>
    )
  }

  return <div className={cl}>{content}</div>
}

export default PersonBadge
