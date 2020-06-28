const path = require("path")

module.exports = {
  siteMetadata: {
    title: `App Realty`,
    description: `Realty Management made easy`,
    author: `@mewben`,
  },
  plugins: [
    `gatsby-plugin-react-helmet`,
    `gatsby-transformer-sharp`,
    `gatsby-plugin-sharp`,
    {
      resolve: `gatsby-plugin-manifest`,
      options: {
        name: `gatsby-starter-default`,
        short_name: `starter`,
        start_url: `/`,
        background_color: `#663399`,
        theme_color: `#663399`,
        display: `minimal-ui`,
        icon: `src/images/favicon.png`, // This path is relative to the root of the site.
      },
    },
    // this (optional) plugin enables Progressive Web App + Offline functionality
    // To learn more, visit: https://gatsby.dev/offline
    // `gatsby-plugin-offline`,
    {
      resolve: `gatsby-plugin-create-client-paths`,
      options: {
        prefixes: [`/projects/*`, `/lots/*`],
      },
    },
    {
      resolve: "gatsby-plugin-root-import",
      options: {
        "@Pages": path.join(__dirname, "src/pages"),
        "@Wrappers": path.join(__dirname, "src/wrappers"),
        "@Providers": path.join(__dirname, "src/providers"),
        "@Components": path.join(__dirname, "src/components"),
        "@Screens": path.join(__dirname, "src/screens"),
        "@Utils": path.join(__dirname, "src/utils"),
        "@Enums": path.join(__dirname, "src/enums"),
        "@Store": path.join(__dirname, "src/store"),
        "@Hooks": path.join(__dirname, "src/hooks"),
      },
    },
    `gatsby-plugin-material-ui`,
    `gatsby-plugin-postcss`,
  ],
}
