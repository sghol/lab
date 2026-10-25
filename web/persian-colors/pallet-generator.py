import math

def lin(L, C, h):
    a, b = C*math.cos(math.radians(h)), C*math.sin(math.radians(h))
    l = (L + 0.3963377774*a + 0.2158037573*b) ** 3
    m = (L - 0.1055613458*a - 0.0638541728*b) ** 3
    s = (L - 0.0894841775*a - 1.2914855480*b) ** 3
    return ( 4.0767416621*l - 3.3077115913*m + 0.2309699292*s,
            -1.2684380046*l + 2.6097574011*m - 0.3413193965*s,
            -0.0041960863*l - 0.7034186147*m + 1.7076147010*s)

def in_gamut(L, C, h):
    return all(-1e-4 <= v <= 1.0001 for v in lin(L, C, h))

def max_chroma(L, h):
    lo, hi = 0.0, 0.5
    for _ in range(30):
        mid = (lo + hi) / 2
        if in_gamut(L, mid, h): lo = mid
        else: hi = mid
    return lo

name = input("Color name: ").strip()
L0, C0, h = map(float, input("OKLCH (L C h): ").replace(",", " ").split()[:3])
base_step = int(input("Base step (100-900): ").strip())

max_dist = max(base_step - 100, 900 - base_step) / 100
K = 0.6   # 0.4 = more saturation, 0.8 = more taper

print()
for step in range(100, 1000, 100):
    L = min(max(L0 + (base_step - step) * 0.001, 0.01), 0.99)
    d = abs(step - base_step) / 100
    C = C0 * (1 - K * d / max_dist)
    C = min(C, max_chroma(L, h))
    C = max(C, 0.005)
    print(f"--persian-{name}-{step}: oklch({L:.4f} {C:.4f} {h});")
