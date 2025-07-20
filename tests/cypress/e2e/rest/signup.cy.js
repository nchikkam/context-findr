import { faker } from '@faker-js/faker';

describe('Registration Page', () => {
    context("register new user", () => {
        const user = {
            // name: faker.person.fullName(),
            // email: faker.internet.email(),
            name: "Test User",
            email: "test@test.com",
            password: "test1243",
        }

        it("register new user", () => {
            cy.request({
                method: "POST", 
                url: "/register", 
                headers: {
                'Content-Type': 'application/json; charset=utf-8',
                },
                body: user,
                failOnStatusCode: false // throws err on duplication
            })
            .then((response) => {
                // expect("token" in response.body).to.eq(true)
                // in real world scenario, user is created only once.
            })
        })
    })
})