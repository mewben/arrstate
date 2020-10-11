import React from "react"
import filesize from "filesize"
import { useMutation, queryCache } from "react-query"
import DeleteIcon from "@material-ui/icons/Delete"
import WarningIcon from "@material-ui/icons/Warning"

import { Td } from "@Components/generic"
import { uploadURL, requestApi } from "@Utils"
import { Button, ConfirmButton } from "@Components/generic/button"

const ListItem = ({ item }) => {
  const [remove] = useMutation(
    () => {
      return requestApi(`/api/files/${item._id}`, "DELETE")
    },
    {
      onSuccess: () =>
        queryCache.invalidateQueries([
          "files",
          { entityType: item.entityType, entityID: item.entityID },
        ]),
    }
  )

  const onDelete = async () => {
    await remove()
  }

  return (
    <tr>
      <Td wrap>
        <a
          className="font-medium text-cool-gray-800 hover:text-blue-500"
          href={uploadURL(item)}
          download
        >
          {item.title}
        </a>
      </Td>
      <Td>{item.ext}</Td>
      <Td align="right">{filesize(item.size)}</Td>
      <Td align="right">
        <ConfirmButton
          button={
            <Button variant="text" circle>
              <DeleteIcon />
            </Button>
          }
          confirmButton={
            <Button color="red" circle>
              <WarningIcon />
            </Button>
          }
          onSubmit={onDelete}
        />
      </Td>
    </tr>
  )
}

export default ListItem
