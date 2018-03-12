import React, { Component } from 'react';
import { StyleSheet, Text, View, Button, TextInput, WebView } from 'react-native';

export default class SearchBar extends Component {
  constructor(props) {
    super(props);
    this.state = {
      query: '',
    };
  }

  render() {
    const { container, search, button } = styles;
    return (
      <View style={container}>
        <TextInput
          style={search}
          onChangeText={query => this.setState({ query })}
          value={this.state.query}
          placeholder="Search.."
        />
        <Button
          style={button}
          title={this.props.loading ? 'loading...' : 'Search'}
          onPress={() => this.props.search(this.state.query)}
        />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    marginTop: 20,
    marginLeft: 10,
    marginRight: 10,
  },
  search: {
    flex: 1,
    borderBottomColor: 'black',
  },
  button: {
    height: 30,
    marginBottom: 8,
  },
});
