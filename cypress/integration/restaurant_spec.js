describe('Restaurant Page Test', function(){
    it('test the title', function(){
        cy.visit('./Front-end/Restaurant.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })

    })

})
