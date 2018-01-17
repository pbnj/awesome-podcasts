const fs = require('fs');
const util = require('util');
const categories = require('../awesome-podcasts.json');

const writeFile = util.promisify(fs.writeFile);

let readmeFile = `<!-- THIS README FILE HAS BEEN GENERATED AUTOMATICALLY. DO NOT EDIT OR MODIFY BY HAND. SEE CONTRIBUTING.MD -->
# Awesome Podcasts
> 😎 Curated list of awesome programming podcasts  [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

---
`;

for (const category of categories) {
  readmeFile += `\n## ${category.category}\n\n`;
  readmeFile += `> ${category.subtitle}\n\n`;

  for (const pod of category.pods) {
    readmeFile += `- [${pod.name}](${pod.url}) - ${pod.desc}\n`;
  }
}

writeFile('../README.md', readmeFile, 'utf-8');
