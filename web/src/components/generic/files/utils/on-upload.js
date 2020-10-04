import { requestApi } from "@Utils"

const filesEndpoint = process.env.GATSBY_S3_UPLOAD_URL_BASE

export const onUpload = async (files = [], { onUpload, ...createOptions }) => {
  for (const file of files) {
    const res = await requestApi("/api/files", "POST", {
      data: {
        title: file.name,
        ext: file.extension,
        mimeType: file.type,
        size: file.size,
        url: file.uploadURL.replace(filesEndpoint, ""),
        ...createOptions,
      },
    })
    onUpload(res.data)
  }
}

const file = {
  source: "react:DashboardModal",
  id:
    "uppy-screen/shot/2020/09/14/at/3/22/07/pm/png-10-10-1d-1d-10-10-1e-1e-10-1e-image/png-48332-1600068133411",
  name: "Screen Shot 2020-09-14 at 3.22.07 PM.png",
  extension: "png",
  meta: {
    relativePath: null,
    name: "Screen Shot 2020-09-14 at 3.22.07 PM.png",
    type: "image/png",
  },
  type: "image/png",
  data: {},
  progress: {
    percentage: 0,
    bytesUploaded: 0,
    bytesTotal: 48332,
    uploadComplete: false,
    uploadStarted: null,
  },
  size: 48332,
  isRemote: false,
  remote: "",
  preview:
    "blob:http://jose.arrstate.lh:8000/c7e16a6c-6ddf-4222-acc2-8dd903488315",
  uploadURL: "http://localhost:8080/files/58557b9f314079fbcc3baf8d4814d0b1",
}
