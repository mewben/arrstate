import React from "react"
import cx from "clsx"

import { map, keys, forEach, find } from "@Utils/lodash"
import Th from "./th"
import Td from "./td"

const TableContext = React.createContext()
export const useTable = () => React.useContext(TableContext)

const Table = ({ className, children }) => {
  const [extraTh, setExtraTh] = React.useState({})

  const registerExtraTh = (arr = []) => {
    const extra = extraTh
    forEach(arr, obj => {
      if (!extra[obj._id]) {
        extra[obj._id] = obj
      }
    })
    setExtraTh(extra)
  }

  const renderExtraTh = () => {
    return map(keys(extraTh), k => {
      return (
        <Th key={k} align="right">
          {extraTh[k]?.name}
        </Th>
      )
    })
  }

  const renderExtraTd = (arr = []) => {
    return map(keys(extraTh), k => {
      // find k in arr
      const item = find(arr, { _id: k })
      return (
        <Td key={k} align="right">
          {item?.value}
        </Td>
      )
    })
  }

  return (
    <TableContext.Provider
      value={{ registerExtraTh, renderExtraTh, renderExtraTd }}
    >
      <div
        className={cx(
          "align-middle inline-block min-w-full shadow-sm overflow-hidden sm:rounded-lg",
          className
        )}
      >
        <table className="w-full">{children}</table>
      </div>
    </TableContext.Provider>
  )
}

export default Table
/*
export const Table2 = ({ className, children }) => {
  return (
    <div
      className={cx(
        "align-middle inline-block min-w-full shadow-sm overflow-hidden sm:rounded-lg",
        className
      )}
    >
      <div className="table w-full">{children}</div>
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
*/
