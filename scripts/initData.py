import requests
import csv
import os

api_url = "https://randomuser.me/api/?results=5000"

response = requests.get(api_url)

if response.status_code == 200:
    data = response.json()

    users = data.get("results", [])

    csv_file_path = "user.csv"

    file_exists = os.path.exists(csv_file_path)

    with open(csv_file_path, mode="a", newline="", encoding="utf-8") as csv_file:
        csv_writer = csv.writer(csv_file)

        if not file_exists:
            csv_writer.writerow(["Name", "Email"])

        for user in users:
            name = f"{user['name']['first']} {user['name']['last']}"
            email = user['email']
            csv_writer.writerow([name, email])

    print(f"Data appended to {csv_file_path}")

else:
    print(f"Error: Unable to fetch data. Status code: {response.status_code}")
