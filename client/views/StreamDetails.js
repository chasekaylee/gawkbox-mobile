import React, { Component } from 'react';
import { StyleSheet, View, WebView, ActivityIndicator } from 'react-native';

export default class StreamDetails extends Component {
  constructor(props) {
    super(props);
    this.state = {
      visable: true,
    };
  }

  hideSpinner() {
    this.setState({ visable: false });
  }

  render() {
    const { name } = this.props.navigation.state.params.stream.channel;
    return (
      <View style={styles.container}>
        <WebView
          source={{
            uri: `https://player.twitch.tv/?channel=${name}`,
          }}
          onLoad={() => this.hideSpinner()}
          style={styles.videoPlayer}
        />
        <WebView
          source={{
            uri: `https://www.twitch.tv/${name}/chat`,
          }}
          style={styles.chat}
        />
        {this.state.visable && (
          <ActivityIndicator style={styles.spinner} size="large" color="#1DB8AA" />
        )}
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
  spinner: {
    position: 'absolute',
    top: '50%',
    left: '50%',
  },
  chat: {
    marginTop: 5,
    maxHeight: 300,
    width: 320,
    flex: 1,
  },
  videoPlayer: {
    marginTop: 20,
    maxHeight: 500,
    width: 320,
    flex: 1,
  },
});
