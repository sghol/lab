from datetime import datetime


packages_data = None
with open("packages") as pkg_file:
    packages_data = pkg_file.read()


packages = packages_data.split("\n\n")

packages_list = []

for package in packages:
    package_row = package.split("\n")
    package_dict = {}
    for row in package_row:
        field, *values = row.split(" :")
        field = field.strip()
        if field in ["Name", "Install Date"]:
            package_dict[field] = "".join(values).strip()
            packages_list.append(package_dict)


packages = []

date_format = "%a %d %b %Y %I:%M:%S %p %z"

for pkg in packages_list:
    name = pkg["Name"]
    install_date = pkg["Install Date"]
    install_date = datetime.strptime(install_date, date_format)
    packages.append({"Name": name, "Install Date": install_date})
#

sorted_pkgs = sorted(packages, key=lambda row: row["Install Date"])

pkgs_name = [pkg["Name"] for pkg in sorted_pkgs]
# pkgs = set(pkgs_name)

# remove repreated packages
pkgs = []
for pkg in pkgs_name:
    if pkg not in pkgs:
        pkgs.append(pkg)

for i, pkg in enumerate(pkgs):
    print(i, pkg)
