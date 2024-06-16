const fs = require("fs");
const path = require("path");
const { exec } = require("child_process");

function getNonIgnoredFilePaths() {
  return new Promise((resolve, reject) => {
    exec(
      "git ls-files --cached --others --exclude-standard",
      (error, stdout, stderr) => {
        if (error) {
          reject(error);
          return;
        }
        if (stderr) {
          reject(new Error(stderr));
          return;
        }

        const paths = stdout.trim().split("\n");
        resolve(paths);
      },
    );
  });
}

// Main function
async function main() {
  const currentDir = process.cwd() + "/..";
  const files = await getNonIgnoredFilePaths();
  const excludeFiles = [
    "go.mod",
    "go.sum",
    ".git",
    "project-content.js",
    "README.md",
    "api.md",
    "architecture.md",
    "env.example",
    ".gitignore",
    ".air.toml",
    "profile.pb.go",
    "profile_grpc.pb.go",
  ];
  let output = "";

  files.forEach((filePath) => {
    filePath = path.join(currentDir, filePath).replaceAll("\\", "/");
    for (const excludeFileName of excludeFiles)
      if (filePath.endsWith(excludeFileName)) return;
    console.log(filePath);
    output += `Path: ${filePath}\n`;
    output += `Content:\n\`\`\`\n${fs.readFileSync(filePath, "utf8")}\n\`\`\`\n\n`;
  });

  const outputFile = "project-content.txt"; // Output file name

  fs.writeFileSync(outputFile, "");
  fs.writeFileSync(outputFile, output);

  console.log(`Results written to ${outputFile}`);
}

main();
