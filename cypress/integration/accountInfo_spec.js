describe('Account Information Page Test', function(){
    it('test the title', function(){
        cy.visit('./Front-end/AccountInfo.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.title().should('eq', 'Account Information')
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
        cy.get('input[type = "Search"]').type("Asian")

        cy.get('button[type = "submit"]')
            .should('have.text','Search')

        cy.get('nav')
            .find('a[id="logout"]')
            .click()
        cy.url().should('include','newHome')
        cy.title().should('eq', 'GatorDash')
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
        cy.visit('./Front-end/AccountInfo.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.get('input[id = "Uname"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'user name')
        cy.get('input[id = "Uname"]').type("root")

        cy.get('input[id = "mobileNumber"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'enter phone number')
        cy.get('input[id = "mobileNumber"]').type("0000000000")

        cy.get('input[id = "email"]')
            .invoke('attr', 'placeholder')
            .should('contain', 'enter email address')
        cy.get('input[id = "email"]').type("root@gmail.com")

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
