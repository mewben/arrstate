export const getURL = ({
  path,
  host = Cypress.env("host"),
  domain = Cypress.env("domain"),
} = {}) => {
  const port = Cypress.env("hostPort")
  return `${domain}.${host}:${port}${path || ""}`
}
