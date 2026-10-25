let resultWorkbook = null;

const fileInput = document.getElementById("fileInput");
const statusEl = document.getElementById("status");
const downloadBtn = document.getElementById("downloadBtn");

function setStatus(msg, kind) {
  statusEl.textContent = msg;
  statusEl.className = kind ? kind : "";
}

fileInput.addEventListener("change", async (e) => {
  const file = e.target.files[0];
  if (!file) return;

  resultWorkbook = null;
  downloadBtn.disabled = true;
  setStatus("Reading file...");

  try {
    // 1. Parse the Excel file
    const buffer = await file.arrayBuffer();
    const workbook = XLSX.read(buffer, { type: "array" });
    const sheet = workbook.Sheets[workbook.SheetNames[0]];
    const rows = XLSX.utils.sheet_to_json(sheet, { header: 1 });

    // 2. Run your validation logic
    //    (this is where your real business rules go)
    const errors = [];
    const results = [["Row", "Column", "Value", "Error"]];

    for (let r = 1; r < rows.length; r++) {
      const row = rows[r];
      if (!row || row.length === 0) continue;

      // Example rule: column B must be a non-empty number
      const colB = row[1];
      if (colB === undefined || colB === "" || isNaN(Number(colB))) {
        errors.push(`Row ${r + 1}: column B must be a number (got "${colB ?? ""}")`);
        results.push([r + 1, "B", colB ?? "", "must be a number"]);
      }

      // Example rule: column A must be non-empty
      const colA = row[0];
      if (colA === undefined || String(colA).trim() === "") {
        errors.push(`Row ${r + 1}: column A must not be empty`);
        results.push([r + 1, "A", "", "must not be empty"]);
      }
    }

    // 3. Build the result workbook
    resultWorkbook = XLSX.utils.book_new();
    const resultSheet = XLSX.utils.aoa_to_sheet(results);
    XLSX.utils.book_append_sheet(resultWorkbook, resultSheet, "Report");

    // 4. Show summary
    if (errors.length === 0) {
      setStatus(`✓ File is valid. ${rows.length - 1} rows checked.`, "ok");
    } else {
      setStatus(`✗ ${errors.length} problem(s) found:\n\n${errors.join("\n")}`, "err");
    }
    downloadBtn.disabled = false;
  } catch (err) {
    setStatus(`✗ Failed to read file: ${err.message}`, "err");
  }
});

downloadBtn.addEventListener("click", () => {
  if (!resultWorkbook) return;
  // Trigger a download of the generated workbook
  XLSX.writeFile(resultWorkbook, "validation-report.xlsx");
});
