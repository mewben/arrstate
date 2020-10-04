import React, { useMemo } from "react"
import PropTypes from "prop-types"
import { DashboardModal } from "@uppy/react"

import { fileUploadClient } from "./utils/client"
import { onUpload as handleUpload } from "./utils/on-upload"

const FilePicker = ({
  onClose,
  onUpload,
  onUploadComplete,
  createOptions,
  maxFiles,
  maxSize,
  accept,
  maxDimensions,
  autoProceed,
  createFile,
}) => {
  const pickerInstance = useMemo(() => {
    const client = fileUploadClient({
      maxFiles,
      maxSize,
      accept,
      autoProceed,
    })

    client.on("upload-success", (file, response) => {
      handleUpload(
        [
          {
            ...file,
            ...response,
          },
        ],
        {
          ...createOptions,
          onUpload,
          createFile,
        }
      )
    })

    client.on("complete", result => {
      if (onUploadComplete) {
        onUploadComplete(result)
      }
    })

    return client
  }, [])

  const getNote = () => {
    if (maxDimensions) {
      return `RECOMMENDED SIZE IS ${maxDimensions[0]}px by ${maxDimensions[1]}px`
    }
  }

  return (
    <DashboardModal
      open
      onRequestClose={onClose}
      uppy={pickerInstance}
      plugins={["URL"]}
      proudlyDisplayPoweredByUppy={false}
      note={getNote()}
      metaFields={[{ id: "name", name: "Name", placeholder: "File name" }]}
    />
  )
}

FilePicker.propTypes = {
  onClose: PropTypes.func.isRequired,
  onUpload: PropTypes.func.isRequired,
  createOptions: PropTypes.object,
  maxFiles: PropTypes.number,
  maxSize: PropTypes.number,
  accept: PropTypes.string,
  maxDimensions: PropTypes.arrayOf(PropTypes.number),
}

export default FilePicker
