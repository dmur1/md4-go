# https://oeis.org/A030109 starting from index 16
# 0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15

for i in range(16):
    original_value = format(i, '04b')
    reversed_value_binary = format(int(original_value[::-1], 2), '04b')
    reversed_value_decimal = int(reversed_value_binary, 2)
    print(f"{i:02}: {original_value} | r: {reversed_value_binary} {reversed_value_decimal}")
