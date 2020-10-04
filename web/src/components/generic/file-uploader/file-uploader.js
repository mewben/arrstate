import React from "react"
import PropTypes from "prop-types"
import Uppy from "@uppy/core"
import Tus from "@uppy/tus"
import { DashboardModal } from "@uppy/react"
import { useTranslation } from "react-i18next"

import { Button } from "@Components/generic"

const FileUploader = ({
  openByClickOn,
  onSuccess,
  maxNumberOfFiles = 1,
  maxFileSize = 20 * 1000000, // 20MB
  allowedFileTypes = ["image/*"], // set this to null to accept everything
}) => {
  const [open, setOpen] = React.useState(false)
  const { t } = useTranslation()
  const uppy = React.useMemo(() => {
    // Do all the configuration here
    return new Uppy({
      restrictions: { maxNumberOfFiles, maxFileSize, allowedFileTypes },
    }).use(Tus, { endpoint: "/files/" })
  }, [])

  React.useEffect(() => {
    return () => uppy.close()
  }, [])

  uppy.on("complete", result => [console.log("complete result", result)])

  const onOpen = () => {
    setOpen(true)
  }

  const onClose = () => {
    setOpen(false)
  }

  const button = openByClickOn ? (
    React.cloneElement(openByClickOn, { onClick: onOpen })
  ) : (
    <Button onClick={onOpen}>{t("upload")}</Button>
  )

  return (
    <>
      {button}
      <DashboardModal
        uppy={uppy}
        open={open}
        onRequestClose={onClose}
        animateOpenClose={false}
        proudlyDisplayPoweredByUppy={false}
        metaFields={[
          { id: "name", name: "Name", placeholder: "FileName" },
          {
            id: "caption",
            name: "Caption",
            placeholder: "Describe what the file is about",
          },
        ]}
        showProgressDetails
      />
    </>
  )
}

FileUploader.propTypes = {
  onSuccess: PropTypes.func.isRequired,
}

export default FileUploader
