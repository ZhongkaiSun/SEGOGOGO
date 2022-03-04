describe('My First Test', function(){
    it('Find an element', function(){
        cy.visit('./Front-end/Sign_Up.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.get('form')
        cy.title().should('eq', 'Sign Up')
    })
    it('has a logo', function () {
        cy.visit('./Front-end/Sign_Up.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.get('form').find('img').should('have.attr', 'src').should('include','gator_logo.png')
    })

    it('visit home page', function(){
        cy.visit('./Front-end/Sign_Up.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.get('form')  // get the containing toolbar
            .find('img[src*="gator_logo.png"]')     // *= gives a partial match on src
            .click()
    })
})


// Arrange - setup initial app state
//1. visit a web page
//2. query for an element
// Act - take an action
//3. interact with that element
// Assert - make an assertion
//4. make an assertion about page content