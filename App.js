import React, { Component } from 'react';
import { StyleSheet } from 'react-native';
import { StackNavigator } from 'react-navigation';

import StreamList from './client/views/StreamList';
import StreamListEntry from './client/views/StreamListEntry';
import StreamDetails from './client/views/StreamDetails';
import Home from './client/views/Home';

export default () => <AppNavigator />;

const AppNavigator = StackNavigator({
  Home: { screen: Home },
  StreamListScreen: { screen: StreamList },
  StreamListEntry: { screen: StreamListEntry },
  StreamDetails: { screen: StreamDetails },
});

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
