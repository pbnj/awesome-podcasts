const inquirer = require('inquirer');
const fs = require('fs');
const { promisify } = require('util');
const categories = require('./awesome-podcasts.json');

const writeFile = promisify(fs.writeFile);

const questions = [
  {
    name: 'name',
    message: 'Podcast Title:',
  },
  {
    name: 'desc',
    message: 'Podcast Description:',
  },
  {
    name: 'url',
    message: 'Podcast URL:',
  },
  {
    name: 'category',
    message: 'Podcast Category',
    type: 'list',
    choices: categories.map(c => c.category),
  },
];

inquirer
  .prompt(questions)
  .then(answer => {
    const { name, desc, url } = answer;
    for (const cat of categories) {
      if (cat.category === answer.category) {
        cat.pods.push({ name, desc, url });
      }
    }

    return categories;
  })
  .then(podcasts =>
    writeFile('awesome-podcasts.json', JSON.stringify(podcasts), 'utf-8')
  );
