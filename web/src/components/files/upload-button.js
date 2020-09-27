import React from "react"
import PropTypes from "prop-types"
import { AddButton } from "@Components/generic/button"

import { FilePicker } from "@Components/generic/files"

const UploadFileButton = ({
  maxFiles,
  maxSize,
  accept,
  folderID,
  entityType,
  entityID,
  onUpload,
  onUploadComplete,
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
          onUploadComplete={onUploadComplete}
          createOptions={{
            entityID,
            entityType,
            folderID,
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
  entityID: PropTypes.string,
  folderID: PropTypes.string,
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
