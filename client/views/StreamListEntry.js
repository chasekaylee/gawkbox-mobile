import React, { Component } from 'react';
import { StyleSheet, Text, View, Button, Image, WebView } from 'react-native';
import { Card } from 'react-native-elements';

export default (props) => {
  const { item } = props;
  return (
    <Card title={item.channel.name} image={{ uri: item.preview.large }}>
      <Text style={{ marginBottom: 10 }}>{item.channel.status}</Text>
      <Button
        icon={{ name: 'code' }}
        backgroundColor="#03A9F4"
        fontFamily="Lato"
        buttonStyle={{
          borderRadius: 0,
          marginLeft: 0,
          marginRight: 0,
          marginBottom: 0,
        }}
        title="Watch Stream"
        onPress={() => props.navigation.navigate('StreamDetails', { stream: item })}
      />
    </Card>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
