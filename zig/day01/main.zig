const std = @import("std");

pub fn main() !void {
    const file = @embedFile("input.txt");
    const part1_result = try solvePart1(file);
    std.debug.print("Part 1: {}\n", .{part1_result});

    const part2_result = try solvePart2(file);
    std.debug.print("Part 2: {}\n", .{part2_result});
}

pub fn solvePart1(input: []const u8) !u32 {
    var lines = std.mem.split(u8, input, "\n");
    var max: u32 = 0;
    var current: u32 = 0;

    while (lines.next()) |line| {
        if (line.len == 0) {
            if (current > max) {
                max = current;
            }
            current = 0;
        } else {
            const value = try std.fmt.parseInt(u32, line, 10);
            current += value;
        }
    }
    return max;
}

pub fn solvePart2(input: []const u8) !u32 {
    var lines = std.mem.split(u8, input, "\n");
    var max = [3]u32{ 0, 0, 0 };
    var current: u32 = 0;

    while (lines.next()) |line| {
        if (line.len == 0) {
            for (max, 0..) |n, i| {
                if (current > n) {
                    if (i == 0) {
                        max[2] = max[1];
                        max[1] = max[0];
                        max[0] = current;
                    } else if (i == 1) {
                        max[2] = max[1];
                        max[1] = current;
                    } else {
                        max[2] = current;
                    }
                    break;
                }
            }
            current = 0;
        } else {
            const value = try std.fmt.parseInt(u32, line, 10);
            current += value;
        }
    }
    return @reduce(.Add, @as(@Vector(3, u32), max));
}

test "test-part1" {
    const file = @embedFile("test.txt");
    const result_part_1 = try solvePart1(file);
    try std.testing.expect(result_part_1 == 24000);
}

test "test-part2" {
    const file = @embedFile("test.txt");
    const result_part_2 = try solvePart2(file);
    try std.testing.expect(result_part_2 == 45000);
}
