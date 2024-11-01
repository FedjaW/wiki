# Filter Command

Users like the filter command, but think it could be better. They want access to more than two filter options and to customize the words that get filtered. It's time to upgrade the filter command feature.

Eine super variante einen filter command zu desginen.

```py
def get_filter_cmd(filters):
    def filter_cmd(content, options, word_pairs):
        if len(options) == 0:
            raise Exception("missing options")
            
        for option in options:
            if option in filters:
                filter_func = filters[option]
                content = filter_func(content, word_pairs)
            else:
                raise Exception("invalid option")
        return content

    return filter_cmd
    
def replace_words(content, word_pairs):
    for pair in word_pairs:
        content = content.replace(pair[0], pair[1])
    return content


def remove_words(content, word_pairs):
    for pair in word_pairs:
        content = content.replace(pair[0], "")
    return content


def capitalize_sentences(content, word_pairs):
    return ". ".join(map(str.capitalize, content.split(". ")))


def uppercase_words(content, word_pairs):
    for pair in word_pairs:
        content = content.replace(pair[0], pair[0].upper())
    return content


filters = {
    "--replace": replace_words,
    "--remove": remove_words,
    "--capitalize": capitalize_sentences,
    "--uppercase": uppercase_words,
}
```