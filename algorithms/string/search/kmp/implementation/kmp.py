# Сложность составления таблицы префиксов - O(m)
def create_prefix_table(pattern):
    prefix_table = [0] * len(pattern)
    j = 0
    i = 1

    while i < len(pattern):
        if pattern[j] == pattern[i]:
            prefix_table[i] = j + 1
            i += 1
            j += 1
        else:
            if j == 0:
                prefix_table[i] = 0
                i += 1
            else:
                j = prefix_table[j - 1]

    return prefix_table


# Сложность KMP-алгоритма - O(n+m)
def KMP(pattern, text):
    prefix_table = create_prefix_table(pattern)
    print(prefix_table)

    pattern_length = len(pattern)
    text_length = len(text)

    i = 0
    j = 0
    while i < text_length:
        if text[i] == pattern[j]:
            i += 1
            j += 1
            if j == pattern_length:
                print("pattern is found")
                break
        else:
            if j > 0:
                j = prefix_table[j - 1]
            else:
                i += 1

    if i == text_length:
        print("pattern is not found")


pattern = "ababac"
text = "ababaqde ababacde"

KMP(pattern, text)
