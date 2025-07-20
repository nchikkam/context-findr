describe('The Home Page', () => {
    context("GET /", () => {
        it("visit home page", () => {
            cy.request("GET", "/").then((response) => {
                expect(response.status).to.eq(200)
                expect(response.body.message).equal("context-finder home page")
            })
        })
    })
})