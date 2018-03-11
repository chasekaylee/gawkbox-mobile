import React, { Component } from 'react';
import { StyleSheet, Text, View, Button, FlatList } from 'react-native';
import { List, ListItem, Card } from 'react-native-elements';
import axios from 'axios';

import SearchBar from '../components/SearchBar';
import StreamListEntry from './StreamListEntry';

export default props => (
  <List>
    <FlatList
      data={props.featured}
      renderItem={({ item: { stream } }) => (
        <StreamListEntry navigation={props.navigation} item={stream} />
      )}
      keyExtractor={item => item.stream._id}
    />
  </List>
);

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
