describe('Restaurant-Cuisine Page Test', function () {
    it('visit the page', function () {
        cy.visit('./Front-end/newHome.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.title().should('eq', 'GatorDash')

        cy.get('#american')  // get the containing toolbar
            .find('img[src*="american.png"]')     // *= gives a partial match on src
            .click()
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
        cy.get('#american')  // get the containing toolbar
            .find('img[src*="american.png"]')     // *= gives a partial match on src
            .click()

        cy.get('nav')
            .find('a[id="home"]')
            .click()
        cy.url().should('include', 'newHome')
        cy.title().should('eq', 'GatorDash')
        cy.get('#american')  // get the containing toolbar
            .find('img[src*="american.png"]')     // *= gives a partial match on src
            .click()

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
            .should('have.text', 'Search')
            .click()
        cy.url().should('include', 'Restaurant')
        cy.visit('./Front-end/newHome.html')
        cy.get('#asian')  // get the containing toolbar
            .find('img[src*="Asian.png"]')     // *= gives a partial match on src
            .click()
    })


    it('test the restaurant list', function () {
        cy.get('#card0')
            .contains("TeaBestTea")

        cy.get('#card0')
            .contains("Rating = 4.2/5.0")

        cy.get('#card0')
            .contains("Delivery Fee = $3.99")

        cy.get('#card1')
            .contains("Yummy")

        cy.get('#card1')
            .contains("Rating = 4.2/5.0")

        cy.get('#card1')
            .contains("Delivery Fee = $4.99")

        cy.get('#card0')
            .find('a[id="ordernow0"]')
            .click()

        cy.url().should('include', 'Restaurant')
        cy.get('#restaurantName')
            .contains("TeaBestTea")
    });


})