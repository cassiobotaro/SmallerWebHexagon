def tax_on(tax_rate):
    def calculate_tax(amount):
        return amount * tax_rate(amount)

    return calculate_tax


def fixed_tax_rate(amount):
    return 0.15


tax_on(fixed_tax_rate)(100)  # Outputs: 15.0
