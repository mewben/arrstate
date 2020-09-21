import Uppy from "@uppy/core"
import Tus from "@uppy/tus"
// import AwsS3Multipart from "@uppy/aws-s3-multipart"
// import GoogleDrive from "@uppy/google-drive"
// import Dropbox from "@uppy/dropbox"
// import Facebook from "@uppy/facebook"
// import Instagram from "@uppy/instagram"
// import OneDrive from "@uppy/onedrive"
// import URL from "@uppy/url"
// import Webcam from "@uppy/webcam"
// import ScreenCapture from "@uppy/screen-capture"
import { FILE_UPLOAD_SETTINGS } from "@Enums/entity/file"
// import { SETTING_PUBLIC } from "@Enums"

// const companionUrl = SETTING_PUBLIC.FILE_UPLOADS.ENDPOINT

export const fileUploadClient = ({
  autoProceed,
  maxFiles = FILE_UPLOAD_SETTINGS.MAX_FILES,
  maxSize = FILE_UPLOAD_SETTINGS.MAX_SIZE,
  accept,
} = {}) =>
  new Uppy({
    autoProceed,
    allowMultipleUploads: maxFiles > 1,
    restrictions: {
      maxNumberOfFiles: maxFiles,
      maxFileSize: maxSize,
      allowedFileTypes: accept,
    },
    locale: {
      strings: {
        youCanOnlyUploadX: {
          0: "You can only upload %{smart_count} file",
          1: "You can only upload %{smart_count} files",
        },
        youHaveToAtLeastSelectX: {
          0: "You have to select at least %{smart_count} file",
          1: "You have to select at least %{smart_count} files",
        },
        exceedsSize2: "This file exceeds maximum allowed size of %{size}",
        youCanOnlyUploadFileTypes: "You can only upload: %{types}",
        companionError: "Connection with Companion failed",
      },
    },
  }).use(Tus, {
    endpoint: "/files/",
  })
// .use(AwsS3Multipart, {
//   companionUrl,
// })
// .use(GoogleDrive, {
//   id: "GoogleDrive",
//   companionUrl,
// })
// .use(Dropbox, {
//   id: "Dropbox",
//   companionUrl,
// })
// .use(Facebook, {
//   id: "Facebook",
//   companionUrl,
// })
// .use(Instagram, {
//   id: "Instagram",
//   companionUrl,
// })
// .use(OneDrive, {
//   id: "OneDrive",
//   companionUrl,
// })
// .use(URL, {
//   id: "URL",
//   companionUrl,
// })
// .use(Webcam, {
//   id: "Webcam",
//   companionUrl,
// })
// .use(ScreenCapture, {
//   id: "ScreenCapture",
//   companionUrl,
// })
