import React from "react"
import { Loading, Error, Empty } from "@Components/generic"

export const ListProvider = ({
  getMethod,
  getMethodParams,
  emptyRenderer,
  contentRenderer,
}) => {
  const { status, data, error, isFetching } = getMethod(getMethodParams)

  const renderEmpty = () => {
    return emptyRenderer ? emptyRenderer() : <Empty />
  }

  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : !data?.total ? (
    renderEmpty()
  ) : (
    contentRenderer(data)
  )
}

export const SingleProvider = ({
  getMethod,
  id,
  emptyRenderer,
  contentRenderer,
}) => {
  const { status, data, error, isFetching } = getMethod(id)

  const renderEmpty = () => {
    return emptyRenderer ? emptyRenderer() : <Empty />
  }

  return status === "loading" || isFetching ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : !data ? (
    renderEmpty()
  ) : (
    contentRenderer(data)
  )
}
