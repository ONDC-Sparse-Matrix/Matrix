import xlsx from 'xlsx';
import fs from 'fs';


// Read the Excel file
const workbook = xlsx.readFile('user_data.xls');
const sheetName = workbook.SheetNames[0];
const sheet = workbook.Sheets[sheetName];

// Read the zipcodes.txt file
const zipcodes = fs.readFileSync('zipcodes.txt', 'utf8').split('\n').map(line => line.trim()).filter(line => line !== '');

// Convert sheet to JSON
const merchants = xlsx.utils.sheet_to_json(sheet);

// Function to get a random subset of pincodes
function getRandomPincodes(pincodeList, count) {
const shuffled = pincodeList.sort(() => 0.5 - Math.random());
return shuffled.slice(0, count).join(', ');
}

// Add Pincodes column
merchants.forEach(merchant => {
// Randomly determine the number of pincodes to assign (for example, between 5 and 10)
const numPincodes = Math.floor(Math.random() * 6) + 5;
merchant.Pincodes = getRandomPincodes(zipcodes, numPincodes);
});

// Convert JSON back to sheet
const newSheet = xlsx.utils.json_to_sheet(merchants);

// Replace old sheet with new sheet
workbook.Sheets[sheetName] = newSheet;

// Write the updated workbook to a new file
xlsx.writeFile(workbook, 'updated_merchants.xlsx');

console.log('Pincodes added and Excel file updated successfully.');