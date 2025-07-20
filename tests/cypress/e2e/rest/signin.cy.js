describe('Login Page', () => {
    context("known user login", () => {
        const user = {
            email: "test@test.com",
            password: "test1243",
        }

        it("signing in registered user", () => {
            cy.request({
                method: "POST", 
                url: "/signin", 
                headers: {
                'Content-Type': 'application/json; charset=utf-8',
                },
                body: user
            })
            .then((response) => {
                expect("token" in response.body).to.eq(true)
                expect("user" in response.body).to.eq(true)

                expect("id" in response.body.user).to.eq(true)
                expect("name" in response.body.user).to.eq(true)
                expect("email" in response.body.user).to.eq(true)
            })
        })
    })


    context("unknown user login", () => {
        const user = {
            email: "unknown@test.com",
            password: "test1243",
        }

        it("signing in unregistered user", () => {
            cy.request({
                method: "POST", 
                url: "/signin", 
                headers: {
                'Content-Type': 'application/json; charset=utf-8',
                },
                body: user,
                failOnStatusCode: false
            })
            .then(response => {
                expect(response.status).to.be.gt(299)  // status returned is 404
            })
        })
    })
})