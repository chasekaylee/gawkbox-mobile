import React, { Component } from 'react';
import { StyleSheet, FlatList } from 'react-native';
import { List } from 'react-native-elements';

import StreamListEntry from './StreamListEntry';

export default ({ navigation, featured }) => (
  <List>
    <FlatList
      data={featured}
      renderItem={({ item: { stream } }) => (
        <StreamListEntry navigation={navigation} item={stream} />
      )}
      keyExtractor={item => item.stream._id}
    />
  </List>
);
