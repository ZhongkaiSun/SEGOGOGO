describe('Restaurant Page Test', function() {
    it('test the title', function () {
        cy.visit('./Front-end/Restaurant.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.title().should('eq', 'Restaurant')
    })

    it('test the navigation bar', function () {
        cy.get('nav')
            .find('img')
            .should('have.attr', 'src')

        cy.get('nav')  // get the containing toolbar
            .find('img[src*="gator_logo.png"]')     // *= gives a partial match on src
            .click()
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.url().should('include', 'newHome')
        cy.title().should('eq', 'GatorDash')
        cy.go('back')

        cy.get('nav')
            .find('a[id="home"]')
            .click()
        cy.url().should('include', 'newHome')
        cy.title().should('eq', 'GatorDash')
        cy.go('back')

        cy.get('nav')
            .find('a[id="History"]')
            .click()
        cy.url().should('include', 'History')
        cy.go('back')

        cy.get('input[type = "Search"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'Search')

        cy.get('input[type = "Search"]').type("Popeyes")

        cy.get('a[id = "search"]')
            .should('have.text','Search')
            .click()
        cy.url().should('include','Restaurant')
    })
    it('test the menu', function () {
        cy.get('#card0')
            .contains("$ 4.79")

        cy.get('#card0')
            .contains("8PC Nuggets A La Carte")

        cy.get('#card0')
            .contains("Calories = 400")
    });


})
