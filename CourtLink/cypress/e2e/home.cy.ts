describe("Home Page Tests", () => {
    beforeEach(() => {
      cy.visit("/");
    });
  
    it("should display the navbar", () => {
      cy.get("app-navbar").should("be.visible");
    });
  
    it("should display the sports selection", () => {
      cy.get("app-sports").should("be.visible");
    });
  
    it("should navigate to the courts page when a sport is selected", () => {
      cy.get("app-sports div").first().click(); 
      cy.url().should("include", "/courts");
    });
  });
  