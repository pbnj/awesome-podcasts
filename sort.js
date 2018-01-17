const fs = require('fs');
const util = require('util');
const categories = require('./awesome-podcasts.json');

const writeFile = util.promisify(fs.writeFile);

for (const cat of categories) {
  cat.pods.sort((a, b) => a.name.localeCompare(b.name));
}

writeFile('awesome-podcasts.json', JSON.stringify(categories), 'utf-8');
