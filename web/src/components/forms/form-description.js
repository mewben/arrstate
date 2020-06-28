import React from "react"

const FormDescription = ({ description }) => {
  if (!description) {
    return null
  }

  return <p className="mt-2 text-sm text-gray-500">{description}</p>
}

export default FormDescription
