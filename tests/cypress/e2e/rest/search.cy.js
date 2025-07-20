
describe('Search for give Context word in one of the files', () => {
    context("known user logged in and uploads a file and Searching Context", () => {
        const user = { email: "test@test.com", password: "test1243", }

        it("search context in knowledge base file(s)", () => {
            cy.request({
                method: "POST", 
                url: "/signin", 
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
                body: user
            })
            .then((response) => {
                const token = response.body.token;

                cy.request({
                    method: 'GET',
                    url: '/api/v1/search?q=better',  // search for word better
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                }).then((response) => {
                    expect(response.status).to.eq(200);
                    const expected = {
                        3: "Beautiful is better than ugly.",
                        4: "Explicit is better than implicit.",
                        5: "Simple is better than complex.",
                        6: "Complex is better than complicated.",
                        7: "Flat is better than nested.",
                        8: "Sparse is better than dense.",
                        17: "Now is better than never.",
                        18: "Although never is often better than *right* now."
                    }
                    
                    for (const [key, value] of Object.entries(response.body.matches)) {
                        expect(value).to.eq(expected[key])
                    }
                });
            })
        })
    })
})