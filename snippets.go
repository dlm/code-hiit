package main

import (
	"math/rand"
)

var snippets = map[Mode][]CodeSnippet{
	EasyCode: {
		{
			Content: `fn add(a: i32, b: i32) -> i32 {
    a + b
}`,
			Language: "Rust",
			Mode:     EasyCode,
		},
		{
			Content: `int sum = 0;
for (int i = 0; i < 10; i++) {
    sum += i;
}`,
			Language: "C++",
			Mode:     EasyCode,
		},
		{
			Content: `let mut x = 5;
x *= 2;
x -= 1;`,
			Language: "Rust",
			Mode:     EasyCode,
		},
		{
			Content: `std::vector<int> nums = {1, 2, 3, 4, 5};
int total = 0;`,
			Language: "C++",
			Mode:     EasyCode,
		},
		{
			Content: `let mask = 0xFF;
let val = data & mask;
if val != 0 {
    count += 1;
}`,
			Language: "Rust",
			Mode:     EasyCode,
		},
		{
			Content: `int flags = 0x1A | 0x04;
bool is_set = (flags & 0x02) != 0;`,
			Language: "C++",
			Mode:     EasyCode,
		},
		{
			Content: `struct Point {
    x: f32,
    y: f32,
}
let p = Point { x: 1.0, y: 2.0 };`,
			Language: "Rust",
			Mode:     EasyCode,
		},
		{
			Content: `auto ptr = new int[10];
ptr[0] = 42;
delete[] ptr;`,
			Language: "C++",
			Mode:     EasyCode,
		},
		{
			Content: `let arr = [1, 2, 3, 4];
for &item in &arr {
    println!("{}", item);
}`,
			Language: "Rust",
			Mode:     EasyCode,
		},
		{
			Content: `char buffer[256];
sprintf(buffer, "val=%d", 0x2A);`,
			Language: "C++",
			Mode:     EasyCode,
		},
		{
			Content: `fn check(n: i32) -> bool {
    n >= 0 && n <= 100
}`,
			Language: "Rust",
			Mode:     EasyCode,
		},
		{
			Content: `int result = (a << 2) | (b & 0x0F);
bool flag = (result != 0);`,
			Language: "C++",
			Mode:     EasyCode,
		},
		{
			Content: `let opts = vec!["--help", "-v", "--debug"];
let cmd = format!("{} {}", bin, opts[1]);`,
			Language: "Rust",
			Mode:     EasyCode,
		},
		{
			Content: `std::array<int, 5> nums = {10, 20, 30, 40, 50};
int first = nums[0];`,
			Language: "C++",
			Mode:     EasyCode,
		},
		{
			Content: `let (x, y) = (3.14, 2.71);
let sum = x + y;`,
			Language: "Rust",
			Mode:     EasyCode,
		},
		{
			Content: `enum Color { Red = 0xFF0000, Green = 0x00FF00 };
Color c = Color::Red;`,
			Language: "C++",
			Mode:     EasyCode,
		},
	},
	MediumCode: {
		{
			Content: `fn parse_coords(s: &str) -> Option<(i32, i32)> {
    let parts: Vec<&str> = s.split(',').collect();
    if parts.len() != 2 {
        return None;
    }
    Some((parts[0].parse().ok()?, parts[1].parse().ok()?))
}`,
			Language: "Rust",
			Mode:     MediumCode,
		},
		{
			Content: `template<typename T>
T max(T a, T b) {
    return (a > b) ? a : b;
}
auto result = max<int>(42, 17);`,
			Language: "C++",
			Mode:     MediumCode,
		},
		{
			Content: `let data = vec![1, 2, 3, 4, 5];
let sum: i32 = data.iter()
    .filter(|&&x| x % 2 == 0)
    .map(|&x| x * 2)
    .sum();`,
			Language: "Rust",
			Mode:     MediumCode,
		},
		{
			Content: `std::map<std::string, int> counts;
for (const auto& item : items) {
    counts[item]++;
}`,
			Language: "C++",
			Mode:     MediumCode,
		},
		{
			Content: `match value {
    0..=10 => println!("0-10"),
    11..=20 => println!("11-20"),
    _ => println!(">20"),
}`,
			Language: "Rust",
			Mode:     MediumCode,
		},
		{
			Content: `let hash = (key ^ 0xDEADBEEF)
    .wrapping_mul(0x9E3779B9)
    .rotate_left(13);
state[idx] ^= hash;`,
			Language: "Rust",
			Mode:     MediumCode,
		},
		{
			Content: `uint32_t rotate_bits(uint32_t val, int n) {
    return (val << n) | (val >> (32 - n));
}
auto result = rotate_bits(0xABCD, 8);`,
			Language: "C++",
			Mode:     MediumCode,
		},
		{
			Content: `enum Status {
    Ok(T),
    Err { code: i32, msg: &'static str },
}
let res = Status::Err { code: -1, msg: "failed" };`,
			Language: "Rust",
			Mode:     MediumCode,
		},
		{
			Content: `std::unique_ptr<Node> create_node(int val) {
    return std::make_unique<Node>(val, nullptr);
}
auto node = create_node(42);`,
			Language: "C++",
			Mode:     MediumCode,
		},
		{
			Content: `let closure = |x: i32, y: i32| -> i32 {
    let sum = x + y;
    sum * sum - (x ^ y)
};
let result = closure(5, 3);`,
			Language: "Rust",
			Mode:     MediumCode,
		},
		{
			Content: `auto cmp = [](const auto& a, const auto& b) {
    return a.priority > b.priority;
};
std::sort(items.begin(), items.end(), cmp);`,
			Language: "C++",
			Mode:     MediumCode,
		},
	},
	HardCode: {
		{
			Content: `impl<T: Clone + Default> Buffer<T> {
    fn resize(&mut self, new_size: usize) -> Result<(), Error> {
        if new_size > MAX_SIZE {
            return Err(Error::TooLarge);
        }
        self.data.resize(new_size, T::default());
        Ok(())
    }
}`,
			Language: "Rust",
			Mode:     HardCode,
		},
		{
			Content: `auto lambda = [&](int x, int y) -> int {
    return (x * 3 + y << 2) & 0xFF;
};
std::transform(v1.begin(), v1.end(), v2.begin(),
               result.begin(), lambda);`,
			Language: "C++",
			Mode:     HardCode,
		},
		{
			Content: `let result: Result<Vec<_>, _> = (0..10)
    .map(|i| -> Result<i32, Error> {
        data.get(i)
            .ok_or(Error::OutOfBounds)?
            .parse::<i32>()
            .map_err(|_| Error::ParseFailed)
    })
    .collect();`,
			Language: "Rust",
			Mode:     HardCode,
		},
		{
			Content: `std::shared_ptr<Node> node =
    std::make_shared<Node>(42, nullptr);
auto weak = std::weak_ptr<Node>(node);
if (auto locked = weak.lock()) {
    locked->value |= 0x80;
}`,
			Language: "C++",
			Mode:     HardCode,
		},
		{
			Content: `unsafe fn process_raw(ptr: *mut u8, len: usize) {
    let slice = std::slice::from_raw_parts_mut(ptr, len);
    for i in 0..len {
        slice[i] ^= 0xAA;
    }
}`,
			Language: "Rust",
			Mode:     HardCode,
		},
		{
			Content: `constexpr auto factorial(int n) -> int {
    return (n <= 1) ? 1 : n * factorial(n - 1);
}
static_assert(factorial(5) == 120);`,
			Language: "C++",
			Mode:     HardCode,
		},
		{
			Content: `fn parse_packet<'a>(buf: &'a [u8]) -> IResult<&'a [u8], Packet> {
    let (rest, header) = be_u16(buf)?;
    let (rest, payload) = take((header & 0x3FFF) as usize)(rest)?;
    Ok((rest, Packet { ty: (header >> 14) as u8, data: payload }))
}`,
			Language: "Rust",
			Mode:     HardCode,
		},
		{
			Content: `template<typename... Ts>
constexpr auto sum(Ts&&... args) {
    return (... + args);
}
static_assert(sum(1, 2, 3, 4) == 10);`,
			Language: "C++",
			Mode:     HardCode,
		},
		{
			Content: `impl<'a, T> Iterator for ChunkIter<'a, T> {
    type Item = &'a [T];
    fn next(&mut self) -> Option<Self::Item> {
        if self.pos >= self.data.len() { return None; }
        let end = (self.pos + self.size).min(self.data.len());
        let chunk = &self.data[self.pos..end];
        self.pos = end;
        Some(chunk)
    }
}`,
			Language: "Rust",
			Mode:     HardCode,
		},
		{
			Content: `auto async_read = [this](auto&&... args) -> Task<int> {
    co_await scheduler_->yield();
    auto result = co_await file_->read(std::forward<decltype(args)>(args)...);
    co_return result | 0x8000;
};`,
			Language: "C++",
			Mode:     HardCode,
		},
		{
			Content: `macro_rules! impl_op {
    ($trait:ident, $method:ident, $op:tt) => {
        impl $trait for Vec2 {
            type Output = Self;
            fn $method(self, rhs: Self) -> Self::Output {
                Self { x: self.x $op rhs.x, y: self.y $op rhs.y }
            }
        }
    };
}`,
			Language: "Rust",
			Mode:     HardCode,
		},
		{
			Content: `template<class T, class = std::enable_if_t<std::is_integral_v<T>>>
constexpr T popcount(T val) {
    return val == 0 ? 0 : (val & 1) + popcount(val >> 1);
}`,
			Language: "C++",
			Mode:     HardCode,
		},
		{
			Content: `let decoder = |bytes: &[u8]| -> Result<(u32, &[u8]), _> {
    if bytes.len() < 4 { return Err(TooShort); }
    let val = u32::from_le_bytes(bytes[..4].try_into()?);
    Ok(((val & 0xFFFFFF) ^ 0x5A5A5A, &bytes[4..]))
};`,
			Language: "Rust",
			Mode:     HardCode,
		},
	},
	NumbersPractice: {
		{
			Content: `let x = 42 + 17 - 8;
let y = 100 / 5 * 2;
let z = 3.14 + 2.71;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `int a = 256 / 16;
int b = 1024 - 512 + 128;
float pi = 3.14159;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `arr[0] = 99;
arr[1] = 88 - 11;
arr[2] = 77 * 2 / 7;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `port = 8080;
timeout = 30.5;
max_conn = 1000 - 50;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `x = 7 + 3 * 2 - 1;
y = 15 / 3 + 10;
z = 99.99 - 0.99;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `rate = 9.8 / 2.0;
count = 500 + 250 - 125;
total = 1 + 2 + 3 + 4 + 5;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `x := 10 * 10 / 2;
y := 50 - 25 + 75;
z := 6.28 + 1.57;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `v[0] = 123;
v[1] = 456 / 2;
v[2] = 789 - 321;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `fps = 60.0;
width = 1920 / 2;
height = 1080 - 100;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `a = 5 * 5 - 10;
b = 18 / 3 + 7;
c = 2.5 + 3.5 - 1.0;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `index = 42 - 10 / 2;
sum = 7 + 8 + 9;
avg = 100.0 / 3.0;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `byte1 = 255 - 128;
byte2 = 64 * 2 + 1;
byte3 = 192 / 3;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `x = 90 - 45 + 20;
y = 13 * 7 / 2;
z = 0.5 + 0.25 + 0.125;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `val = 888 / 8 - 11;
res = 33 + 66 - 22;
num = 5.5 * 2.0;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `x1 = 16 / 4 + 20;
y1 = 35 * 2 - 15;
z1 = 9.0 - 4.5;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `dim_x = 640 / 2;
dim_y = 480 - 80;
scale = 2.0 * 1.5;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `ip = "192.168.1.254";
port = 3000 + 8080;
subnet = "10.0.0.0/24";`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `version = "2.7.18";
build = 20231205;
patch = 4 + 3 - 1;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `pos_x = 125.75;
pos_y = -34.5 + 100;
pos_z = 0.0 - 12.3;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `timeout_ms = 5000;
retry_count = 3 * 2;
delay = 100 + 50 + 25;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `slice = data[10:50];
chunk = arr[0:100:10];
last = items[-5:];`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `percent = 85.5 / 100.0;
ratio = 16 / 9;
score = 987 - 123;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `mask = 0xFF00 + 0x00FF;
bits = 0b1010 | 0b0101;
oct = 0o755 - 0o644;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `rgb = (255, 128, 64);
alpha = 0.75 + 0.25;
color = 0xABCDEF;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `mac = "00:1A:2B:3C:4D:5E";
vlan = 100 + 200;
mtu = 1500 - 28;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `size_mb = 1024 * 1024;
quota_gb = 500 / 2;
limit_kb = 2048 + 512;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `temp_c = 25.5 + 3.2;
temp_f = (9 / 5) * 20;
kelvin = 273 + 100;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `date = "2024-03-15";
time = 14 * 60 + 30;
epoch = 1234567890;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `coords = (40.7128, -74.0060);
lat = 51.5074 - 0.1278;
lon = -0.0 + 139.6917;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `buffer[128] = 0xAA;
offset = 0x1000 + 0x200;
length = 512 - 64 + 8;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `freq_hz = 440 * 2;
sample_rate = 48000 / 2;
channels = 2 + 0;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `range_min = 0.001;
range_max = 999.999;
step = 0.1 * 10;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `code = 200 + 4;
status = 500 - 100;
redirect = 300 + 1 + 7;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `pid = 12345;
uid = 1000 + 1;
gid = 100 * 10;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `page = 42 / 2;
per_page = 25 + 25;
total = 1337 - 337;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `addr = 0xDEADBEEF;
base = 0x400000;
end = 0xFFFFFF - 0x100;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `seed = 42 * 1337;
rand = 9876 / 4;
hash = 0x5A5A5A5A;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `bits_16 = 65535;
bits_32 = 4294967295;
bits_8 = 255 / 1;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `grid[7][8] = 99;
matrix[3][3] = 1.5;
cube[0][1][2] = 42;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `weight = 72.5 + 0.5;
height = 180 - 5;
bmi = 72.5 / 1.8 / 1.8;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
		{
			Content: `interval = 1000 / 60;
duration = 3600 * 24;
elapsed = 123.456;`,
			Language: "Numbers",
			Mode:     NumbersPractice,
		},
	},
	HexNumbers: {
		{
			Content: `let mask = 0xFF0A & 0x0F0F;
let flag = 0xDEAD | 0xBEEF;
let check = mask ^ flag;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `addr := 0x7FFF_F0A0;
offset := addr + 0x20;
page := offset & 0xFFFF_F000;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `let color = 0x1A2B3C;
let alpha = 0xFF;
let pixel = (color << 8) | alpha;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `const KEY: u32 = 0x9E37_79B1;
state ^= KEY.wrapping_add(0x1337);
state = state.rotate_left(5);`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `let bytes = [0xDE, 0xAD, 0xBE, 0xEF];
let sum = bytes[0] as u16 + bytes[3] as u16;
let carry = (sum & 0xFF00) >> 8;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `uint16_t header = 0x7F;
header |= 0x20;
footer = header ^ 0x5A;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `hash = (seed ^ 0x5A5A5A5A) + 0x1BADB002;
mix = (hash << 7) ^ (hash >> 3);
mask = mix & 0xFFFF;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `guid := 0xAABBCCDDu32;
mask := guid & 0x00FF00FF;
bits := (guid >> 8) & 0x00FF00FF;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `pc = 0x0040_1000;
jump = pc + 0x20;
stack = 0x7FFF_FFF0;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `let table = [0x0, 0x1, 0x8, 0xF];
let idx = i & 0x3;
let nibble = table[idx];`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `crc := 0xFFFF;
crc ^= 0x00FF;
crc = (crc << 4) | (crc >> 12);`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `value = 0xABC ^ 0x123;
swap = (value << 8) | (value >> 8);
mask = swap & 0x0FFF;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `salt := 0x6C696769;
token := salt ^ 0x41414141;
tag := token & 0xFFFF0000;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
		{
			Content: `addr_low := 0x0010;
addr_high := 0xFF00;
pointer := addr_high | addr_low;`,
			Language: "Hex Numbers",
			Mode:     HexNumbers,
		},
	},
	SymbolsPractice: {
		{
			Content: `arr[0] = {x: 1, y: 2};
dict["key"] = "value";
obj->field = nullptr;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `if (a && b || !c) {
    return (x > y) ? x : y;
}`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `mask = ~0xFF & 0xAA;
val = (bits << 4) | 0x0F;
check = !(flag ^ 0x01);`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `"string" + 'c' + r"raw";
path = "/usr/local/bin";
url = "https://example.com?q=test&n=5";`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `list = [1, 2, 3];
tuple = (a, b, c);
set = {x, y, z};`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `fn(a, b, c);
obj.method(x, y);
ptr->call(&arg);`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `regex = /[a-z]+@\w+\.\w{2,}/;
pattern = ^start.*end$;
match = \d{3}-\d{4};`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `expr = (a + b) * (c - d);
calc = {[(x / y) % z]};
nest = <<val>>;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `cmd = "ls -la | grep '.txt'";
pipe = cat file.txt | wc -l;
redir = echo "test" > out.log;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `@decorator
#pragma once
$variable = $$value;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `a += 10;
b -= 5;
c *= 2;
d /= 4;
e %= 3;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `cmp = (a == b) && (c != d);
logic = (x <= y) || (z >= w);
test = !(p < q);`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `map = {"k1": v1, "k2": v2};
access = obj["prop"];
slice = arr[1:5];`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `lambda = (x, y) => x * y;
arrow = () => {return 42;};
func = |a, b| a + b;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `type = Vec<T>;
generic = HashMap<K, V>;
bounds = Box<dyn Trait>;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `ternary = x ? y : z;
null_check = val ?? default;
optional = obj?.prop?.nested;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `bits = 0b1010 & 0b1100;
hex = 0xAB | 0xCD;
complement = ~mask;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `printf("%d %s %f\n", n, s, f);
format!("{}: {} = {}", k, op, v);
sprintf(buf, "<%p>", ptr);`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `path = "./dir/../file.ext";
glob = "**/*.{js,ts}";
home = "~/Documents/";`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `try {
    throw Error("msg");
} catch (e) {
    log(e);
}`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `template<class T>
std::vector<T*> ptrs;
unique_ptr<Node> node;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `a++; ++b;
c--; --d;
x += y *= z;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `match val {
    Some(x) => process(x),
    None => panic!("error"),
}`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `arr = [[1, 2], [3, 4]];
grid = {{"a", "b"}, {"c", "d"}};
deep = [[[x]]];`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `@staticmethod
def func(*args, **kwargs):
    pass`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `enum Flag { A = 1 << 0, B = 1 << 1 };
mask = Flag::A | Flag::B;
check = (val & mask) != 0;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `ptr = &mut value;
ref = *ptr;
deref = **double_ptr;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `sql = "SELECT * FROM t WHERE id = ?";
query = 'INSERT INTO log VALUES (?, ?)';
escape = "quote: \"text\"";`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `#include <vector>
using namespace std;
typedef map<string, int> Dict;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
		{
			Content: `range = 0..100;
exclusive = 0..<10;
infinite = 0..;`,
			Language: "Symbols",
			Mode:     SymbolsPractice,
		},
	},
	BracketsPractice: {
		{
			Content: `fn(a, b, c);
obj.method(x, y);
ptr->call(&arg);`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `arr[0] = vec[i];
map["key"] = data[idx];
grid[x][y] = val;`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `if (a && b) {
    return (x > y) ? x : y;
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `list = [1, 2, 3];
tuple = (a, b, c);
set = {x, y, z};`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `lambda = (x, y) => x * y;
arrow = () => {return 42;};
func = |a, b| a + b;`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `type = Vec<T>;
generic = HashMap<K, V>;
bounds = Box<dyn Trait>;`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `expr = (a + b) * (c - d);
calc = {[(x / y) % z]};
nest = <<val>>;`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `arr = [[1, 2], [3, 4]];
grid = {{"a", "b"}, {"c", "d"}};
deep = [[[x]]];`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `match val {
    Some(x) => process(x),
    None => panic!("error"),
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `template<typename T>
T max(T a, T b) {
    return (a > b) ? a : b;
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `std::map<std::string, int> counts;
for (const auto& item : items) {
    counts[item]++;
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `let data = vec![1, 2, 3];
let sum: i32 = data.iter()
    .filter(|&&x| x % 2 == 0)
    .sum();`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `auto lambda = [&](int x) -> int {
    return (x * 3) & 0xFF;
};`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `impl<T: Clone> Buffer<T> {
    fn resize(&mut self, size: usize) {
        self.data.resize(size, T::default());
    }
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `std::vector<std::vector<int>> matrix = {
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9}
};`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `map = {"k1": v1, "k2": v2};
access = obj["prop"];
slice = arr[1:5];`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `enum Color { Red = 0xFF0000 };
flags = (a | b) & mask;
check = (val & 0x80) != 0;`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `result = [(x, y) for x in range(10) if x % 2 == 0];
dict = {k: v for k, v in items};`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `fn parse<'a>(buf: &'a [u8]) -> Result<&'a [u8], Error> {
    let (rest, data) = take(4)(buf)?;
    Ok((rest, data))
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `auto cmp = [](const auto& a, const auto& b) {
    return a.priority > b.priority;
};
std::sort(items.begin(), items.end(), cmp);`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `macro_rules! vec {
    ($($x:expr),*) => {
        {
            let mut v = Vec::new();
            $(v.push($x);)*
            v
        }
    };
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `template<class T>
std::unique_ptr<T> create() {
    return std::make_unique<T>();
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `closure = |x: i32, y: i32| -> i32 {
    let sum = x + y;
    (sum * sum) - (x ^ y)
};`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `try {
    throw Error("msg");
} catch (e) {
    log(e);
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `struct Point {
    x: f32,
    y: f32,
}
let p = Point { x: 1.0, y: 2.0 };`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `impl Iterator for Range {
    type Item = i32;
    fn next(&mut self) -> Option<Self::Item> {
        if self.pos < self.end {
            let val = self.pos;
            self.pos += 1;
            Some(val)
        } else {
            None
        }
    }
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `constexpr auto factorial(int n) -> int {
    return (n <= 1) ? 1 : n * factorial(n - 1);
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `let decoder = |bytes: &[u8]| -> Result<(u32, &[u8]), _> {
    let val = u32::from_le_bytes(bytes[..4].try_into()?);
    Ok((val & 0xFFFFFF, &bytes[4..]))
};`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `template<typename... Ts>
constexpr auto sum(Ts&&... args) {
    return (... + args);
}`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
		{
			Content: `@decorator
def func(*args, **kwargs):
    result = {k: v for k, v in zip(args, kwargs.values())}
    return [process(x) for x in result.items()]`,
			Language: "Brackets",
			Mode:     BracketsPractice,
		},
	},
	RegexPatterns: {
		{
			Content: `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$
^(?:https?://)?(?:www\.)?[a-z0-9-]+(?:\.[a-z]{2,}){1,3}(?:/[^\s?#]+)?/?$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `\b(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\b
^\[(?:[a-f0-9]{1,4}:){7}[a-f0-9]{1,4}\]$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `^#[0-9a-fA-F]{6}$
^(?:rgb|hsl)a?\(\s*\d{1,3}\s*,\s*\d{1,3}%?\s*,\s*\d{1,3}%?(?:\s*,\s*(?:0?\.\d+|1\.0|1))?\s*\)$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[^\w\s]).{12,64}$
^(?!.*(\w)\1\1)[A-Za-z0-9_-]{8,32}$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `^(?<proto>https?)://(?<host>[^/:]+)(?::(?<port>\d{2,5}))?(?<path>/[^\s?#]*)?(?:\?(?<query>[^\s#]*))?(?:#(?<frag>\S+))?$
^(?<year>\d{4})-(?<month>0[1-9]|1[0-2])-(?<day>0[1-9]|[12]\d|3[01])$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `^(?:\+1-?)?(?:\(?\d{3}\)?[-.\s]?)?\d{3}[-.\s]?\d{4}$
^\$?\d{1,3}(?:,\d{3})*(?:\.\d{2})?$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `\[(?P<label>[^\]]+)\]\((?P<url>[^)]+)\)
\*\*(?P<bold>[^*]+)\*\*|__(?P<underline>[^_]+)__`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `(?:^|\s)@[A-Za-z0-9_]{1,15}(?=\s|$)
(?:^|\s)#[A-Za-z_][\w-]*`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `\b[A-Z]{2}\d{3}(?:-\d{3})?\b
^(?:[A-Z]{2}\d{2}\s?){3}[A-Z]{2}$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `^(?:\.\./|\.?/)?(?:[\w.-]+/)*[\w.-]+\.(?:go|rs|py|js|ts|cpp|h)$
^(?:[\w-]+\/)*[\w-]+(?:\.[\w-]+)?@(?:\d+\.\d+\.\d+|\*)$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `(?:--[a-z0-9-]+(?:=\S+)?|\-[a-zA-Z]+)
(?:^|\s)(?:-I\s*[^\s]+|-D\w+=\w+|/Ox)(?=\s|$)`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `^(?:(Mon|Tue|Wed|Thu|Fri|Sat|Sun),\s)?\d{2}\s[A-Z][a-z]{2}\s\d{4}\s\d{2}:\d{2}:\d{2}\sGMT$
^([01]?\d|2[0-3]):[0-5]\d(:[0-5]\d)?(Z|[+-][01]\d:[0-5]\d)?$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `^\((?:[^)(]+|\((?1)\))*\)$
^\[(?:[^[\]]+|\[(?1)\])*\]$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `(?<!\\)\$(\w+)
\\\$(?:\{(?P<name>[A-Za-z_]\w*)\}|(?P<idx>\d+))`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `(?<=^|\s)version\s*=\s*["']?(?P<ver>\d+\.\d+\.\d+(?:-[\w.]+)?)["']?(?=\s|$)
^\s*(?:[\w-]+\s*=\s*\".*?\\"\s*)+$`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
		{
			Content: `(?x)
^
(?:[A-Za-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[A-Za-z0-9!#$%&'*+/=?^_{|}~-]+)*
 |
 \"(?:[^\"\\r\\n]|\\\\.)+\")
@
(?:[A-Za-z0-9](?:[A-Za-z0-9-]*[A-Za-z0-9])?\.)+[A-Za-z]{2,}
$
`,
			Language: "Regex",
			Mode:     RegexPatterns,
		},
	},
}

func GetRandomSnippet(mode Mode) CodeSnippet {
	if mode == Custom {
		customSnippets, err := LoadCustomSnippets()
		if err != nil || len(customSnippets) == 0 {
			displayPath := "~/.config/code-hiit/snippets.json"
			if path, pathErr := getCustomSnippetsPath(); pathErr == nil {
				displayPath = path
			}
			return CodeSnippet{
				Content:  "# No custom snippets found\n# Create " + displayPath + " to add your own!\n# Example:\n# {\n#   \"snippets\": [\n#     {\"content\": \"your code here\", \"language\": \"Go\"}\n#   ]\n# }",
				Language: "Info",
				Mode:     Custom,
			}
		}
		return customSnippets[rand.Intn(len(customSnippets))]
	}

	snippetList := snippets[mode]
	if len(snippetList) == 0 {
		return CodeSnippet{}
	}
	return snippetList[rand.Intn(len(snippetList))]
}

func GetSnippet(mode Mode, index int) CodeSnippet {
	if mode == Custom {
		customSnippets, err := LoadCustomSnippets()
		if err != nil || len(customSnippets) == 0 {
			return GetRandomSnippet(Custom) // Return help message
		}
		return customSnippets[index%len(customSnippets)]
	}

	snippetList := snippets[mode]
	if len(snippetList) == 0 {
		return CodeSnippet{}
	}
	return snippetList[index%len(snippetList)]
}

func GetNextSnippet(mode Mode, currentIndex int) (CodeSnippet, int) {
	if mode == Custom {
		customSnippets, err := LoadCustomSnippets()
		if err != nil || len(customSnippets) == 0 {
			return GetRandomSnippet(Custom), 0 // Return help message
		}
		nextIndex := (currentIndex + 1) % len(customSnippets)
		return customSnippets[nextIndex], nextIndex
	}

	snippetList := snippets[mode]
	if len(snippetList) == 0 {
		return CodeSnippet{}, 0
	}
	nextIndex := (currentIndex + 1) % len(snippetList)
	return snippetList[nextIndex], nextIndex
}
