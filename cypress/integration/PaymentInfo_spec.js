describe('Payment Information Page Test', function(){
    it('test the title', function(){
        cy.visit('./Front-end/PaymentInfo.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.title().should('eq', 'Payment Information')
    })

    it('should have a navigation bar', function () {
        cy.get('nav')
            .find('img')
            .should('have.attr', 'src')
            .should('include','gator_logo.png','Engineering%20Gator.png')

        cy.get('nav')  // get the containing toolbar
            .find('img[src*="gator_logo.png"]')     // *= gives a partial match on src
            .click()
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.url().should('include','newHome')
        cy.title().should('eq', 'GatorDash')
        cy.go('back')

        cy.get('nav')
            .find('a[id="home"]')
            .click()
        cy.url().should('include','newHome')
        cy.title().should('eq', 'GatorDash')
        cy.go('back')

        cy.get('nav')
            .find('a[id="History"]')
            .click()
        cy.url().should('include','History')
        cy.go('back')

        cy.get('input[type = "Search"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'Search')
        cy.get('input[type = "Search"]').type("KFC")

        cy.get('a[id = "search"]')
            .should('have.text','Search')
            .click()
        cy.url().should('include','Restaurant')

        cy.go('back')



    })

    it('should visit home page', function(){
        cy.get('nav')  // get the containing toolbar
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
        cy.visit('./Front-end/PaymentInfo.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.get('input[id = "cardHolder"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'card holder')
        cy.get('input[id = "cardHolder"]').type("Hongru")

        cy.get('input[id = "cardNumber"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'enter card number')
        cy.get('input[id = "cardNumber"]').type("0000000000000000")

        cy.get('input[id = "expdate"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'MM/YY')
        cy.get('input[id = "expdate"]').type("09/24")

        cy.get('input[id = "securityCode"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'XXX')
        cy.get('input[id = "securityCode"]').type("123")

        cy.get('input[id = "address1"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'enter address line 1')
        cy.get('input[id = "address1"]').type("2901 SW 13th St")

        cy.get('input[id = "address2"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'enter address line 2')
        cy.get('input[id = "address2"]').type("APT 250")

        cy.get('input[id = "city"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'City')
        cy.get('input[id = "city"]').type("Gainesville")

        cy.get('input[id = "state"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'State')
        cy.get('input[id = "state"]').type("Florida")

        cy.get('input[id = "zipcode"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'ZIP code')
        cy.get('input[id = "zipcode"]').type("32608")
    })

    it('test submit', function () {
        //cy.get('button[id="save"]').click()
        //cy.go('back')

        cy.get('button[id="backhome"]')
            .click()
        cy.url().should('include','newHome')
        cy.title().should('eq', 'GatorDash')
        cy.go('back')
    });

})
