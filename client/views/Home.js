import React, { Component } from 'react';
import { StyleSheet, View } from 'react-native';
import axios from 'axios';

import SearchBar from '../components/SearchBar';
import StreamList from './StreamList';
import StreamListEntry from './StreamListEntry';
import Offline from './Offline';

export default class Home extends Component {
  constructor(props) {
    super(props);
    this.state = {
      featured: [],
      stream: {
        average_fps: null,
        channel: {},
        created_at: '',
        delay: null,
        game: '',
        is_playlist: null,
        preview: {},
        video_height: null,
        viewers: null,
        _id: null,
      },
      hasSearchData: false,
      offline: false,
    };
    this.search = this.search.bind(this);
    this.fetchStream = this.fetchStream.bind(this);
  }

  async componentWillMount() {
    await this.fetchFeatured();
  }

  search(query) {
    this.fetchStream(query);
  }

  fetchFeatured() {
    this.setState({ loading: true });
    axios
      .get('http://localhost:8080/api/featured')
      .then((data) => {
        this.setState({
          featured: data.data.featured,
          loading: false,
        });
      })
      .catch(console.error);
  }

  fetchStream(searchTerm) {
    this.setState({
      loading: true,
      hasSearchData: false,
      offline: false,
    });
    axios
      .get(`http://localhost:8080/api/search?query=${searchTerm}`)
      .then(({ data }) => {
        this.setState({
          stream: data.stream,
          hasSearchData: true,
          loading: false,
        });
      })
      .catch((error) => {
        this.setState({ loading: false });
        if (error.response && error.response.status === 400) {
          this.setState({ offline: true });
        } else {
          console.log(error);
        }
      });
  }

  render() {
    let list = null;
    let searchEntry = null;
    let offlineMessage = null;
    if (!this.state.offline) {
      if (this.state.hasSearchData) {
        searchEntry = (
          <StreamListEntry navigation={this.props.navigation} item={this.state.stream} />
        );
      } else {
        list = <StreamList navigation={this.props.navigation} featured={this.state.featured} />;
      }
    } else {
      offlineMessage = <Offline />;
    }

    return (
      <View>
        <SearchBar search={this.search} loading={this.state.loading} />
        {list}
        {searchEntry}
        {offlineMessage}
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
