import React from "react"
import cx from "clsx"

const Table = ({ className, children }) => {
  return (
    <div
      className={cx(
        "align-middle inline-block min-w-full shadow-sm overflow-hidden sm:rounded-lg",
        className
      )}
    >
      <table className="w-full">{children}</table>
    </div>
  )
}

export default Table

export const Table2 = ({ className, children }) => {
  return (
    <div
      className={cx(
        "align-middle inline-block min-w-full shadow-sm overflow-hidden sm:rounded-lg",
        className
      )}
    >
      <div className="table w-full">{children}</div>
      {/* <table className="min-w-full table-fixed">{children}</table> */}
    </div>
  )
}

export const THead2 = ({ className, children }) => {
  return <div className={cx("table-header-group", className)}>{children}</div>
}

export const TBody2 = ({ className, children }) => {
  return <div className={cx("table-row-group", className)}>{children}</div>
}

export const TFoot2 = ({ className, children }) => {
  return <div className={cx("table-footer-group", className)}>{children}</div>
}

export const TRow2 = ({ className, children }) => {
  return <div className={cx("table-row", className)}>{children}</div>
}
