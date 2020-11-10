import React from "react"
import defaultTheme from "tailwindcss/defaultTheme"
import {
  Chart,
  antvLight,
  registerTheme,
  createThemeByStylesheet,
} from "bizcharts"

const ChartWrapper = ({ children, ...props }) => {
  antvLight.fontFamily = ["Inter var", ...defaultTheme.fontFamily.sans].join(
    ","
  )

  registerTheme("default", createThemeByStylesheet(antvLight))

  return (
    <Chart {...props} theme="default">
      {children}
    </Chart>
  )
}

export default ChartWrapper
