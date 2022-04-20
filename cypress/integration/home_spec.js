describe('Home Page Test', function () {
    it('test the title', function () {
        cy.visit('./Front-end/newHome.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.get('form')
        cy.title().should('eq', 'GatorDash')
    })

    it('test the login and logout', function () {
        cy.get('nav')
            .find('a[id="Login"]')
            .click()
        cy.get('input[name = "Uname"]').type("root")
        cy.get('input[name = "Pass"]').type("123456")
        cy.get('#log')
            .click()

        cy.get('nav')
            .find('a[id="Logout"]')
            .click()
    });

})
