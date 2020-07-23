import { reduce, get } from "@Utils/lodash"

// group blocks type adjacent
export const groupBlocks = blocks => {
  return reduce(
    blocks,
    (prev, curr) => {
      if (
        prev.length &&
        curr.type === get(prev, [prev.length - 1, 0, "type"])
      ) {
        prev[prev.length - 1].push(curr)
      } else {
        prev.push([curr])
      }
      return prev
    },
    []
  )
}
