import { visit, getURL } from "../../utils"

context("Routes", () => {
  describe("If not logged in", () => {
    it("should redirect to /login", () => {
      cy.visit(getURL({ path: "/" }))
      cy.location("pathname").should("equal", "/login")
    })
  })
})
