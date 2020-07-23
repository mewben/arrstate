import dayjs from "dayjs"

export const formatDate = (d, format) => {
  return dayjs(d).format(format)
}
