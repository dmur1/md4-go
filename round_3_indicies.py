for i in range(16):
    original_value = format(i, '04b')
    reversed_value_binary = format(int(original_value[::-1], 2), '04b')
    reversed_value_decimal = int(reversed_value_binary, 2)
    print(f"{i:02}: {original_value} | r: {reversed_value_binary} {reversed_value_decimal}")
