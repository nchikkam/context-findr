
describe('upload by users', () => {
    context("known user logged in and uploads a file", () => {
        const user = { email: "test@test.com", password: "test1243", }

        it("verify upload by logged in user works", () => {
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

                cy.fixture('test.txt', 'utf-8').then((fileContent) => {
                    const blob = Cypress.Blob.binaryStringToBlob(fileContent, 'txt');
                    const formData = new FormData();
                    formData.append('file', blob, 'test.txt');
   
                    cy.request({
                        method: 'POST',
                        url: '/api/v1/upload',
                        headers: {
                        'Content-Type': 'multipart/form-data',
                        'Authorization': `Bearer ${token}`
                        },
                        body: formData,
                    }).then((response) => {
                        expect(response.status).to.eq(200);
                    });

                });
            })
        })

        it("verify upload with wrong token doesn't work", () => {
            cy.fixture('test.txt', 'utf-8').then((fileContent) => {
                const blob = Cypress.Blob.binaryStringToBlob(fileContent, 'txt');
                const formData = new FormData();
                formData.append('file', blob, 'test.txt');

                cy.request({
                    method: 'POST',
                    url: '/api/v1/upload',
                    headers: {
                    'Content-Type': 'multipart/form-data',
                    'Authorization': `Bearer wrong token`
                    },
                    body: formData,
                    failOnStatusCode: false
                }).then((response) => {
                    expect(response.status).to.eq(401);
                    expect(response.statusText).to.eq("Unauthorized");
                });
            });
        })
    })
})