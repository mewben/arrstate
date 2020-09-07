import React from "react"

import { fullName } from "@Utils"
import PersonBadge from "./person-badge"

// Person full name with badge
const Person = ({ person, hideBadge, hideName }) => {
  return (
    <div className="flex items-center space-x-2">
      {!hideBadge && <PersonBadge person={person} />}
      {!hideName && <span>{fullName(person?.name)}</span>}
    </div>
  )
}

export default Person
