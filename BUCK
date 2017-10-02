go_binary(
    name = 'main',
    srcs = [
        'main.go',
    ],
    # The issue: deps inside go_binary are always assumed
    # to be a library.  we cannot include other rules that 
    # synthesize data files or whatnot
    deps = [
        ':generated_asset',
    ],
)

# a totally artificial example of a build rule with side affects
# that we might want to include from a golang build target.
genrule(
    name = 'generated_asset',
    out = 'generated_asset',
    cmd = 'echo \'hello generated asset world\' > $OUT',
)
