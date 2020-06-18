export const visit = url => {
  cy.visit(url)

  cy.get(".loader").should("not.exist")
}
