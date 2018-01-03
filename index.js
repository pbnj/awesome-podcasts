const fs = require("fs")
const util = require("util")

const readFile = util.promisify(fs.readFile)
const writeFile = util.promisify(fs.writeFile)

const main = async () => {
  let awesomePodcasts
  try {
    awesomePodcasts = await readFile("awesome-podcasts.json", "utf-8")
  } catch (e) {
    console.error(e)
  }
  const categories = JSON.parse(awesomePodcasts)
  let readmeFile = `<!-- THIS README FILE HAS BEEN GENERATED AUTOMATICALLY. DO NOT EDIT OR MODIFY BY HAND. SEE CONTRIBUTING.MD -->
# Awesome Podcasts
> 😎 Curated list of awesome programming podcasts  [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

---
`
  categories.sort((a, b) => a.category.localeCompare(b.category))

  for (const category of categories) {
    readmeFile += `\n## ${category.category}\n\n`
    readmeFile += `> ${category.subtitle}\n\n`

    category.pods.sort((a, b) => a.name.localeCompare(b.name))

    for (const pod of category.pods) {
      readmeFile += `- [${pod.name}](${pod.url}) - ${pod.desc}\n`
    }
  }

  try {
    writeFile("README.md", readmeFile)
  } catch (e) {
    console.error(e)
  }
}

main()
