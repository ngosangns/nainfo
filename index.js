const fs = require('fs');
const path = require('path');

// Function to list all files and folders with tree structure
function listFilesAndFolders(dir, prefix = '') {
    const files = fs.readdirSync(dir);
    let output = '';
    files.forEach(file => {
        const filePath = path.join(dir, file);
        const stats = fs.statSync(filePath);
        if (stats.isDirectory()) {
            output += `${prefix}${file}/\n`;
            output += listFilesAndFolders(filePath, `${prefix}  `); // Recursive call for subdirectories
        } else {
            output += `${prefix}${file}\n`;
        }
    });
    return output;
}

// Function to print path and content of each file, including files in child folders
function printFilePathAndContent(dir) {
    const files = fs.readdirSync(dir);
    files.forEach(file => {
        const filePath = path.join(dir, file);
        const stats = fs.statSync(filePath);
        if (stats.isFile()) {
            console.log(`Path: ${filePath}`);
            console.log(`Content:\n\`\`\`\n${fs.readFileSync(filePath, 'utf8')}\n\`\`\`\n`);
        } else if (stats.isDirectory()) {
            console.log(`Folder: ${filePath}/`);
            printFilePathAndContent(filePath); // Recursive call for subdirectories
        }
    });
}

// Main function
function main() {
    const currentDir = process.cwd();
    console.log("List of files and folders in the current directory:");
    console.log(listFilesAndFolders(currentDir));

    console.log("Printing path and content of each file, including files in child folders:");
    printFilePathAndContent(currentDir);
}

main();
