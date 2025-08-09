class TaxCalculator:
    def __init__(self, tax_rate_repository):
        self.tax_rate_repository = tax_rate_repository

    def tax_on(self, amount):
        return amount * self.tax_rate_repository.tax_rate(amount)


class FixedTaxRateRepository:
    def tax_rate(self, amount):
        return 0.15


tax_rate_repository = FixedTaxRateRepository()
my_calculator = TaxCalculator(tax_rate_repository)
print(my_calculator.tax_on(100))  # Outputs: 15.0
