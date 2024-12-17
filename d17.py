def main():
    want = [2,4,1,5,7,5,0,3,4,1,1,6,5,5,3,0]

    start = 8 ** (len(want) - 1)
    print(start)

    n = start
    inc = 1
    wantPos = 0

    while True:
        n += inc
        out = run(n, want)
        if len(out) > 8:
            print(n, out)


def run(start, want):
    out = []
    idx = 0
    a = start

    while a != 0:
        # print(f"loop: {a=}")
        b = (a % 8) ^ 5
        c = a // (2**b)
        a = a // 8
        b = b ^ c  ^ 6

        out.append(b % 8)
        if out[-1] != want[idx]:
            return out

        idx += 1


    return out

    # print(f"after, {a=} {b=} {c=}")





if __name__ == "__main__":
    main()
