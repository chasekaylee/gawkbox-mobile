import React, { Component } from 'react';
import { StyleSheet, Text, View, Button, Image, WebView } from 'react-native';

export default props => (
  <View style={styles.container}>
    <WebView
      source={{
        uri: `https://player.twitch.tv/?channel=${
          props.navigation.state.params.stream.channel.name
        }`,
      }}
      style={{
        marginTop: 20,
        maxHeight: 500,
        width: 320,
        flex: 1,
      }}
    />
    <WebView
      source={{
        uri: `https://www.twitch.tv/${props.navigation.state.params.stream.channel.name}/chat`,
      }}
      style={{
        marginTop: 5,
        maxHeight: 300,
        width: 320,
        flex: 1,
      }}
    />
  </View>
);

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
