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
    it('test nav bar', function () {
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
        cy.go('back')
    });
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
    it('test Cuisine category', function () {
        cy.get('#american')  // get the containing toolbar
            .find('img[src*="american.png"]')     // *= gives a partial match on src
            .click()
        cy.get('#category').contains("American")
        cy.go('back')

        cy.get('#coffeeandtea')  // get the containing toolbar
            .find('img[src*="Coffee.png"]')     // *= gives a partial match on src
            .click()
        cy.get('#category').contains("Coffee")
        cy.go('back')

    });

    it('test cart', function () {
        cy.get('#cd-cart-trigger').click()
        cy.get('#cd-cart')
            .find('a[id = "checkout"]')
            .click()
        cy.on('window:alert', (txt) => {
            //Mocha assertions
            expect(txt).to.contains('Please add something into your cart! Or login first.');
        })
        cy.visit('./Front-end/newHome.html')
    });

    it('test favorite restaurants', function () {
        cy.get('#fav1')
            .contains('Popeyes')
        cy.get('#fav1')
            .contains('Order Now')


    });
})
