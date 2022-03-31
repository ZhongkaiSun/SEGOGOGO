describe('Account Information Page Test', function(){
    it('test the title', function(){
        cy.visit('./Front-end/AccountInfo.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.title().should('eq', 'AccountInfo')
    })

})
