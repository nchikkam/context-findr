
describe('Upload by a specific user', () => {
    context("known user logged in and uploads a file", () => {
        const user = { email: "test@test.com", password: "test1243", }

        it("Verify upload by logged in user works", () => {
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
                    url: '/api/v1/uploads',
                    headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                    'Authorization': `Bearer ${token}`
                    }
                }).then((response) => {
                    expect(response.status).to.eq(200);
                    for(let uploaded_file of response.body.files){
                        expect(uploaded_file.email).to.eq(user.email);    
                    }
                });
            })
        })
    })
})