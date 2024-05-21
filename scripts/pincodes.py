import csv

csv_file_path = "pin_codes.csv"

start_pin = 100000
end_pin = 101000
pin_codes = list(range(start_pin, end_pin + 1))

with open(csv_file_path, mode="w", newline="", encoding="utf-8") as csv_file:
    csv_writer = csv.writer(csv_file)

    csv_writer.writerow(["Pin Code"])

    for pin_code in pin_codes:
        csv_writer.writerow([pin_code])

print(f"Pin codes generated and saved to {csv_file_path}")
