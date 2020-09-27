import React from "react"
import { Controller } from "react-hook-form"
import { useTranslation } from "react-i18next"
import DeleteIcon from "@material-ui/icons/Delete"
import PublishIcon from "@material-ui/icons/Publish"
import WarningIcon from "@material-ui/icons/Warning"
import cx from "clsx"

import { Button, ConfirmButton } from "@Components/generic/button"
import { UploadFileButton } from "@Components/files"
import { pick, isEmpty } from "@Utils/lodash"
import { uploadURL } from "@Utils"
import { FILE_TYPES } from "@Enums"

const attachmentFields = ["_id", "title", "ext", "mimeType", "size", "url"]

const AttachmentField = ({ name, ...props }) => {
  return (
    <Controller
      name={name}
      render={({ onChange, value }) => {
        return <Component onChange={onChange} value={value} {...props} />
      }}
    />
  )
}

const Component = ({
  value,
  onChange,
  label,
  entityType,
  entityID,
  isAttachment,
  className,
  isOutlined = true,
  size,
  isFullWidth,
}) => {
  const { t } = useTranslation()

  const hasValue = !isEmpty(value)
  const imageURL = React.useMemo(() => {
    return uploadURL(value)
  }, [value])

  const style = hasValue
    ? {
        backgroundImage: `url("${imageURL}")`,
      }
    : undefined

  const handleChange = data => {
    console.log("handleChange", data)
    if (isEmpty(data)) {
      onChange(null)
    } else {
      onChange(isAttachment ? pick(data, attachmentFields) : data)
    }
  }

  let content
  if (!hasValue) {
    content = (
      <UploadFileButton
        entityType={entityType}
        entityID={entityID}
        onUpload={handleChange}
        accept={FILE_TYPES.IMAGE}
        maxFiles={1}
      >
        <div className="image-badge-trigger">
          <div className="image-badge-text">{label || t("image")}</div>
          <div className="image-badge-icon">
            <PublishIcon />
          </div>
        </div>
      </UploadFileButton>
    )
  } else {
    content = (
      <div className="image-badge-trigger">
        <ConfirmButton
          button={
            <Button circle>
              <DeleteIcon />
            </Button>
          }
          confirmButton={
            <Button color="red" circle>
              <WarningIcon />
            </Button>
          }
          onSubmit={handleChange.bind(null, {})}
        />
      </div>
    )
  }
  return (
    <div
      className={cx("image-badge", className, size, {
        outlined: isOutlined,
        fullwidth: isFullWidth,
        "has-value": hasValue,
      })}
    >
      <div className="image-badge-content" style={style}>
        <div className="image-badge-trigger">{content}</div>
      </div>
    </div>
  )
}

export default AttachmentField
