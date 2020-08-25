import React from "react"
import Typography from "@material-ui/core/Typography"
import Skeleton from "@material-ui/lab/Skeleton"

const Loading = ({ variant = "default", typo }) => {
  if (variant === "default") {
    return (
      <div className="flex w-full h-full items-center justify-center">
        <svg
          className="h-8 w-8 text-cool-gray-400"
          viewBox="0 0 80 80"
          fill="currentColor"
        >
          <path
            d="M40,72C22.4,72,8,57.6,8,40C8,22.4,
		22.4,8,40,8c17.6,0,32,14.4,32,32c0,1.1-0.9,2-2,2
		s-2-0.9-2-2c0-15.4-12.6-28-28-28S12,24.6,12,40s12.6,
		28,28,28c1.1,0,2,0.9,2,2S41.1,72,40,72z"
          >
            <animateTransform
              attributeType="xml"
              attributeName="transform"
              type="rotate"
              from="0 40 40"
              to="360 40 40"
              dur="0.6s"
              repeatCount="indefinite"
            />
          </path>
        </svg>
      </div>
    )
  }

  if (variant === "skeleton") {
    return (
      <Typography variant={typo}>
        <Skeleton />
      </Typography>
    )
  }
}

export default Loading
