go_binary(
    name = 'main',
    srcs = [
        'main.go',
    ],
    deps = [
        ':generated_asset',
    ],
)

genrule(
    name = 'generated_asset',
    out = 'generated_asset',
    cmd = 'echo \'hello generated asset world\' > $OUT',
)
