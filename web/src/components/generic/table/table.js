import React from "react"

const Table = ({ children }) => {
  return (
    <div className="align-middle inline-block min-w-full shadow-sm overflow-hidden sm:rounded-lg">
      <table className="min-w-full">{children}</table>
    </div>
  )
}

export default Table
