import React from "react"
import PropTypes from "prop-types"
import { AddButton } from "@Components/generic/button"

import { FilePicker } from "@Components/generic/files"

const UploadFileButton = ({
  maxFiles = 1,
  maxSize,
  accept,
  folderId,
  entityType,
  entityId,
  onUpload,
  children,
  buttonText = "Upload",
  autoProceed,
  maxDimensions,
}) => {
  const [isPickerOpen, setIsPickerOpen] = React.useState(false)

  const handleUpload = files => {
    if (onUpload) {
      onUpload(files)
    }
  }

  const openPicker = () => {
    setIsPickerOpen(true)
  }

  const closePicker = () => {
    setIsPickerOpen(false)
  }

  return (
    <>
      {!children ? (
        <AddButton title={buttonText} onClick={openPicker} />
      ) : (
        React.cloneElement(children, { onClick: openPicker })
      )}
      {isPickerOpen && (
        <FilePicker
          onClose={closePicker}
          onUpload={handleUpload}
          createOptions={{
            entityId,
            entityType,
            folderId,
          }}
          maxFiles={maxFiles}
          maxSize={maxSize}
          accept={accept}
          maxDimensions={maxDimensions}
          autoProceed={autoProceed}
        />
      )}
    </>
  )
}

UploadFileButton.propTypes = {
  buttonText: PropTypes.string,
  entityType: PropTypes.string,
  entityId: PropTypes.string,
  folderId: PropTypes.string,
  onUpload: PropTypes.func,
  accept: PropTypes.oneOfType([
    PropTypes.string,
    PropTypes.arrayOf(PropTypes.string),
  ]),
  maxDimensions: PropTypes.array,
  maxFiles: PropTypes.number,
  maxSize: PropTypes.number,
}

export default UploadFileButton
