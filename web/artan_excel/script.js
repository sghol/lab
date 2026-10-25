const loadButton = document.getElementById("load");
const downloadButton = document.getElementById("download");
const fileInput = document.getElementById("fileInput");
const statusDiv = document.getElementById("status");

let currentFile = null;
let resultRows = null; // ← declared and filled below

function parseExcel(rows) {
  const [header, ...dataRows] = rows;
  const NAME_COL = 1;

  const result = [];
  const newHeader = [...header];
  newHeader.splice(NAME_COL, 1, "نام", "نام خانوادگی");
  result.push(newHeader);

  for (const row of dataRows) {
    if (!row.length || !row[0]) continue;

    const out = [...row];
    while (out.length <= NAME_COL) out.push("");

    const parts = String(out[NAME_COL] ?? "")
      .trim()
      .split(/\s+/)
      .filter(Boolean);
    const firstName = parts[0] ?? "";
    const lastName = parts.slice(1).join(" ");

    out.splice(NAME_COL, 1, firstName, lastName);
    result.push(out);
  }
  return result;
}

async function readExcel(file) {
  const data = await file.arrayBuffer();
  const wb = XLSX.read(data);
  const ws = wb.Sheets[wb.SheetNames[0]];
  const rows = XLSX.utils.sheet_to_json(ws, { header: 1 });
  return parseExcel(rows); // ← returns the rows
}

function writeExcel() {
  if (!resultRows) return;
  try {
    const wb = XLSX.utils.book_new();
    XLSX.utils.book_append_sheet(wb, XLSX.utils.aoa_to_sheet(resultRows), "نتیجه");

    const name = currentFile.name.replace(/\.(xlsx|xls|csv)$/i, "") + "-اصلاح شده.xlsx";
    XLSX.writeFile(wb, name);

    statusDiv.textContent = "فایل در پوشه دانلود ذخیره شد";
  } catch (err) {
    console.error(err);
    statusDiv.textContent = "خطا در ساخت فایل: " + err.message;
  }
}

async function chooseFile() {
  const file = fileInput.files[0];
  if (!file) return;

  currentFile = file;
  statusDiv.textContent = `فایل انتخاب شده: ${file.name} (${(file.size / 1024).toFixed(1)} KB)`;

  try {
    resultRows = await readExcel(file); // ← stored!
    downloadButton.disabled = false;
    downloadButton.classList.replace("button-secondary", "button-success");
  } catch (err) {
    console.error(err);
    statusDiv.textContent = "خطا در خواندن فایل: " + err.message;
    downloadButton.disabled = true;
    downloadButton.classList.replace("button-success", "button-secondary");
  }
}


// --- main ---
// on cancel for filechooser
fileInput.addEventListener("cancel", () => {
  if (currentFile) return;
  statusDiv.textContent = "فایل اکسل را انتخاب کنید";
  downloadButton.classList.remove("button-success");
});

fileInput.addEventListener("change", chooseFile);
loadButton.addEventListener("click", () => fileInput.click());
downloadButton.addEventListener("click", writeExcel);
