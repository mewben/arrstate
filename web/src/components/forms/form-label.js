import React from "react"

const FormLabel = ({ label }) => {
  if (!label) {
    return null
  }
  return <div>{label}</div>
}

export default FormLabel
