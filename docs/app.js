/* global React, ReactDOM, fetch, semanticUIReact */
const podcastURL =
  "https://raw.githubusercontent.com/petermbenjamin/awesome-podcasts/master/awesome-podcasts.json"

const { Container, Divider, Dropdown, Header, List, Segment } = semanticUIReact

class App extends React.Component {
  state = { categories: [], showCategories: [] }

  componentDidMount() {
    fetch(podcastURL)
      .then(r => r.json())
      .then(d => {
        this.setState({
          categories: d.categories,
          showCategories: d.categories,
        })
      })
  }

  componentDidUpdate(prevProps, prevState) {
    if (this.state.showCategories.length === 0) {
      this.setState({ showCategories: this.state.categories })
    }
  }

  handleDropdownChange = (event, { name, value }) => {
    this.setState({
      showCategories: this.state.categories.filter(cat =>
        value.includes(cat.name)
      ),
    })
  }

  render() {
    const { categories, showCategories } = this.state
    return (
      <Container>
        <Segment.Group horizontal>
          <Segment>
            <Header as="h1">Awesome Podcasts</Header>
          </Segment>
          <Segment>
            <Dropdown
              placeholder="Filter Categories"
              fluid
              multiple
              search
              selection
              options={categories.map(c => ({
                key: c.name,
                value: c.name,
                text: c.name,
              }))}
              onChange={this.handleDropdownChange}
            />
          </Segment>
        </Segment.Group>
        {showCategories.map(category => (
          <Segment key={category.name}>
            <Header as="h2">
              {category.name}
              <Header.Subheader>{category.subtitle}</Header.Subheader>
            </Header>
            <Divider />
            <List divided relaxed selection>
              {category.pods.map(pod => (
                <List.Item key={pod.name}>
                  <List.Content>
                    <a href={pod.url}>
                      <List.Header>{pod.name}</List.Header>
                      <List.Description>{pod.desc}</List.Description>
                    </a>
                  </List.Content>
                </List.Item>
              ))}
            </List>
          </Segment>
        ))}
      </Container>
    )
  }
}

const root = document.getElementById("root")
const elem = <App />
ReactDOM.render(elem, root)
