describe('Login Page Test', function(){
    it('test the title', function(){
        cy.visit('./Front-end/Login.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.get('form')
        cy.title().should('eq', 'Login')
    })
    it('has a logo', function () {
        cy.get('form').find('img').should('have.attr', 'src').should('include','gator_logo.png')
    })

    it('visit home page', function(){
        cy.get('form')  // get the containing toolbar
            .find('img[src*="gator_logo.png"]')     // *= gives a partial match on src
            .click()
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.url().should('include','newHome')
        cy.title().should('eq', 'GatorDash')
    })
    it('test the input', function() {
        cy.visit('./Front-end/Login.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.get('input[name = "Uname"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'Username')
        cy.get('input[name = "Uname"]').type("root")
        cy.get('input[name = "Pass"]').invoke('attr', 'placeholder')
            .should('contain', 'Password')
        cy.get('input[name = "Pass"]').type("root@gmail.com")
    })
    it('test the checkbox', function(){
        cy.get('#check').check("Remember me")
    })
    it('check the Login button', function(){
        cy.get('#log').click()
        cy.on('window:alert',(txt)=> {
            //Mocha assertions
            expect(txt).to.contains('Successfully');
        })
    })
    it('check the Sign up button', function(){
        cy.get('#sign-up').click()
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.url().should('include','Sign_Up')
        cy.title().should('eq', 'Sign Up')
        cy.go('back')
    })
})
