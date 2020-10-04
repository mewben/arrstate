import { isEmpty } from "@Utils/lodash"

const filesEndpoint = process.env.GATSBY_S3_FILES_ENDPOINT

export const uploadURL = file => {
  if (isEmpty(file)) {
    return ""
  }

  return `${filesEndpoint}${file.url}`
}
