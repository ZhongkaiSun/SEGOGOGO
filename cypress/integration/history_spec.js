describe('Order History Page Test', function() {
    it('test the title', function () {
        cy.visit('./Front-end/History.html')
        Cypress.on('uncaught:exception', (err, runnable) => {
            // returning false here prevents Cypress from
            // failing the test
            return false
        })
        cy.title().should('eq', 'History')
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
        cy.get('input[type = "Search"]').type("Asian")

        cy.get('button[type = "submit"]')
            .should('have.text', 'Search')

        cy.get('nav')
            .find('a[id="login"]')
            .click()
        cy.url().should('include', 'Login')
        cy.title().should('eq', 'Login')
        cy.go('back')

    })

    it('should visit home page', function () {
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
    })

    it('test alert info', function () {
        cy.get('button[class="btn-close""]')
            .click()

    });


    it('test dropdown list', function () {

    });

    it('test order history table', function () {
        cy.get("tr")
        cy.get('#table1> tbody > tr > td:nth-child(1)').each(($elm, index, $list)=> {
            // text captured from column1
            const t = $elm.text()
            // matching criteria
            if (t.includes('Selenium')) {
                // next sibling captured
                cy.get('#table1 > tbody > tr > td:nth-child(1)')
                    .eq(index).next().then(function (d) {
                    // text of following sibling
                    const r = d.text()
                    //assertion
                    expect(r).to.contains('Commercial')
                })
            }
        })
    })

})
