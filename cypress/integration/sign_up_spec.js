describe('Sign Up Page Test', function(){
    it('test the title', function(){
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
        cy.get('form').find('img').should('have.attr', 'src').should('include','gator_logo.png')
    })

    it('visit home page', () => {
        cy.get('form')  // get the containing toolbar
            .find('img[src*="gator_logo.png"]')     // *= gives a partial match on src
            .click()
        cy.url().should('include','newHome')
        cy.title().should('eq', 'GatorDash')
    })
    it('test the input', function(){
        cy.visit('./Front-end/Sign_Up.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.get('input[name = "Uname"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'User Name')
        cy.get('input[name = "Uname"]').type("root")
        cy.get('input[name = "Email"]').invoke('attr', 'placeholder')
            .should('contain', 'E-mail Address')
        cy.get('input[name = "Email"]').type("root@gmail.com")

        cy.get('#CPass').invoke('attr', 'placeholder')
            .should('contain', 'Enter Your Password')
        cy.get('#CPass').type("1")
        cy.get('#RPass') .invoke('attr', 'placeholder')
            .should('contain', 'Re-enter Your Password')
        cy.get('#RPass').type("2")
        cy.get('#notice').contains("Passwords do not match")

        cy.get('#CPass').clear().type("123456")
        cy.get('#RPass').clear().type("123456")
        cy.get('#notice').contains("Passwords match")
    })
    it('test the checkbox', function(){
        cy.get('#check').check("I agree to all the terms of service.")
    })
    it('check the Sign up button', function(){
        cy.get('#sign-up').click()
        cy.on('window:alert',(txt)=> {
            //Mocha assertions
            expect(txt).to.contains('The user already exists');
        })
    })
})


// Arrange - setup initial app state
//1. visit a web page
//2. query for an element
// Act - take an action
//3. interact with that element
// Assert - make an assertion
//4. make an assertion about page content